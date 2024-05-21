package gomeng

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (c *Client) genReqParams(payload map[string]interface{}, reqType string, deviceTokens ...string) map[string]interface{} {
	p := map[string]interface{}{
		"appkey":          c.key,
		"timestamp":       time.Now().Unix(),
		"type":            reqType,
		"payload":         payload,
		"production_mode": c.productMode,
	}
	if len(deviceTokens) > 0 {
		p["device_tokens"] = strings.Join(deviceTokens, ",")
	}
	return p
}

func (c *Client) doPost(param map[string]interface{}, endpoint string) (*ResponseMessage, error) {
	url := BaseURL + endpoint
	// sign
	sign, err := c.doSign(http.MethodPost, url, param)
	if err != nil {
		return nil, fmt.Errorf("Sign failed, error: %s", err.Error())
	}

	b, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	resp, err := c.Post(fmt.Sprintf("%s?sign=%s", url, sign), "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rm := &ResponseMessage{}
	if err := rm.Unmarshal(body); err != nil {
		return nil, err
	}
	return rm, nil
}
