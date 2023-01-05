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
	CardType   string      `json:"card_type"`
	BaseInfo   *BaseInfo   `json:"base_info,omitempty"`
	Groupon    *Groupon    `json:"groupon,omitempty"`
	MemberCard *MemberCard `json:"member_card,omitempty"`
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

type MemberCard struct {
	ActivateURL      string       `json:"activate_url"`
	AdvancedInfo     AdvancedInfo `json:"advanced_info"`
	AutoActivate     bool         `json:"auto_activate"`
	BackgroundPicURL string       `json:"background_pic_url"`
	BaseInfo         BaseInfo     `json:"base_info"`
	BonusRule        BonusRule    `json:"bonus_rule"`
	CustomCell1      CustomCell1  `json:"custom_cell1"`
	CustomCell2      CustomCell1  `json:"custom_cell2"`
	CustomField1     CustomField1 `json:"custom_field1"`
	CustomField2     CustomField1 `json:"custom_field2"`
	Discount         int64        `json:"discount"`
	Prerogative      string       `json:"prerogative"`
	SupplyBalance    bool         `json:"supply_balance"`
	SupplyBonus      bool         `json:"supply_bonus"`
	BonusCleared     string       `json:"bonus_cleared"`
	BonusRules       string       `json:"bonus_rules"`
}

type AdvancedInfo struct {
	Abstract        Abstract        `json:"abstract"`
	BusinessService []string        `json:"business_service"`
	TextImageList   []TextImageList `json:"text_image_list"`
	TimeLimit       []TimeLimit     `json:"time_limit"`
	UseCondition    UseCondition    `json:"use_condition"`
}

type Abstract struct {
	Abstract    string   `json:"abstract"`
	IconURLList []string `json:"icon_url_list"`
}

type TextImageList struct {
	ImageURL string `json:"image_url"`
	Text     string `json:"text"`
}

type TimeLimit struct {
	BeginHour   *int64 `json:"begin_hour,omitempty"`
	BeginMinute *int64 `json:"begin_minute,omitempty"`
	EndHour     *int64 `json:"end_hour,omitempty"`
	EndMinute   *int64 `json:"end_minute,omitempty"`
	Type        string `json:"type"`
}

type UseCondition struct {
	AcceptCategory          string `json:"accept_category"`
	CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount"`
	RejectCategory          string `json:"reject_category"`
}

type BonusRule struct {
	CostBonusUnit        int64 `json:"cost_bonus_unit"`
	CostMoneyUnit        int64 `json:"cost_money_unit"`
	IncreaseBonus        int64 `json:"increase_bonus"`
	InitIncreaseBonus    int64 `json:"init_increase_bonus"`
	LeastMoneyToUseBonus int64 `json:"least_money_to_use_bonus"`
	MaxIncreaseBonus     int64 `json:"max_increase_bonus"`
	MaxReduceBonus       int64 `json:"max_reduce_bonus"`
	ReduceMoney          int64 `json:"reduce_money"`
}

type CustomCell1 struct {
	Name             string `json:"name"`
	Tips             string `json:"tips"`
	URL              string `json:"url"`
	AppBrandPass     string `json:"app_brand_pass"`
	AppBrandUserName string `json:"app_brand_user_name"`
}

type CustomField1 struct {
	NameType         string `json:"name_type"`
	Name             string `json:"name"`
	URL              string `json:"url"`
	AppBrandPass     string `json:"app_brand_pass"`
	AppBrandUserName string `json:"app_brand_user_name"`
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
		BeginTimestamp int    `json:"begin_timestamp,omitempty"`
		EndTimestamp   int    `json:"end_timestamp,omitempty"`
	} `json:"date_info"`
	Sku struct {
		Quantity int `json:"quantity"`
	} `json:"sku"`
	UseLimit                  int    `json:"use_limit"`
	GetLimit                  int    `json:"get_limit"` //每人领取数量限制
	UseCustomCode             bool   `json:"use_custom_code"`
	BindOpenid                bool   `json:"bind_openid"`
	CanShare                  bool   `json:"can_share"`
	CanGiveFriend             bool   `json:"can_give_friend"`
	LocationIDList            []int  `json:"location_id_list"`
	CenterTitle               string `json:"center_title"`
	CenterSubTitle            string `json:"center_sub_title"`
	CenterURL                 string `json:"center_url"`
	CustomURLName             string `json:"custom_url_name"`
	CustomURL                 string `json:"custom_url"`
	CustomURLSubTitle         string `json:"custom_url_sub_title"`
	CustomAppBrandUserName    string `json:"custom_app_brand_user_name"`
	CustomAppBrandPass        string `json:"custom_app_brand_pass"`
	PromotionURLName          string `json:"promotion_url_name"`
	PromotionURL              string `json:"promotion_url"`
	PromotionAppBrandUserName string `json:"promotion_app_brand_user_name"`
	PromotionAppBrandPass     string `json:"promotion_app_brand_pass"`
	PromotionUrlSubTitle      string `json:"promotion_url_sub_title"`
	Source                    string `json:"source"`
	CenterAppBrandUserName    string `json:"center_app_brand_user_name"`
	CenterAppBrandPass        string `json:"center_app_brand_pass"`
}

type RespCardCard struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	CardId  string `json:"card_id"`
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

type RequestSetCardWhitelist struct {
	Openid   []string `json:"openid"`
	Username []string `json:"username"`
}

type RespSetCardWhitelist struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
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

// 批量查询卡券列表
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
type RequestCardBatchGet struct {
	Offset     int      `json:"offset"`
	Count      int      `json:"count"`
	StatusList []string `json:"status_list"`
}

