package gomeng_test

import (
	"fmt"
	"testing"

	umeng "github.com/crazytaxii/gomeng"
)

func TestUmengInit(t *testing.T) {
	umeng.Init2(
		true,
		"app_key_4_android",
		"app_secret_4_android",
	)
}

func TestPush2SingleUser(t *testing.T) {
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

	err := umeng.Push2SingleUser("android", "AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG", &payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestPush2MultiUsers(t *testing.T) {
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

	err := umeng.Push2MultiUsers("android", []string{"AtOAal-11NoRhG1KJv_aq1aij5O_aWwMlvvklGNu1LmG"}, &payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestPush2AllUsers(t *testing.T) {
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

	err := umeng.Push2AllUsers("android", &payload)
	if err != nil {
		fmt.Println("err:", err)
	}
}
