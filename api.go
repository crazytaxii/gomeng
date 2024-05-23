package gomeng

import "fmt"

const (
	baseURL = "https://msgapi.umeng.com"

	apiPush      = "api/send"
	apiBroadcast = "api/send"
)

type requestType string

const (
	unicastRequest   requestType = "unicast"
	listcastRequest  requestType = "listcast"
	broadcastRequest requestType = "broadcast"
)

type ReturnState string

const (
	SuccessState ReturnState = "SUCCESS"
	FailState    ReturnState = "FAIL"
)

type (
	Data struct {
		MessageID  string `json:"msg_id"`
		TaskID     string `json:"task_id"`
		ErrMessage string `json:"error_msg"`
		ErrCode    string `json:"error_code"`
	}
	ResponseMessage struct {
		Ret  ReturnState `json:"ret"`
		Data `json:"data"`
	}
)

func (rm *ResponseMessage) Error() error {
	if rm.Ret == SuccessState {
		return nil
	}
	return fmt.Errorf("error %s: %s", rm.ErrCode, rm.ErrMessage)
}
