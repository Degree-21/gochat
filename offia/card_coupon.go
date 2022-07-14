/*
@Time : 2022/7/14 3:36 PM
@Author : 21
@File : card_coupon
@Software: GoLand
*/
package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// create card coupon details
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html
type RequestCreateCard struct {
	Card *Card `json:"card"`
}

// card 结构体
type Card struct {
	CardType string    `json:"card_type"`
	BaseInfo *BaseInfo `json:"base_info"`
	Groupon  *Groupon  `json:"groupon"`
}

type Groupon struct {
	AdvancedInfo struct {
		UseCondition struct {
			AcceptCategory          string `json:"accept_category"`
			RejectCategory          string `json:"reject_category"`
			CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount"`
		} `json:"use_condition"`
		Abstract struct {
			Abstract    string   `json:"abstract"`
			IconURLList []string `json:"icon_url_list"`
		} `json:"abstract"`
		TextImageList []struct {
			ImageURL string `json:"image_url"`
			Text     string `json:"text"`
		} `json:"text_image_list"`
		TimeLimit []struct {
			Type        string `json:"type"`
			BeginHour   int    `json:"begin_hour,omitempty"`
			EndHour     int    `json:"end_hour,omitempty"`
			BeginMinute int    `json:"begin_minute,omitempty"`
			EndMinute   int    `json:"end_minute,omitempty"`
		} `json:"time_limit"`
		BusinessService []string `json:"business_service"`
	} `json:"advanced_info"`
	DealDetail string `json:"deal_detail"`
}

type BaseInfo struct {
	LogoURL      string `json:"logo_url"`
	BrandName    string `json:"brand_name"`
	CodeType     string `json:"code_type"`
	Title        string `json:"title"`
	Color        string `json:"color"`
	Notice       string `json:"notice"`
	ServicePhone string `json:"service_phone"`
	Description  string `json:"description"`
	DateInfo     struct {
		Type           string `json:"type"`
		BeginTimestamp int    `json:"begin_timestamp"`
		EndTimestamp   int    `json:"end_timestamp"`
	} `json:"date_info"`
	Sku struct {
		Quantity int `json:"quantity"`
	} `json:"sku"`
	UseLimit          int    `json:"use_limit"`
	GetLimit          int    `json:"get_limit"`
	UseCustomCode     bool   `json:"use_custom_code"`
	BindOpenid        bool   `json:"bind_openid"`
	CanShare          bool   `json:"can_share"`
	CanGiveFriend     bool   `json:"can_give_friend"`
	LocationIDList    []int  `json:"location_id_list"`
	CenterTitle       string `json:"center_title"`
	CenterSubTitle    string `json:"center_sub_title"`
	CenterURL         string `json:"center_url"`
	CustomURLName     string `json:"custom_url_name"`
	CustomURL         string `json:"custom_url"`
	CustomURLSubTitle string `json:"custom_url_sub_title"`
	PromotionURLName  string `json:"promotion_url_name"`
	PromotionURL      string `json:"promotion_url"`
	Source            string `json:"source"`
}

type RespCardCard struct {
	CardId string `json:"card_id"`
}

// 投放卡券
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html
type RequestCreateCardCouponQrCode struct {
	ActionName    string `json:"action_name"`
	ExpireSeconds int    `json:"expire_seconds"`
	ActionInfo    struct {
		Card struct {
			CardId       string `json:"card_id"`
			Code         string `json:"code"`
			Openid       string `json:"openid"`
			IsUniqueCode bool   `json:"is_unique_code"`
			OuterStr     string `json:"outer_str"`
		} `json:"card"`
	} `json:"action_info"`
}

// 投放卡券返回值
type RespCreateCardCouponQrCode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	ShowQrcodeUrl string `json:"show_qrcode_url"`
}

//  查询 code 接口
//   https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
type RequestCardCodeGet struct {
	CardId       string `json:"card_id"`
	Code         string `json:"code"`
	CheckConsume bool   `json:"check_consume"`
}

// 查询 code 结果
type RespCardCodeGet struct {
	Card struct {
		CardId    string `json:"card_id"`
		BeginTime int    `json:"begin_time"`
		EndTime   int    `json:"end_time"`
	} `json:"card"`
	Openid         string `json:"openid"`
	CanConsume     bool   `json:"can_consume"`
	UserCardStatus string `json:"user_card_status"`
}

