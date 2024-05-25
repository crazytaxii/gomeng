# gomeng ![CI Status](https://github.com/crazytaxii/gomeng/actions/workflows/ci.yaml/badge.svg)

友盟消息推送 Golang SDK
友盟开发者中心 U-Push API 集成文档：[传送门](https://developer.umeng.com/docs/66632/detail/68343)

## 使用示例

**需要先在友盟官网创建应用获得 App Key 和 App Secret！**

- `Unicast()` 单播
- `Listcast()` 多播
- `Broadcast()`广播（默认每天可推送10次）

```go
import (
    umeng "github.com/crazytaxii/gomeng"
)

func main() {
    cfg := &umeng.Config{
        AppKey: "app_key",
        AppSecret: "app_master_secret",
        ProductionMode: false,
    }
    client := umeng.NewClient(cfg)

    demo := umeng.Payload{
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

    retMsg, err := client.Unicast(context.Background() ,demo, "device_token")
    if err != nil {
        log.Fatalf("failed to unicast: %v", err)
    }
    if err := rm.Error(); err != nil {
        // business error
        log.Fatal(err)
    }

    retMsg, err := client.Listcast(context.Background(), demo, "device_token1", "device_token2", "device_token3")
    if err != nil {
        log.Fatalf("failed to listcast: %v", err)
    }
    if err := rm.Error(); err != nil {
        // business error
        log.Fatal(err)
    }

    retMsg, err := client.Broadcast(context.Background(), demo)
    if err != nil {
        fmt.Fatalf("failed to broadcast: %v", err)
    }
    if err := rm.Error(); err != nil {
        // business error
        log.Fatal(err)
    }
}
```
