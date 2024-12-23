package alipay_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/lnyyj/alipay"
)

func TestClient_MerchantOrderSyncCreate(t *testing.T) {
	nowt := time.Now().Format("2006-01-02 15:04:05")
	bizInfo := alipay.BusinessInfo{
		PileId:            "pile001",
		ChargingStartTime: nowt,
		ChargingEndTime:   nowt,
		AbnormalReason:    "",
		Number:            "1",
		SocketNum:         "2",
		RefundAmount:      "12",
		ProviderName:      "能充电",
		ChargingTime:      "1",
		EndReason:         "",
		AbnormalTime:      nowt,
	}
	bizInfoBuf, _ := json.Marshal(bizInfo)

	payload := alipay.MerchantOrderSync{
		OutBizNo:          "test0001",
		BuyerID:           "",
		BuyerOpenID:       "",
		OrderType:         "SERVICE_ORDER",
		OrderCreateTime:   nowt,
		OrderModifiedTime: nowt,
		Amount:            0,
		PayAmount:         0,
		DiscountAmount:    0,
		ShopInfo: alipay.ShopInfo{
			Address:        "蚂蚁A空间P5停车库",
			Name:           "猛犸文三路站",
			MerchantShopID: "shop0001",
		},
		ExtInfo: []alipay.OrderExtInfo{
			alipay.OrderExtInfo{
				ExtKey:   "business_info",
				ExtValue: string(bizInfoBuf),
			},
			alipay.OrderExtInfo{
				ExtKey:   "merchant_biz_type",
				ExtValue: "TWO_WHEELED_CHARGING_PILE",
			},
			alipay.OrderExtInfo{
				ExtKey:   "merchant_order_status",
				ExtValue: "CREATE",
			},
			alipay.OrderExtInfo{
				ExtKey:   "merchant_order_status",
				ExtValue: "CREATE",
			},
		},
		ItemOrderList: []alipay.ItemOrderList{},
	}

	t.Log(payload)
	result, err := client.MerchantOrderSync(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	if result.IsFailure() {
		t.Fatal(result.Msg, result.SubMsg)
	}
	t.Logf("%v", result)
}
