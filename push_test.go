package gomeng

import (
	"context"
	"os"
	"testing"
)

func newTest() (client *Client, payload Payload, token string) {
	cfg := &Config{
		AppKey:    os.Getenv("APP_KEY"),
		AppSecret: os.Getenv("APP_SECRET"),
		Timeout:   defaultTimeout,
	}
	return NewClient(cfg), Payload{
		"display_type": "notification",
		"body": map[string]interface{}{
			"ticker":      "test_ticker",
			"title":       "test_title",
			"text":        "test_text",
			"builder_id:": 1,
			"custom": map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			"after_open": "go_app",
			"play_sound": "true",
		},
	}, os.Getenv("DEVICE_TOKEN")
}

func TestPush2SingleUser(t *testing.T) {
	c, payload, token := newTest()
	rm, err := c.Unicast(context.Background(), payload, token)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	if err := rm.Error(); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
	}
}

func TestPush2MultiUsers(t *testing.T) {
	c, payload, token := newTest()
	rm, err := c.ListCast(context.Background(), payload, token)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	if err := rm.Error(); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
	}
}

func TestPush2AllUsers(t *testing.T) {
	c, payload, _ := newTest()
	rm, err := c.Broadcast(context.Background(), payload)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	if err := rm.Error(); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
	}
}
