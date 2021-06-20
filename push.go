package gomeng

/**
 * 推送给单用户（单播）
 */
func (c *Client) Push(payload map[string]interface{}, deviceToken string) error {
	resp, err := c.doPost(c.genReqParams(payload, "unicast", deviceToken), APIPush)
	if err != nil {
		return err
	}
	if err := resp.Error(); err != nil {
		return err
	}
	return nil
}

/**
 * 推送给多用户（列播）
 */
func (c *Client) ListCast(payload map[string]interface{}, deviceTokens ...string) error {
	resp, err := c.doPost(c.genReqParams(payload, "listcast", deviceTokens...), APIPush)
	if err != nil {
		return err
	}
	if err := resp.Error(); err != nil {
		return err
	}
	return nil
}

/**
 * 推送给所有用户（广播）
 * 默认每天可推送10次
 */
func (c *Client) Broadcast(payload map[string]interface{}) error {
	resp, err := c.doPost(c.genReqParams(payload, "broadcast"), APIBroadcast)
	if err != nil {
		return err
	}
	if err := resp.Error(); err != nil {
		return err
	}
	return nil
}
