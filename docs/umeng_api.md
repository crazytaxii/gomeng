# 友盟 U-Push Golang SDK

**使用 API 前需要在Web后台的 App 应用信息页面获取 App key 和 App Master Secret，同时在 Web 后台添加服务器 IP 地址至白名单安或关闭IP白名单。**

## 向指定的用户推送消息（单播）

| 参数 | 说明 |
| --- | --- |
| deviceToken | 友盟消息推送服务对设备的唯一标识 |
| payload | 消息具体内容 |

- Android 平台 payload

    ```json
    {
        "appkey": "xx", // 必填，应用唯一标识
        "timestamp": "xx", // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
        "type": "xx", // 必填，消息发送类型,其值可以为:
                    // unicast-单播
                    // listcast-列播，要求不超过500个device_token
                    // filecast-文件播，多个device_token可通过文件形式批量发送
                    // broadcast-广播
                    // groupcast-组播，按照filter筛选用户群, 请参照filter参数
                    // customizedcast，通过alias进行推送，包括以下两种case:
                    // - alias: 对单个或者多个alias进行推送
                    // - file_id: 将alias存放到文件后，根据file_id来推送
        "device_tokens": "xx", // 当type=unicast时必填, 表示指定的单个设备
                            // 当type=listcast时必填, 要求不超过500个, 以英文逗号分隔
        "alias_type": "xx", // 当type=customizedcast时必填
                            // alias的类型, alias_type可由开发者自定义, 开发者在SDK中
                            // 调用setAlias(alias, alias_type)时所设置的alias_type
        "alias": "xx", // 当type=customizedcast时, 选填(此参数和file_id二选一)
                    // 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
                    // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
        "file_id": "xx", // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
                        // 当type=customizedcast时，选填(此参数和alias二选一)
                        // file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应的alias_type必须和接口参数alias_type一致
                        // 使用文件播需要先调用文件上传接口获取file_id，参照"文件上传"
        "filter": {}, // 当type=groupcast时必填，用户筛选条件，如用户标签、渠道等
        "payload": {
            // 必填，JSON格式，具体消息内容(Android最大为1840B)
            "display_type": "xx", // 必填，消息类型: notification(通知)、message(消息)
            "body": {
                // 必填，消息体
                // 当display_type=message时，body的内容只需填写custom字段
                // 当display_type=notification时，body包含如下参数:
                // 通知展现内容:
                "ticker": "xx", // 必填，通知栏提示文字
                "title": "xx", // 必填，通知标题
                "text": "xx", // 必填，通知文字描述

                // 自定义通知图标:
                "icon": "xx", // 可选，状态栏图标ID
                "largeIcon": "xx", // 可选，通知栏拉开后左侧图标ID
                "img": "xx", // 可选，通知栏大图标的URL链接

                // 自定义通知声音:
                "sound": "xx", // 可选，通知声音

                // 自定义通知样式:
                "builder_id": "xx", // 可选，默认为0，用于标识该通知采用的样式。使用该参数时，开发者必须在SDK里面实现自定义通知栏样式

                // 通知到达设备后的提醒方式
                "play_vibrate": "true/false", // 可选，收到通知是否震动，默认为"true"
                "play_lights": "true/false", // 可选，收到通知是否闪灯，默认为"true"
                "play_sound": "true/false", // 可选，收到通知是否发出声音，默认为"true"

                // 点击"通知"的后续行为，默认为打开app。
                "after_open": "xx", // 可选，默认为"go_app"，值可以为:
                                    // "go_app": 打开应用
                                    // "go_url": 跳转到URL
                                    // "go_activity": 打开特定的activity
                                    // "go_custom": 用户自定义内容
                "url": "xx", // 当after_open=go_url时必填
                            // 通知栏点击后跳转的URL，要求以http或者https开头
                "activity": "xx", // 当after_open=go_activity时必填，通知栏点击后打开的Activity
                "custom": {}, // 当display_type=message时必填
                            // 当display_type=notification且after_open=go_custom时必填
                            // 用户自定义内容，可以为字符串或者JSON格式
            },
            "extra": {
                // 可选，JSON格式，用户自定义key-value。只对"通知"(display_type=notification)生效
                // 可以配合通知到达后，打开App/URL/Activity使用
                "key1": "value1",
                "key2": "value2",
            }
        },
        "policy": {
            // 可选，发送策略
            "start_time": "xx", // 可选，定时发送时，若不填写表示立即发送
                                // 定时发送时间不能小于当前时间
                                // 格式: "yyyy-MM-dd HH:mm:ss"
            "expire_time": "xx", // 可选，消息过期时间，其值不可小于发送时间或者start_time
            "max_send_num": "xx", // 可选，发送限速，每秒发送的最大条数。最小值1000
            "out_biz_no": "xx" // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送
        },
        "production_mode": "true/false", // 可选，正式/测试模式。默认为true
                                        // 测试模式只会将消息发给测试设备。测试设备需要到web上添加
                                        // Android: 测试设备属于正式设备的一个子集
        "description": "xx", // 可选，发送消息描述，建议填写
        "mipush": "true/false", // 可选，默认为false。当为true时，表示MIUI、EMUI、Flyme系统设备离线转为系统下发
        "mi_activity": "xx", // 可选，mipush值为true时生效，表示走系统通道时打开指定页面acitivity的完整包路径
    }
    ```

