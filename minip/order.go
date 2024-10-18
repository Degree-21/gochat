package minip

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

//https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%9B%9B%E3%80%81%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E5%88%97%E8%A1%A8

// GetOrderList 获取订单列表
type ParamsGetOrderList struct {
	BeginTime  int64  `json:"begin_time,omitempty"`  // 起始时间，时间戳形式
	EndTime    int64  `json:"end_time,omitempty"`    // 结束时间（含），时间戳形式
	OrderState int    `json:"order_state,omitempty"` // 订单状态枚举
	OpenID     string `json:"openid,omitempty"`      // 支付者openid
	LastIndex  string `json:"last_index,omitempty"`  // 翻页时使用的最后索引
	PageSize   int    `json:"page_size,omitempty"`   // 翻页时使用，返回列表的长度
}

type ResultGetOrderList struct {
	LastIndex string              `json:"last_index"` // 翻页时使用。
	HasMore   bool                `json:"has_more"`   // 是否有下一页
	OrderList []*GetOrderListData `json:"order_list"` //订单列表
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
	Contact        *Contact `json:"contact"`                   //联系方式。
}

type Contact struct {
	ConsignorContact string `json:"consignor_contact,omitempty"` // 寄件人联系方式
	ReceiverContact  string `json:"receiver_contact,omitempty"`  // 收件人联系方式
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
