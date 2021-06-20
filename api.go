package gomeng

import (
	"encoding/json"
	"fmt"
)

const (
	BaseURL = "https://msgapi.umeng.com/api/"

	APIPush      = "send"
	APIBroadcast = "send"
)

type ResponseMessage struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgID   string `json:"msg_id"`
		TaskID  string `json:"task_id"`
		ErrMsg  string `json:"error_msg"`
		ErrCode string `json:"error_code"`
	} `json:"data"`
}

func (rm *ResponseMessage) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, rm); err != nil {
		return err
	}
	return nil
}

func (rm *ResponseMessage) Error() error {
	if rm.Ret != "SUCCESS" {
		return fmt.Errorf("Umeng push failed, error message: %s, error code: %s",
			rm.Data.ErrMsg, rm.Data.ErrCode)
	}
	return nil
}
