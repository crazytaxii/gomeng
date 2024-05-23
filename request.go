package gomeng

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Payload map[string]interface{}

func (c *Client) genRequestParams(payload Payload, reqType requestType, deviceTokens ...string) map[string]interface{} {
	p := map[string]interface{}{
		"appkey":          c.cfg.AppKey,
		"timestamp":       time.Now().Unix(),
		"type":            reqType,
		"payload":         payload,
		"production_mode": c.cfg.ProductionMode,
	}
	if len(deviceTokens) > 0 {
		p["device_tokens"] = strings.Join(deviceTokens, ",")
	}
	return p
}

func (c *Client) doPost(ctx context.Context, param map[string]interface{}, endpoint string) (*ResponseMessage, error) {
	// e.g. https://msgapi.umeng.com/api/send
	url, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	// do sign
	sign, err := sign(http.MethodPost, url, c.cfg.AppSecret, b)
	if err != nil {
		return nil, fmt.Errorf("sign error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, joinSign(url, sign), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.rawhttp.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rm := &ResponseMessage{}
	if err := json.Unmarshal(body, rm); err != nil {
		return nil, err
	}
	return rm, nil
}
