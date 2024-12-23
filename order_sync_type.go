package alipay

// 二轮充电桩 充电订单信息同步 https://opendocs.alipay.com/mini/02ctgm
type MerchantOrderSync struct {
	AuxParam
	OutBizNo          string          `json:"out_biz_no"`                // 必填 外部订单号。由商家自定义。  注意：同一笔订单更新状态时，需与首次同步入参一致。
	BuyerID           string          `json:"buyer_id"`                  // 必填 支付宝用户唯一标识。新商户建议使用open_id替代该字段。对于新商户，user_id字段未来计划逐步回收，存量商户可继续使用。如使用open_id，请确认 应用-开发配置-openid配置管理 已启用。无该配置项，可查看 openid配置申请
	BuyerOpenID       string          `json:"buyer_open_id"`             //
	OrderType         string          `json:"order_type"`                // 必填 订单类型。固定为SERVICE_ORDER（服务订单）。
	OrderCreateTime   string          `json:"order_create_time"`         // 必填 订单创建时间，即该笔订单真实的创建时间。时间格式为 yyyy-MM-dd HH:mm:ss。
	OrderModifiedTime string          `json:"order_modified_time"`       // 必填 订单修改时间，时间格式为 yyyy-MM-dd HH:mm:ss.SSS，订单状态或内容发生变更时需要同步更新该时间。用于订单状态或数据变化较快的顺序控制，防止乱序。order_modified_time较晚的同步会被最终存储，order_modified_time相同的两次同步会被幂等处理。
	ServiceCode       string          `json:"service_code,omitempty"`    // 服务code：传入小程序后台提报的服务id，将订单与服务关联。service_code 通过 alipay.open.app.service.apply(服务提报申请)接口提报服务后获取。
	SourceApp         string          `json:"source_app,omitempty"`      // 订单来源，默认Alipay(支付宝小程序)，钉钉来源Dingtalk，其他来源枚举值定义联系支付宝订单中心添加
	TradeNo           string          `json:"trade_no,omitempty"`        // 支付宝交易号。注意：若该状态传递了trade_no，则同步该订单的后续状态时都需要传递trade_no。 此字段关联支付宝账单详情卡片，建议回传真实有效交易号。
	Amount            float64         `json:"amount"`                    // 必填 订单金额，单位为元。
	PayAmount         float64         `json:"pay_amount,omitempty"`      // 实际支付金额。精确到小数点后两位。单位为元
	DiscountAmount    float64         `json:"discount_amount,omitempty"` // 优惠金额。传入 discount_info_list  时必填。单位为元
	ShopInfo          ShopInfo        `json:"shop_info"`                 // 必填 门店信息。可查看下文: 门店信息:
	ExtInfo           []OrderExtInfo  `json:"ext_info"`                  // 必填 订单扩展字段。可查看下文 订单扩展字段。
	ItemOrderList     []ItemOrderList `json:"item_order_list"`           // 必填 商品信息列表。可查看下文: 商品信息列表
}

func (t MerchantOrderSync) APIName() string {
	return "alipay.merchant.order.sync"
}

func (t MerchantOrderSync) Params() map[string]string {
	var m = make(map[string]string)

	return m
}

type MerchantOrderSyncRsp struct {
	Error
	RecordID         string             `json:"record_id"`
	OrderID          string             `json:"order_id"`
	OrderStatus      string             `json:"order_status"`
	DistributeResult []DistributeResult `json:"distribute_result"`
	SyncSuggestions  []SyncSuggestions  `json:"sync_suggestions"`
}

type DistributeResult struct {
	SceneCode           string `json:"scene_code"`
	SceneName           string `json:"scene_name"`
	NotDistributeReason string `json:"not_distribute_reason"`
}
type SyncSuggestions struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ShopInfo struct {
	Address        string `json:"address"`
	Name           string `json:"name"`
	MerchantShopID string `json:"merchant_shop_id"`
}
type OrderExtInfo struct {
	ExtKey   string `json:"ext_key"`
	ExtValue string `json:"ext_value"`
}
type ItemOrderList struct {
	ItemName string         `json:"item_name"`
	ExtInfo  []OrderExtInfo `json:"ext_info"`
}

type BusinessInfo struct {
	PileId            string `json:"pile_id,omitempty"`         // 否	否	否	字段名称:充电桩ID字段说明:运营商体系内充电桩的唯一标识ID，行业有要求回传时必填。
	ChargingStartTime string `json:"charging_start_time"`       // 是	否	否	字段名称:充电开始时间字段说明:系统显示充电开始时间点,时间格式为 yyyy-MM-dd HH:mm:ss。
	ChargingEndTime   string `json:"charging_end_time"`         // 是	否	是	字段名称:充电结束时间字段说明:系统显示充电结束时间点,时间格式为 yyyy-MM-dd HH:mm:ss。
	AbnormalReason    string `json:"abnormal_reason,omitempty"` // 否	否	否	字段名称:充电异常原因字段说明:充电异常时的原因
	Number            string `json:"number"`                    // 是	否	否	字段名称:充电插座号字段说明:充电插座的编号
	SocketNum         string `json:"socket_num,omitempty"`      // 否	否	否	字段名称:充电桩插座数字段说明:充电桩包含插座数，行业有要求时必填
	RefundAmount      string `json:"refund_amount,omitempty"`   // 否	否	否	字段名称:退款金额字段说明:退款金额，单位为元
	ProviderName      string `json:"provider_name"`             // 是	否	是	字段名称:服务提供方字段说明:服务商或商家名称
	ChargingTime      string `json:"charging_time"`             // 是	否	是	字段名称:充电时长字段说明:充电总时长，为数字，单位为分钟
	EndReason         string `json:"end_reason"`                // 是	否	否	字段名称:充电结束原因字段说明:充电结束的原因
	AbnormalTime      string `json:"abnormal_time,omitempty"`   // 否	否	否	字段名称:充电异常发生时间字段说明:充电发生异常的时间点,时间格式为 yyyy-MM-dd HH:mm:ss。
}