// 批量查询卡券列表结果
type RespCardBatchGet struct {
	CardIdList []string `json:"card_id_list"`
	TotalNum   int      `json:"total_num"`
}

// 更改卡券信息
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
type RequestCardUpdate struct {
	CardId     string      `json:"card_id"`
	MemberCard *MemberCard `json:"member_card"`
}

// 更改卡券信息 结果
type RespCardUpdate struct {
	SendCheck bool `json:"send_check"`
}

// 修改库存
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#6
type RequestCardModifySocket struct {
	CardId             string `json:"card_id"`
	IncreaseStockValue int    `json:"increase_stock_value"`
	ReduceStockValue   int    `json:"reduce_stock_value"`
}

// 修改库存返回结构体 nil

// 修改卡券code
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#6
type RequestCardCodeUpdate struct {
	Code    string `json:"code"`
	CardId  string `json:"card_id"`
	NewCode string `json:"new_code"`
}

// 修改卡券code nil

// 删除卡券
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#6
type RequestCardDelete struct {
	CardId string `json:"card_id"`
}

// 删除卡券返回结构体 nil

// 设置卡券失效接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#6
//设置卡券失效接口 非自定义结构体
type RequestCardCodeUnavailable struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

//设置卡券失效接口 自定义结构体
type RequestCardCodeUnavailableAuto struct {
	Code   string `json:"code"`
	CardId string `json:"cardId"`
}

// 统计卡券概况数据接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#6
type RequestGetCardBizUinInfo struct {
	BeginDate  string `json:"begin_date"`
	EndDate    string `json:"end_date"`
	CondSource int    `json:"cond_source"`
}

type RequestCardCodeDecrypt struct {
	EncryptCode string `json:"encrypt_code"`
}

type RespCardCodeDecrypt struct {
	Code string `json:"code"`
}

// 统计卡券概况数据接口 返回数据结构体
type RespGetCardBizUinInfo struct {
	List *GetCardBizUinInfoList `json:"list"`
}

type GetCardBizUinInfoList struct {
	RefDate     string `json:"ref_date"`
	ViewCnt     int    `json:"view_cnt"`
	ViewUser    int    `json:"view_user"`
	ReceiveCnt  int    `json:"receive_cnt"`
	ReceiveUser int    `json:"receive_user"`
	VerifyCnt   int    `json:"verify_cnt"`
	VerifyUser  int    `json:"verify_user"`
	GivenCnt    int    `json:"given_cnt"`
	GivenUser   int    `json:"given_user"`
	ExpireCnt   int    `json:"expire_cnt"`
	ExpireUser  int    `json:"expire_user"`
}

type LandingPageCreateRequest struct {
	Banner   string     `json:"banner"`
	CanShare bool       `json:"can_share"`
	CardList []CardList `json:"card_list"`
	Scene    string     `json:"scene"`
	Title    string     `json:"title"`
}

type CardList struct {
	CardID   string `json:"card_id"`
	ThumbURL string `json:"thumb_url"`
}

type LandingPageCreateResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	PageID  int64  `json:"page_id"`
	URL     string `json:"url"`
}

// 创建卡券
func CreateCardCoupon(request *RequestCreateCard, result *RespCardCard) wx.Action {
	return wx.NewPostAction(urls.CardCreate,
		//wx.WithQuery("access_token", accessToken),
		wx.WithBody(func() ([]byte, error) {
			//_, _ := json.Marshal(request)
			//fmt.Println(string(str))
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

//设置测试白名单
func SetCardWhitelist(request *RequestSetCardWhitelist, result *RespSetCardWhitelist) wx.Action {
	return wx.NewPostAction(urls.CardWhitelist,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
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

// 批量获取卡券
func GetBatchCardList(request *RequestCardBatchGet, result *RespCardBatchGet) wx.Action {
	return wx.NewPostAction(urls.CardBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 更新卡券信息
func UpdateCard(request *RequestCardUpdate, result *RespCardUpdate) wx.Action {
	return wx.NewPostAction(urls.CardUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 修改卡券库存接口
func CardModifySocket(request *RequestCardModifySocket) wx.Action {
	return wx.NewPostAction(urls.CardModifyStock,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
	)
}

// 修改卡券code
func CardCodeUpdate(request *RequestCardCodeUpdate) wx.Action {
	return wx.NewPostAction(urls.CardCodeUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
	)
}

// 删除卡券
func CardDelete(request *RequestCardDelete) wx.Action {
	return wx.NewPostAction(urls.CardDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
	)
}

// 设置卡券失效接口 自定义
func CardCodeUnavailableAuto(request *RequestCardCodeUnavailableAuto) wx.Action {
	return wx.NewPostAction(urls.CardCodeUnavailable,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
	)
}

// 设置卡券失效接口 非自定义
func CardCodeUnavailable(request *RequestCardCodeUnavailable) wx.Action {
	return wx.NewPostAction(urls.CardCodeUnavailable,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
	)
}

// 拉取卡券概况
func CardBizUinInfo(request *RequestGetCardBizUinInfo, result *RespGetCardBizUinInfo) wx.Action {
	return wx.NewPostAction(urls.CardBizUinInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(request)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)

}

func CardCodeDecrypt(req *RequestCardCodeDecrypt, result *RespCardCodeDecrypt) wx.Action {
	return wx.NewPostAction(urls.CardCodeDecrypt,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(req)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func CardFreeGet(req *RequestGetCardBizUinInfo, result *GetCardBizUinInfoList) wx.Action {
	return wx.NewPostAction(urls.CardFreeGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(req)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func LandingPageCreate(req *LandingPageCreateRequest, result *LandingPageCreateResponse) wx.Action {
	return wx.NewPostAction(urls.CardLandingPageCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(req)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
}
