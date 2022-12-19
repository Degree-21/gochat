package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type VipCardActivateRequest struct {
	ActivateBeginTime     *string `json:"activate_begin_time,omitempty"`      // 激活后的有效起始时间，若不填写默认以创建时的 data_info 为准。Unix时间戳格式。
	ActivateEndTime       *string `json:"activate_end_time,omitempty"`        // 激活后的有效截至时间，若不填写默认以创建时的 data_info 为准。Unix时间戳格式。
	BackgroundPicURL      *string `json:"background_pic_url,omitempty"`       // 商家自定义会员卡背景图，须 先调用 上传图片接口 将背景图上传至CDN，否则报错， 卡面设计请遵循 微信会员卡自定义背景设计规范
	CardID                *string `json:"card_id,omitempty"`                  // 卡券id，自定义code卡券必填
	Code                  string  `json:"code"`                               // 领取会员卡用户获得得code
	InitBalance           *int64  `json:"init_balance,omitempty"`             // 初始余额，不填为0
	InitBonus             *int64  `json:"init_bonus,omitempty"`               // 初始积分，不填为0
	InitBonusRecord       *string `json:"init_bonus_record,omitempty"`        // 积分同步说明
	InitCustomFieldValue1 *string `json:"init_custom_field_value1,omitempty"` // 创建时字段custom_field1定义类型的初始值
	InitCustomFieldValue2 *string `json:"init_custom_field_value2,omitempty"` // 创建时字段custom_field2定义类型的初始值
	InitCustomFieldValue3 *string `json:"init_custom_field_value3,omitempty"` // 创建时字段custom_field3定义类型的初始值
	MembershipNumber      string  `json:"membership_number"`                  // 会员卡编号，会员卡编号，由开发者填入，作为序列号显示在用户的卡包里。可与 Code 码保持等值。
}

type VipCardActivateResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type GetVipCardInfoRequest struct {
	CardID string `json:"card_id"`
	Code   string `json:"code"`
}

type GetVipCardInfoResponse struct {
	Bonus            int64           `json:"bonus"`      // 积分信息
	Errcode          int64           `json:"errcode"`    // 错误码，0为正常
	Errmsg           string          `json:"errmsg"`     // 错误信息
	HasActive        bool            `json:"has_active"` // 该卡是否已经被激活，true表示已经被激活，false表示未被激活
	MembershipNumber string          `json:"membership_number"`
	Nickname         string          `json:"nickname"`         // 用户昵称
	Openid           string          `json:"openid"`           // 用户在本公众号内唯一识别码
	Sex              string          `json:"sex"`              // 用户性别
	UserCardStatus   string          `json:"user_card_status"` // 当前用户的会员卡状态，NORMAL 正常 EXPIRE 已过期 GIFTING 转赠中 GIFT_SUCC 转赠成功 GIFT_TIMEOUT 转赠超时 DELETE; 已删除，UNAVAILABLE 已失效
	UserInfo         VipCardInfoUser `json:"user_info"`
}

type VipCardInfoUser struct {
	CommonFieldList []CommonFieldListElement `json:"common_field_list"` // 开发者设置的会员卡会员信息类目，如等级。
	CustomFieldList []CustomFieldListElement `json:"custom_field_list"` // 开发者设置的会员卡会员信息类目，如等级。
}

type CommonFieldListElement struct {
	Name  string `json:"name"`  // 会员信息类目名称
	Value string `json:"value"` // 会员卡信息类目值，比如等级值等
}

type CustomFieldListElement struct {
	Name      string   `json:"name"`  // 会员信息类目名称
	Value     string   `json:"value"` // 会员卡信息类目值，比如等级值等
	ValueList []string `json:"value_list"`
}

type UpdateVipCardUserInfoRequest struct {
	AddBalance        int64          `json:"add_balance"`
	AddBonus          int64          `json:"add_bonus"`
	BackgroundPicURL  string         `json:"background_pic_url"`
	Balance           int64          `json:"balance"`
	Bonus             int64          `json:"bonus"`
	CardID            string         `json:"card_id"`
	Code              string         `json:"code"`
	CustomFieldValue1 string         `json:"custom_field_value1"`
	CustomFieldValue2 string         `json:"custom_field_value2"`
	NotifyOptional    NotifyOptional `json:"notify_optional"`
	RecordBalance     string         `json:"record_balance"`
	RecordBonus       string         `json:"record_bonus"`
}

type NotifyOptional struct {
	IsNotifyBalance      bool `json:"is_notify_balance"`
	IsNotifyBonus        bool `json:"is_notify_bonus"`
	IsNotifyCustomField1 bool `json:"is_notify_custom_field1"`
}

type UpdateVipCardUserInfoResponse struct {
	Errcode       int64  `json:"errcode"`
	Errmsg        string `json:"errmsg"`
	ResultBonus   int    `json:"result_bonus"`
	ResultBalance int    `json:"result_balance"`
	Openid        string `json:"openid"`
}

type UpdateVipCardRequest struct {
	CardID     string     `json:"card_id"`
	MemberCard MemberCard `json:"member_card"`
}

type UpdateVipCardResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	SendCheck bool   `json:"send_check"`
}

//激活会员卡
func VipCardActivate(req *VipCardActivateRequest, result *VipCardActivateResponse) wx.Action {
	return wx.NewPostAction(urls.VipCardActivate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(req)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
}

//获取会员卡信息
func GetVipCardInfo(req *GetVipCardInfoRequest, result *GetVipCardInfoResponse) wx.Action {
	return wx.NewPostAction(urls.GetVipCardInfo, wx.WithBody(func() ([]byte, error) {
		return json.Marshal(req)
	}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
}

//更新用户信息
func UpdateVipCardUserInfo(req *UpdateVipCardUserInfoRequest, result *UpdateVipCardUserInfoResponse) wx.Action {
	return wx.NewPostAction(urls.UpdateVipCardUserInfo, wx.WithBody(func() ([]byte, error) {
		return json.Marshal(req)
	}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
}

//更新会员卡信息
func UpdateVipCardInfo(req *UpdateVipCardRequest, result *UpdateVipCardResponse) wx.Action {
	return wx.NewPostAction(urls.UpdateVipCardInfo, wx.WithBody(func() ([]byte, error) {
		return json.Marshal(req)
	}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}))
}