- iOS 平台 payload

    ```json
    {
        "appkey": "xx", // 必填，应用唯一标识
        "timestamp": "xx", // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
        "type": "xx", // 必填，消息发送类型,其值可以为:
                    // unicast-单播
                    // listcast-列播，要求不超过500个device_token
                    // filecast-文件播，多个device_token可通过文件形式批量发送
                    // broadcast-广播
                    // groupcast-组播，按照filter筛选用户群, 请参照filter参数
                    // customizedcast，通过alias进行推送，包括以下两种case:
                    // - alias: 对单个或者多个alias进行推送
                    // - file_id: 将alias存放到文件后，根据file_id来推送
        "device_tokens": "xx", // 当type=unicast时必填, 表示指定的单个设备
                            // 当type=listcast时必填, 要求不超过500个, 以英文逗号分隔
        "alias_type": "xx", // 当type=customizedcast时必填
                            // alias的类型, alias_type可由开发者自定义, 开发者在SDK中
                            // 调用setAlias(alias, alias_type)时所设置的alias_type
        "alias": "xx", // 当type=customizedcast时, 选填(此参数和file_id二选一)
                    // 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
                    // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
        "file_id": "xx", // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
                        // 当type=customizedcast时，选填(此参数和alias二选一)
                        // file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应
                        // 的alias_type必须和接口参数alias_type一致
                        // 使用文件播需要先调用文件上传接口获取file_id，参照"2.4文件上传接口"
        "filter": {}, // 当type=groupcast时必填，用户筛选条件，如用户标签、渠道等
        "payload":
        {
            // 必填，JSON格式，具体消息内容(iOS最大为2012B)
            "aps":
            {
                // 必填，严格按照APNs定义来填写
                "alert": {
                    // 当content-available=1时(静默推送)，可选; 否则必填
                    // 可为JSON类型和字符串类型
                    "title": "title",
                    "subtitle": "subtitle",
                    "body": "body"
                },
                "badge": "xx", // 可选
                "sound": "xx", // 可选
                "content-available": 1, // 可选，代表静默推送
                "category": "xx", // 可选，注意: iOS8才支持该字段
            },
            "key1": "value1",
            "key2": "value2",
        },
        "policy":
        {
            // 可选，发送策略
            "start_time": "xx", // 可选，定时发送时间，若不填写表示立即发送。
                                // 定时发送时间不能小于当前时间
                                // 格式: "yyyy-MM-dd HH:mm:ss"
                                // 注意，start_time只对任务生效
            "expire_time": "xx", // 可选，消息过期时间，其值不可小于发送时间或者start_time(如果填写了的话)
                                // 如果不填写此参数，默认为3天后过期。格式同start_time
            "out_biz_no": "xx", // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送
            "apns_collapse_id": "xx" // 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展最新的一条，字段长度不得超过64bytes
        },
        "production_mode": "true/false", // 可选，正式/测试模式。默认为true
        "description": "xx" // 可选，发送消息描述，建议填写
    }
    ```
