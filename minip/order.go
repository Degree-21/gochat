package minip

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"time"
)

//https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%9B%9B%E3%80%81%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E5%88%97%E8%A1%A8

// GetOrderList 获取订单列表
type ParamsGetOrderList struct {
	PayTimeRange *PayTimeRange `json:"pay_time_range"`        // 支付时间
	OrderState   int           `json:"order_state,omitempty"` // 订单状态枚举 订单状态枚举：(1) 待发货；(2) 已发货；(3) 确认收货；(4) 交易完成；(5) 已退款。
	OpenID       string        `json:"openid,omitempty"`      // 支付者openid
	LastIndex    string        `json:"last_index,omitempty"`  // 翻页时使用的最后索引
	PageSize     int           `json:"page_size,omitempty"`   // 翻页时使用，返回列表的长度
}

type PayTimeRange struct {
	BeginTime int64 `json:"begin_time,omitempty"` // 起始时间，时间戳形式
	EndTime   int64 `json:"end_time,omitempty"`   // 结束时间（含），时间戳形式
}

type ResultGetOrderList struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`

	LastIndex string              `json:"last_index"`           // 翻页时使用。
	HasMore   bool                `json:"has_more"`             // 是否有下一页
	OrderList []*GetOrderListData `json:"order_list,omitempty"` //订单列表
}

type GetOrderListData struct {
	TransactionID   string    `json:"transaction_id,omitempty"`    // 原支付交易对应的微信订单号
	MerchantID      string    `json:"merchant_id,omitempty"`       // 商户号
	SubMerchantID   string    `json:"sub_merchant_id,omitempty"`   // 二级商户号
	MerchantTradeNo string    `json:"merchant_trade_no,omitempty"` // 商户内部订单号
	Description     string    `json:"description,omitempty"`       // 商品描述
	PaidAmount      int64     `json:"paid_amount,omitempty"`       // 支付金额
	OpenID          string    `json:"openid,omitempty"`            // 支付者openid
	TradeCreateTime int64     `json:"trade_create_time,omitempty"` // 交易创建时间
	PayTime         int64     `json:"pay_time,omitempty"`          // 支付时间
	OrderState      int       `json:"order_state,omitempty"`       // 订单状态
	InComplaint     bool      `json:"in_complaint,omitempty"`      // 是否处在交易纠纷中
	Shipping        *Shipping `json:"shipping,omitempty"`          // 发货信息
}

type Shipping struct {
	DeliveryMode        int             `json:"delivery_mode,omitempty"`         // 发货模式
	LogisticsType       int             `json:"logistics_type,omitempty"`        // 物流模式
	FinishShipping      bool            `json:"finish_shipping,omitempty"`       // 是否已完成全部发货
	GoodsDesc           string          `json:"goods_desc,omitempty"`            // 商品描述
	FinishShippingCount int             `json:"finish_shipping_count,omitempty"` // 完成发货次数
	ShippingList        []*ShippingInfo `json:"shipping_list,omitempty"`         // 物流信息列表
}

type ShippingInfo struct {
	TrackingNo     string   `json:"tracking_no,omitempty"`     // 物流单号
	ExpressCompany string   `json:"express_company,omitempty"` // 物流公司编码
	GoodsDesc      string   `json:"goods_desc,omitempty"`      // 商品描述
	UploadTime     int64    `json:"upload_time,omitempty"`     // 上传时间
	ItemDesc       string   `json:"item_desc"`
	Contact        *Contact `json:"contact,omitempty"`
}

type Contact struct {
	ConsignorContact string `json:"consignor_contact,omitempty"` // 寄件人联系方式
	ReceiverContact  string `json:"receiver_contact,omitempty"`  // 收件人联系方式
}

type ParamsIsTradeManaged struct {
	Appid string `json:"appid"`
}

// Order represents the order information for logistics.

type ParamsUploadShippingOrder struct {
	OrderKey       *OrderKey       `json:"order_key"`
	LogisticsType  int             `json:"logistics_type"`             //物流模式，发货方式枚举值：1、实体物流配送采用快递公司进行实体物流配送形式 2、同城配送 3、虚拟商品，虚拟商品，例如话费充值，点卡等，无实体配送形式 4、用户自提
	DeliveryMode   int             `json:"delivery_mode"`              //发货模式，发货模式枚举值：1、UNIFIED_DELIVERY（统一发货）2、SPLIT_DELIVERY（分拆发货） 示例值: UNIFIED_DELIVERY
	IsAllDelivered bool            `json:"is_all_delivered,omitempty"` //分拆发货模式时必填，用于标识分拆发货模式下是否已全部发货完成，只有全部发货完成的情况下才会向用户推送发货完成通知。示例值: true/false
	ShippingList   []*ShippingInfo `json:"shipping_list"`              // 物流信息列表，发货物流单列表，支持统一发货（单个物流单）和分拆发货（多个物流单）两种模式，多重性: [1, 10]
	UploadTime     time.Time       `json:"upload_time"`                //上传时间，用于标识请求的先后顺序 示例值: `2022-12-15T13:29:35.120+08:00`
	Payer          *Payer          `json:"payer"`                      //	支付者，支付者信息
}
type OrderKey struct {
	OrderNumberType int    `json:"order_number_type"`        // 订单单号类型，用于确认需要上传详情的订单。枚举值1，使用下单商户号和商户侧单号；枚举值2，使用微信支付单号。
	TransactionId   string `json:"transaction_id,omitempty"` // 原支付交易对应的微信订单号
	MchId           string `json:"mch_id,omitempty"`         //支付下单商户的商户号，由微信支付生成并下发。
	OutTradeNo      string `json:"out_trade_no,omitempty"`   //支付下单商户的商户号，由微信支付生成并下发。
}

// Payer represents the payer information.
type Payer struct {
	OpenID string `json:"openid"` //用户标识，用户在小程序appid下的唯一标识。 下单前需获取到用户的Openid 示例值: oUpF8uMuAJO_M2pxb1Q9zNjWeS6o 字符字节限制: [1, 128]
}

type ResultUploadShippingOrder struct {
}

type ResultTradeManaged struct {
	IsTradeManaged bool `json:"is_trade_managed"`
}

func GetOrderList(params *ParamsGetOrderList, result *ResultGetOrderList) wx.Action {
	return wx.NewPostAction(urls.MinipOrderGetOrderList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UploadShippingOrder(params *ParamsUploadShippingOrder, result *ResultUploadShippingOrder) wx.Action {
	return wx.NewPostAction(urls.MinipOrderUploadShippingInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func IsTradeManaged(params *ParamsIsTradeManaged, result *ResultTradeManaged) wx.Action {
	return wx.NewPostAction(urls.MinipOrderIsTradeManaged,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
