package gomeng

import "context"

// 推送给单用户（单播）
func (c *Client) Unicast(ctx context.Context, payload Payload, deviceToken string) (resp *ResponseMessage, err error) {
	return c.doPost(ctx, c.genRequestParams(payload, unicastRequest, deviceToken), apiPush)
}

// 推送给多用户（列播）
func (c *Client) ListCast(ctx context.Context, payload Payload, deviceTokens ...string) (resp *ResponseMessage, err error) {
	return c.doPost(ctx, c.genRequestParams(payload, listcastRequest, deviceTokens...), apiPush)
}

// 推送给所有用户（广播）
// 默认每天只可推送10次
func (c *Client) Broadcast(ctx context.Context, payload Payload) (resp *ResponseMessage, err error) {
	return c.doPost(ctx, c.genRequestParams(payload, broadcastRequest), apiBroadcast)
}
