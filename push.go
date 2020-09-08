package gomeng

import (
	"fmt"
	"strings"
	"time"
)

/**
 * 推送给单用户（单播）
 */
func (c *Client) Push(payload map[string]interface{}, deviceToken string) error {
	url := fmt.Sprintf("%s/%s", BaseURL, Push)
	param := map[string]interface{}{
		"appkey":          c.AppKey,
		"timestamp":       time.Now().Unix(),
		"type":            "unicast",
		"device_tokens":   deviceToken,
		"payload":         payload,
		"production_mode": c.ProductMode,
	}
	return c.Request("POST", url, param)
}

/**
 * 推送给多用户（列播）
 */
func (c *Client) ListCast(payload map[string]interface{}, deviceTokens ...string) error {
	url := fmt.Sprintf("%s/%s", BaseURL, Push)
	param := map[string]interface{}{
		"appkey":          c.AppKey,
		"timestamp":       time.Now().Unix(),
		"type":            "listcast",
		"device_tokens":   strings.Join(deviceTokens, ","),
		"payload":         payload,
		"production_mode": c.ProductMode,
	}
	return c.Request("POST", url, param)
}

/**
 * 推送给所有用户（广播）
 * 默认每天可推送10次
 */
func (c *Client) Broadcast(payload map[string]interface{}) error {
	url := fmt.Sprintf("%s/%s", BaseURL, Broadcast)
	param := map[string]interface{}{
		"appkey":          c.AppKey,
		"timestamp":       time.Now().Unix(),
		"type":            "broadcast",
		"payload":         payload,
		"production_mode": c.ProductMode,
	}
	return c.Request("POST", url, param)
}
