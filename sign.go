package gomeng

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

/**
 * 友盟签名算法
 * 拼接请求方法、url、post-body及应用的app_master_secret
 * 将D形成字符串计算MD5值，形成一个32位的十六进制（字母小写）字符串，即为本次请求sign（签名）的值
 */
func (c *Client) doSign(method string, url string, param map[string]interface{}) (string, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("%s%s%s%s", method, url, string(body), c.AppMasterSecret)
	hasher := md5.New()
	_, err = hasher.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
