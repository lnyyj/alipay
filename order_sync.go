package alipay

import "context"

// MerchantOrderSync 推送订单 https://opendocs.alipay.com/mini/043zb5
func (c *Client) MerchantOrderSync(ctx context.Context, param MerchantOrderSync) (result *MerchantOrderSyncRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}
