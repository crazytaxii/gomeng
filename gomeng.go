package gomeng

import (
	"net/http"
	"time"
)

const DefaultTimeout = 10 * time.Second

type Client struct {
	productMode bool
	key         string
	secret      string
	*http.Client
}

func NewClient(productMode bool, key, secret string, timeout time.Duration) *Client {
	return &Client{
		productMode: productMode,
		key:         key,
		secret:      secret,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}
