package gomeng

import ()

const DOMAIN = "msgapi.umeng.com"

type Client struct {
	ProductMode     bool
	Platform        string
	AppKey          string
	AppMasterSecret string
}

func NewClient(productMode bool, appKey string, appMasterSecret string) *Client {
	return &Client{
		ProductMode:     productMode,
		AppKey:          appKey,
		AppMasterSecret: appMasterSecret,
	}
} // NewClient()
