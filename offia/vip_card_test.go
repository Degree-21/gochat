package offia

import (
	"context"
	"fmt"
	"testing"
)

const (
	VipAppid     = ""
	VipAppSecret = ""
	VipToken     = ""
)

//pLqZ6wxzzT7LEuf54fwOfJCv-Yf8
//创建卡券
func TestCreateVipCard(t *testing.T) {
	oa := New(VipAppid, VipAppSecret)

	request := &RequestCreateCard{
		Card: &Card{
			CardType: "MEMBER_CARD",
			MemberCard: &MemberCard{
				Prerogative:   "123",
				SupplyBonus:   true,
				SupplyBalance: false,
				AutoActivate:  true,
				BaseInfo: BaseInfo{
					LogoURL:       "https://qny-cj.9w9.com/bd8f520221130140727.jpg",
					CodeType:      "CODE_TYPE_NONE",
					BrandName:     "活动抽奖",
					Title:         "活动抽奖会员卡",
					Color:         "Color010",
					Notice:        "notice",
					Description:   "description",
					UseCustomCode: true,
					Sku: struct {
						Quantity int `json:"quantity"`
					}{Quantity: 100000000},
					CenterTitle:            "点击领取",
					CenterSubTitle:         "点击领取！！！",
					CenterAppBrandUserName: "gh_c672c121cb23@app",
					CenterAppBrandPass:     "pages/home/index",
					GetLimit:               1,
					DateInfo: struct {
						Type           string `json:"type"`
						BeginTimestamp int    `json:"begin_timestamp"`
						EndTimestamp   int    `json:"end_timestamp"`
					}{
						Type: "DATE_TYPE_PERMANENT",
					},
				},
				CustomField1: CustomField1{
					NameType: "FIELD_NAME_TYPE_LEVEL",
					URL:      "https://www.qq.com",
				},
			},
		},
	}

	resp := &RespCardCard{}
	err := oa.Do(context.Background(), VipToken, CreateCardCoupon(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is 1(%v)", err.Error()))
		return
	}
	fmt.Println(resp)
	fmt.Println("error end")

}

//设置白名单
func TestSetCardWhitelist(t *testing.T) {
	oa := New(VipAppid, VipAppSecret)

	request := &RequestSetCardWhitelist{
		Openid: []string{"oLqZ6wxbP0fbzpsPOkIHyGqMAiho", "oLqZ6wx6b3P9c3EjuvcCdw4P3AVA"},
	}

	resp := &RespSetCardWhitelist{}
	err := oa.Do(context.Background(), VipToken, SetCardWhitelist(request, resp))
	fmt.Println("result:", resp)
	if err != nil {
		fmt.Println(fmt.Errorf("error message is 1(%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCreateCardCouponQrCode2(t *testing.T) {
	// 素材
	oa := New(VipAppid, VipAppSecret)

	request := &RequestCreateCardCouponQrCode{
		ActionName: "QR_CARD",
		ActionInfo: struct {
			Card struct {
				CardId       string `json:"card_id"`
				Code         string `json:"code"`
				Openid       string `json:"openid"`
				IsUniqueCode bool   `json:"is_unique_code"`
				OuterStr     string `json:"outer_str"`
			} `json:"card"`
		}{
			Card: struct {
				CardId       string `json:"card_id"`
				Code         string `json:"code"`
				Openid       string `json:"openid"`
				IsUniqueCode bool   `json:"is_unique_code"`
				OuterStr     string `json:"outer_str"`
			}{
				CardId: "pLqZ6w3Xnr3IbQbDMdFlQvaFV4oI",
				Code:   "110201201247",
			},
		},
	}

	resp := &RespCreateCardCouponQrCode{
		Ticket:        "",
		ExpireSeconds: 0,
		Url:           "",
		ShowQrcodeUrl: "",
	}

	err := oa.Do(context.Background(), VipToken, CreateCardCouponQrCode(request, resp))
	fmt.Println("result:", resp)
	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestUpdateVipCardUserInfo(t *testing.T) {
	// 素材
	oa := New(VipAppid, VipAppSecret)

	request := &UpdateVipCardUserInfoRequest{
		AddBalance:        0,
		AddBonus:          100,
		BackgroundPicURL:  "",
		Balance:           0,
		Bonus:             0,
		CardID:            "pLqZ6w6g2LUW9949GvJiwtmPeQGs",
		Code:              "110201201246",
		CustomFieldValue1: "",
		CustomFieldValue2: "",
		NotifyOptional:    NotifyOptional{},
		RecordBalance:     "",
		RecordBonus:       "",
	}

	resp := &UpdateVipCardUserInfoResponse{}

	err := oa.Do(context.Background(), VipToken, UpdateVipCardUserInfo(request, resp))
	fmt.Println("result:", resp)
	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestGetVipCardInfo(t *testing.T) {
	// 素材
	oa := New(VipAppid, VipAppSecret)

	request := &GetVipCardInfoRequest{
		CardID: "pLqZ6wxzzT7LEuf54fwOfJCv-Yf8",
		Code:   "110201201245",
	}

	resp := &GetVipCardInfoResponse{}

	err := oa.Do(context.Background(), VipToken, GetVipCardInfo(request, resp))
	fmt.Println("result:", resp)
	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}
