package gomeng

import (
	"testing"
)

const (
	TestAppKey      = "app_key"
	TestAppSecret   = "app_secret"
	TestProductMode = false
	TestDeviceToken = "AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG"
)

func TestPush2SingleUser(t *testing.T) {
	c := NewClient(TestProductMode, TestAppKey, TestAppSecret, DefaultTimeout)
	payload := map[string]interface{}{
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
	}

	if err := c.Push(payload, TestDeviceToken); err != nil {
		t.Fatalf("err: %v", err)
	}
}

func TestPush2MultiUsers(t *testing.T) {
	c := NewClient(TestProductMode, TestAppKey, TestAppSecret, DefaultTimeout)
	payload := map[string]interface{}{
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
	}

	if err := c.ListCast(payload, "AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG"); err != nil {
		t.Fatalf("err: %v", err)
	}
}

func TestPush2AllUsers(t *testing.T) {
	c := NewClient(TestProductMode, TestAppKey, TestAppSecret, DefaultTimeout)
	payload := map[string]interface{}{
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
	}

	if err := c.Broadcast(payload); err != nil {
		t.Fatalf("err: %v", err)
	}
}
