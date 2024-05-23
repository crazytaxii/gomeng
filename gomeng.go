package gomeng

import (
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

type Config struct {
	AppKey         string        `json:"app_key" yaml:"appKey"`
	AppSecret      string        `json:"app_secret" yaml:"appSecret"`
	ProductionMode bool          `json:"production_mode" yaml:"productionMode"`
	Timeout        time.Duration `json:"timeout" yaml:"timeout"`
}

type Client struct {
	cfg     *Config
	rawhttp *http.Client
}

func NewClient(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
		rawhttp: &http.Client{
			Timeout: fallback2DefaultIfZero(cfg.Timeout),
		},
	}
}
