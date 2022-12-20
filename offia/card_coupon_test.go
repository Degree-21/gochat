/*
@Time : 2022/7/14 3:37 PM
@Author : 21
@File : card_coupon_test
@Software: GoLand
*/
package offia

import (
	"context"
	"fmt"
	"testing"
)

const accessToken string = ""

const Appid string = ""

const AppSecret string = ""

func TestDebugUploadCardCoupon(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html
	result := new(ResultMaterialAdd)
	// 素材
	oa := New(Appid, AppSecret)
	err := oa.Do(context.Background(), accessToken, UploadImage("../mock/test.jpg", result))
	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	request := &RequestCreateCard{Card: &Card{
		CardType: "",
		BaseInfo: nil,
		Groupon:  nil,
	}}

	cardResult := &RespCardCard{}
	// 上传卡券
	err = oa.Do(context.Background(), accessToken, CreateCardCoupon(request, cardResult))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("success")
}

func TestCreateCardCouponQrCode(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html
	//result := new(ResultMaterialAdd)
	// 素材
	oa := New(Appid, AppSecret)

	request := &RequestCreateCardCouponQrCode{
		ActionName:    "",
		ExpireSeconds: 0,
		ActionInfo: struct {
			Card struct {
				CardId       string `json:"card_id"`
				Code         string `json:"code"`
				Openid       string `json:"openid"`
				IsUniqueCode bool   `json:"is_unique_code"`
				OuterStr     string `json:"outer_str"`
			} `json:"card"`
		}{},
	}

	resp := &RespCreateCardCouponQrCode{
		Ticket:        "",
		ExpireSeconds: 0,
		Url:           "",
		ShowQrcodeUrl: "",
	}

	err := oa.Do(context.Background(), accessToken, CreateCardCouponQrCode(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestGetCardCode(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
	oa := New(Appid, AppSecret)

	request := &RequestCardCodeGet{
		CardId:       "",
		Code:         "",
		CheckConsume: false,
	}

	resp := &RespCardCodeGet{
		Card: struct {
			CardId    string `json:"card_id"`
			BeginTime int    `json:"begin_time"`
			EndTime   int    `json:"end_time"`
		}{},
		Openid:         "",
		CanConsume:     false,
		UserCardStatus: "",
	}

	err := oa.Do(context.Background(), accessToken, GetCardCode(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")

}

func TestConsumeCardCode(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
	oa := New(Appid, AppSecret)

	request := &RequestCardCodeConsume{
		Code:   "",
		CardId: "",
	}

	resp := &RespCardCodeConsume{
		Card: struct {
			CardId string `json:"card_id"`
		}{},
		Openid: "",
	}

	err := oa.Do(context.Background(), accessToken, ConsumeCardCode(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")

}

func TestGetUserCardList(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
	oa := New(Appid, AppSecret)

	request := &RequestUserCardList{
		Openid: "",
		CardId: "",
	}

	resp := &RespUserCardList{
		CardList:     nil,
		HasShareCard: false,
	}

	err := oa.Do(context.Background(), accessToken, GetUserCardList(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")

}

func TestGetCard(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html
	oa := New(Appid, AppSecret)

	request := &RequestCardGet{
		CardId: "",
	}

	resp := &RespCardGet{Card: &CardGetCard{
		CardType: "",
		Discount: nil,
	}}

	err := oa.Do(context.Background(), accessToken, GetCard(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")

}

func TestGetBatchCardList(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardBatchGet{
		Offset:     0,
		Count:      0,
		StatusList: nil,
	}

	resp := &RespCardBatchGet{
		CardIdList: nil,
		TotalNum:   0,
	}

	err := oa.Do(context.Background(), accessToken, GetBatchCardList(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestUpdateCard(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardUpdate{
		CardId:     "",
		MemberCard: nil,
	}

	resp := &RespCardUpdate{SendCheck: false}

	err := oa.Do(context.Background(), accessToken, UpdateCard(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

// 修改卡券库存
func TestCardModifySocket(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardModifySocket{
		CardId:             "",
		IncreaseStockValue: 0,
		ReduceStockValue:   0,
	}

	err := oa.Do(context.Background(), accessToken, CardModifySocket(request))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCardCodeUpdate(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardCodeUpdate{
		Code:    "",
		CardId:  "",
		NewCode: "",
	}
	err := oa.Do(context.Background(), accessToken, CardCodeUpdate(request))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCardDelete(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardDelete{}
	err := oa.Do(context.Background(), accessToken, CardDelete(request))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCardCodeUnavailable(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardCodeUnavailable{}
	err := oa.Do(context.Background(), accessToken, CardCodeUnavailable(request))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCardCodeUnavailableAuto(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestCardCodeUnavailableAuto{}
	err := oa.Do(context.Background(), accessToken, CardCodeUnavailableAuto(request))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is (%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}

func TestCardBizUinInfo(t *testing.T) {
	// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
	oa := New(Appid, AppSecret)

	request := &RequestGetCardBizUinInfo{
		BeginDate:  "",
		EndDate:    "",
		CondSource: 0,
	}

	resp := &RespGetCardBizUinInfo{}
	err := oa.Do(context.Background(), accessToken, CardBizUinInfo(request, resp))

	if err != nil {
		fmt.Println(fmt.Errorf("error message is 1(%v)", err.Error()))
		return
	}
	fmt.Println("error end")
}
