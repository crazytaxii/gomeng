package gomeng

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func sign(method, url, secret string, raw []byte) (string, error) {
	buf := bytes.NewBufferString(method)
	buf.WriteString(url)
	buf.Write(raw)
	buf.WriteString(secret)
	hasher := md5.New()
	if _, err := hasher.Write(buf.Bytes()); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func joinSign(url, sign string) string {
	// e.g. https://msgapi.umeng.com/api/send?sign=xxx
	return fmt.Sprintf("%s?sign=%s", url, sign)
}

func fallback2DefaultIfZero(timeout time.Duration) time.Duration {
	if timeout > 0 {
		return timeout
	}
	return defaultTimeout
}
