package gomeng

const (
	BaseURL = "https://msgapi.umeng.com/api"

	Push      = "send"
	Broadcast = "send"
)

type ResponseMessage struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgID   string `json:"msg_id"`
		TaskID  string `json:"task_id"`
		ErrMsg  string `json:"error_msg"`
		ErrCode int    `json:"error_code"`
	} `json:"data"`
}
