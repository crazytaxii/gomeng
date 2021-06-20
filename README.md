# gomeng

友盟消息推送 Golang SDK
友盟开发者中心 U-Push API 集成文档：[传送门](https://developer.umeng.com/docs/66632/detail/68343)

## 使用示例

**需要先在友盟官网创建应用获得 App Key 和 App Secret！**

+ `Push()` 单播
+ `Listcast()` 多播
+ `Broadcast()`广播（默认每天可推送10次）

```go
import (
    umeng "github.com/crazytaxii/gomeng"
)

func main() {
    client := umeng.NewClient(false, "app_key", "app_master_secret", 10 * time.Second)
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

    if err := client.Push(payload, "device_token"); err != nil {
        log.Fatalf("err: %v", err)
    }

    if err := client.Listcast(payload, "device_token1", "device_token2", "device_token3"); err != nil {
        log.Fatalf("err: %v", err)
    }

    if err := client.Broadcast(payload); err != nil {
        fmt.Fatalf("err: %v", err)
    }
}
```
