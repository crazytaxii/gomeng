package gomeng_test

import (
	"fmt"
	"testing"

	umeng "github.com/crazytaxii/gomeng"
)

const (
	APP_KEY           = "app_key"
	APP_MASTER_SECRET = "app_master_secret"
	PRODUCT_MODE      = false
)

func TestPush2SingleUser(t *testing.T) {
	client := umeng.NewClient(PRODUCT_MODE, APP_KEY, APP_MASTER_SECRET)
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

	err := client.Push2SingleUser("AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG",
		&payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestPush2MultiUsers(t *testing.T) {
	client := umeng.NewClient(PRODUCT_MODE, APP_KEY, APP_MASTER_SECRET)
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

	err := client.Push2MultiUsers([]string{"AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG"},
		&payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestPush2AllUsers(t *testing.T) {
	client := umeng.NewClient(PRODUCT_MODE, APP_KEY, APP_MASTER_SECRET)
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

	err := client.Push2AllUsers(&payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}
