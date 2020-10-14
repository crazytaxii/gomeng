package gomeng

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func (c *Client) Request(method, url string, param map[string]interface{}) error {
	// sign
	sign, err := c.doSign(method, url, param)
	if err != nil {
		return fmt.Errorf("Sign failed, error: %s", err.Error())
	}

	bytes, err := json.Marshal(sign)
	if err != nil {
		return fmt.Errorf("JSON marshal failed, error: %s", err.Error())
	}

	_, body, errs := gorequest.New().Post(fmt.Sprintf("%s?sign=%s", url, sign)).Send(string(bytes)).End()
	if len(errs) > 0 {
		return fmt.Errorf("HTTP request failed, error: %s", errs[0])
	}
	res := &ResponseMessage{}
	err = json.Unmarshal([]byte(body), res)
	if err != nil {
		return fmt.Errorf("JSON unmarshal failed, error: %s", err.Error())
	}
	if res.Ret != "SUCCESS" {
		return fmt.Errorf("Umeng push failed, error message: %s, error code: %d",
			res.Data.ErrMsg, res.Data.ErrCode)
	}

	return nil
}
