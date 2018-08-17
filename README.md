# gomeng
友盟消息推送Golang SDK
友盟开发者中心U-Push API集成文档：[传送门](https://developer.umeng.com/docs/66632/detail/68343)

## 安装
```Bash
$ go get github.com/crazytaxii/gomeng
```

## 使用示例

**需要先在友盟官网创建应用获得AppKey和AppSecret！**

+ `Push2SingleUser()` 推送给单用户 （单播类消息暂无推送限制）
+ `Push2MultiUsers()` 推送给多用户
+ `Push2AllUsers()` 推送给所有用户（默认每天可推送10次）

```Go
import (
    umeng "github.com/crazytaxii/gomeng"
)

func push() {
    client := umeng.NewClient(false, "app_key", "app_master_secret")
    payload := map[string]interface{}{
        "display_type": "notification",
        "body": map[string]interface{}{
            "ticker":   "test_ticker",
            "title":    "test_title",
            "text":     "test_text",
            "builder:": 1,
            "custom": map[string]interface{}{
                "key1": "value1",
                "key2": "value2",
                "key3": "value3",
            },
            "after_open": "go_app",
            "play_sound": "true",
        },
    }

    err := client.Push2SingleUser("device_token", &payload)
    if err != nil {
        fmt.Println("err:", err)
    }

    err := client.Push2MultiUsers([]string{"device_token1", "device_token2", "device_token3"},
        &payload)
    if err != nil {
        fmt.Println("err:", err)
    }

    err := client.Push2AllUsers(&payload)
    if err != nil {
        fmt.Println("err:", err)
    }
}
```