// 核销 code 的请求体
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
type RequestCardCodeConsume struct {
	Code   string `json:"code"`
	CardId string `json:"card_id"`
}

// 核销 code 的结构
type RespCardCodeConsume struct {
	Card struct {
		CardId string `json:"card_id"`
	} `json:"card"`
	Openid string `json:"openid"`
}

// 用户已领取卡券接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
type RequestUserCardList struct {
	Openid string `json:"openid"`
	CardId string `json:"card_id"`
}

// 用户已领取卡券接口 返回结构体
type RespUserCardList struct {
	CardList []struct {
		Code   string `json:"code"`
		CardId string `json:"card_id"`
	} `json:"card_list"`
	HasShareCard bool `json:"has_share_card"`
}

// 卡券详情
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
type RequestCardGet struct {
	CardId string `json:"card_id"`
}

// 卡券详情返回数据
type RespCardGet struct {
	Card *CardGetCard `json:"card"`
}

type CardGetDiscount struct {
	BaseInfo struct {
		Id        string `json:"id"`
		LogoUrl   string `json:"logo_url"`
		CodeType  string `json:"code_type"`
		BrandName string `json:"brand_name"`
		Title     string `json:"title"`
		DateInfo  struct {
			Type           string `json:"type"`
			FixedTerm      int    `json:"fixed_term"`
			FixedBeginTerm int    `json:"fixed_begin_term"`
		} `json:"date_info"`
		Color          string `json:"color"`
		Notice         string `json:"notice"`
		Description    string `json:"description"`
		LocationIdList []int  `json:"location_id_list"`
		GetLimit       int    `json:"get_limit"`
		CanShare       bool   `json:"can_share"`
		CanGiveFriend  bool   `json:"can_give_friend"`
		Status         string `json:"status"`
		Sku            struct {
			Quantity      int `json:"quantity"`
			TotalQuantity int `json:"total_quantity"`
		} `json:"sku"`
		CreateTime   int           `json:"create_time"`
		UpdateTime   int           `json:"update_time"`
		AreaCodeList []interface{} `json:"area_code_list"`
	} `json:"base_info"`
	Discount     int `json:"discount"`
	AdvancedInfo struct {
		TimeLimit []struct {
			Type string `json:"type"`
		} `json:"time_limit"`
		TextImageList        []interface{} `json:"text_image_list"`
		BusinessService      []interface{} `json:"business_service"`
		ConsumeShareCardList []interface{} `json:"consume_share_card_list"`
		Abstract             struct {
			Abstract    string   `json:"abstract"`
			IconUrlList []string `json:"icon_url_list"`
		} `json:"abstract"`
		ShareFriends bool `json:"share_friends"`
	} `json:"advanced_info"`
}

type CardGetCard struct {
	CardType string           `json:"card_type"`
	Discount *CardGetDiscount `json:"discount"`
}

// 创建卡券
func CreateCardCoupon(request *RequestCreateCard, result *RespCardCard) wx.Action {
	return wx.NewPostAction(urls.CardCreate,
		//wx.WithQuery("access_token", accessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 投放卡券
func CreateCardCouponQrCode(request *RequestCreateCardCouponQrCode, result *RespCreateCardCouponQrCode) wx.Action {
	return wx.NewPostAction(urls.CardQrcodeCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 查询code 接口
func GetCardCode(request *RequestCardCodeGet, result *RespCardCodeGet) wx.Action {
	return wx.NewPostAction(urls.CardCodeGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 核销 code 接口
func ConsumeCardCode(request *RequestCardCodeConsume, result *RespCardCodeConsume) wx.Action {
	return wx.NewPostAction(urls.CardCodeConsume,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 获取用户获取的卡券
func GetUserCardList(request *RequestUserCardList, result *RespUserCardList) wx.Action {
	return wx.NewPostAction(urls.CardUserList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 获取卡券详情
func GetCard(request *RequestCardGet, result *RespCardGet) wx.Action {
	return wx.NewPostAction(urls.CardGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
