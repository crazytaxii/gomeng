package gomeng

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

/**
 * 推送给单用户（单播）
 */
func Push2SingleUser(platform string, deviceToken string,
	payload *map[string]interface{}) error {
	url := fmt.Sprintf("%s%s", DOMAIN, "/api/send")

	if !(platform == PLATFORM_IOS || platform == PLATFORM_ANDROID) {
		return fmt.Errorf("Umeng push failed, wrong platform")
	}

	param := map[string]interface{}{
		"appkey":        configMap[platform].AppKey,
		"timestamp":     time.Now().Unix(),
		"type":          "unicast",
		"device_tokens": deviceToken,
		"payload":       *payload,
	}

	// 签名
	sign, err := doSign("POST", url, param, configMap[platform].AppMasterSecret)
	if err != nil {
		return err
	}

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}

	request := gorequest.New()
	_, body, _ := request.Post(fmt.Sprintf("%s?sign=%s", url, sign)).
		Send(string(data)).End()
	result := new(struct {
		Ret  string `json:"ret"`
		Data struct {
			MsgId   string `json:"msg_id"`
			ErrMsg  string `json:"error_msg"`
			ErrCode string `json:"error_code"`
		} `json:"data"`
	})
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return err
	}
	fmt.Println("result:", result)
	if result.Ret != "SUCCESS" {
		return fmt.Errorf("Umeng push failed, error message: %s, error code: %s",
			result.Data.ErrMsg, result.Data.ErrCode)
	}
	return nil
} // Push2SingleUser()

/**
 * 推送给多用户（列播）
 */
func Push2MultiUsers(platform string, listDeviceToken []string,
	payload *map[string]interface{}) error {
	url := fmt.Sprintf("%s%s", DOMAIN, "/api/send")

	if !(platform == PLATFORM_IOS || platform == PLATFORM_ANDROID) {
		return fmt.Errorf("Umeng push failed, wrong platform")
	}

	param := map[string]interface{}{
		"appkey":        configMap[platform].AppKey,
		"timestamp":     time.Now().Unix(),
		"type":          "listcast",
		"device_tokens": strings.Join(listDeviceToken, ","),
		"payload":       *payload,
	}

	// 签名
	sign, err := doSign("POST", url, param, configMap[platform].AppMasterSecret)
	if err != nil {
		return err
	}

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}

	request := gorequest.New()
	_, body, _ := request.Post(fmt.Sprintf("%s?sign=%s", url, sign)).
		Send(string(data)).End()
	result := new(struct {
		Ret  string `json:"ret"`
		Data struct {
			MsgId   string `json:"msg_id"`
			ErrMsg  string `json:"error_msg"`
			ErrCode string `json:"error_code"`
		} `json:"data"`
	})
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return err
	}
	if result.Ret != "SUCCESS" {
		return fmt.Errorf("Umeng push failed, error message: %s, error code: %s",
			result.Data.ErrMsg, result.Data.ErrCode)
	}
	return nil
} // Push2MultiUsers()

/**
 * 推送给所有用户（广播）
 * 默认每天可推送10次
 */
func Push2AllUsers(platform string, payload *map[string]interface{}) error {
	url := fmt.Sprintf("%s%s", DOMAIN, "/api/send")

	if !(platform == PLATFORM_IOS || platform == PLATFORM_ANDROID) {
		return fmt.Errorf("Umeng push failed, wrong platform")
	}

	param := map[string]interface{}{
		"appkey":    configMap[platform].AppKey,
		"timestamp": time.Now().Unix(),
		"type":      "broadcast",
		"payload":   *payload,
	}

	// 签名
	sign, err := doSign("POST", url, param, configMap[platform].AppMasterSecret)
	if err != nil {
		return err
	}

	data, err := json.Marshal(param)
	if err != nil {
		return err
	}

	request := gorequest.New()
	_, body, _ := request.Post(fmt.Sprintf("%s?sign=%s", url, sign)).
		Send(string(data)).End()
	result := new(struct {
		Ret  string `json:"ret"`
		Data struct {
			MsgId   string `json:"msg_id"`
			ErrMsg  string `json:"error_msg"`
			ErrCode string `json:"error_code"`
		} `json:"data"`
	})
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return err
	}
	if result.Ret != "SUCCESS" {
		return fmt.Errorf("Umeng push failed, error message: %s, error code: %s",
			result.Data.ErrMsg, result.Data.ErrCode)
	}
	return nil
} // Push2AllUsers()
