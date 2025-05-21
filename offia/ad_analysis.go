package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"strconv"
)

//广告位类型名称（ad_slot）	广告位类型
//SLOT_ID_BIZ_BOTTOM	公众号底部广告
//SLOT_ID_BIZ_MID_CONTEXT	公众号文中广告
//SLOT_ID_BIZ_VIDEO_END	公众号视频后贴
//SLOT_ID_BIZ_SPONSOR	公众号互选广告
//SLOT_ID_BIZ_CPS	公众号返佣商品
//SLOT_ID_WEAPP_BANNER	小程序banner
//SLOT_ID_WEAPP_REWARD_VIDEO	小程序激励视频
//SLOT_ID_WEAPP_INTERSTITIAL	小程序插屏广告
//SLOT_ID_WEAPP_VIDEO_FEEDS	小程序视频广告
//SLOT_ID_WEAPP_VIDEO_BEGIN	小程序视频前贴
//SLOT_ID_WEAPP_BOX

type AdSlot string

const (
	SLOT_ID_BIZ_BOTTOM         AdSlot = "SLOT_ID_BIZ_BOTTOM"
	SLOT_ID_BIZ_MID_CONTEXT    AdSlot = "SLOT_ID_BIZ_MID_CONTEXT"
	SLOT_ID_BIZ_VIDEO_END      AdSlot = "SLOT_ID_BIZ_VIDEO_END"
	SLOT_ID_BIZ_SPONSOR        AdSlot = "SLOT_ID_BIZ_SPONSOR"
	SLOT_ID_BIZ_CPS            AdSlot = "SLOT_ID_BIZ_CPS"
	SLOT_ID_WEAPP_BANNER       AdSlot = "SLOT_ID_WEAPP_BANNER"
	SLOT_ID_WEAPP_REWARD_VIDEO AdSlot = "SLOT_ID_WEAPP_REWARD_VIDEO"
	SLOT_ID_WEAPP_INTERSTITIAL AdSlot = "SLOT_ID_WEAPP_INTERSTITIAL"
	SLOT_ID_WEAPP_VIDEO_FEEDS  AdSlot = "SLOT_ID_WEAPP_VIDEO_FEEDS"
	SLOT_ID_WEAPP_VIDEO_BEGIN  AdSlot = "SLOT_ID_WEAPP_VIDEO_BEGIN"
	SLOT_ID_WEAPP_BOX          AdSlot = "SLOT_ID_WEAPP_BOX"
)

//参数	是否必须	说明
//page	是	返回第几页数据
//page_size	是	当页返回数据条数
//start_date	是	获取数据的开始时间 yyyy-mm-dd
//end_date	是	获取数据的结束时间 yyyy-mm-dd
//ad_slot	否	广告位类型名称
// ParamsPublisherAdPosGeneral 表示获取公众号分广告位数据的请求参数
type ParamsPublisherAdPosGeneral struct {
	Page      int    `url:"page"`
	PageSize  int    `url:"page_size"`
	StartDate string `url:"start_date"`        // 获取数据的开始时间 yyyy-mm-dd
	EndDate   string `url:"end_date"`          // 获取数据的结束时间 yyyy-mm-dd
	AdSlot    string `url:"ad_slot,omitempty"` // 广告位类型名称
}

// ResultPublisherAdPosGeneral 表示获取公众号分广告位数据的响应结构体
type ResultPublisherAdPosGeneral struct {
	BaseResp struct {
		ErrMsg string `json:"err_msg"`
		Ret    int    `json:"ret"`
	} `json:"base_resp"`
	List []struct {
		SlotID        int64   `json:"slot_id"`
		AdSlot        string  `json:"ad_slot"`
		Date          string  `json:"date"`
		ReqSuccCount  int     `json:"req_succ_count"`
		ExposureCount int     `json:"exposure_count"`
		ExposureRate  float64 `json:"exposure_rate"`
		ClickCount    int     `json:"click_count"`
		ClickRate     float64 `json:"click_rate"`
		Income        int     `json:"income"`
		Ecpm          float64 `json:"ecpm"`
		SlotStr       string  `json:"slot_str"`
	} `json:"list"`
	Summary struct {
		ReqSuccCount  int     `json:"req_succ_count"`
		ExposureCount int     `json:"exposure_count"`
		ExposureRate  float64 `json:"exposure_rate"`
		ClickCount    int     `json:"click_count"`
		ClickRate     float64 `json:"click_rate"`
		Income        int     `json:"income"`
		Ecpm          float64 `json:"ecpm"`
	} `json:"summary"`
	TotalNum int `json:"total_num"`
}

// GetPublisherAdPosGeneral 获取公众号分广告位数据
func GetPublisherAdPosGeneral(params *ParamsPublisherAdPosGeneral, result *ResultPublisherAdPosGeneral) wx.Action {
	return wx.NewGetAction(urls.PublisherAdPosGeneral,
		wx.WithQuery("page", strconv.Itoa(params.Page)),
		wx.WithQuery("action", "publisher_adpos_general"),
		wx.WithQuery("page_size", strconv.Itoa(params.PageSize)),
		wx.WithQuery("start_date", params.StartDate),
		wx.WithQuery("end_date", params.EndDate),
		//wx.WithQuery("ad_slot", params.AdSlot),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ParamsPublisherCpsGeneral 表示获取公众号返佣商品数据的请求参数
type ParamsPublisherCpsGeneral struct {
	Page      int    `url:"page"`
	PageSize  int    `url:"page_size"`
	StartDate string `url:"start_date"` // 获取数据的开始时间 yyyy-mm-dd
	EndDate   string `url:"end_date"`   // 获取数据的结束时间 yyyy-mm-dd
}

// ResultPublisherCpsGeneral 表示获取公众号返佣商品数据的响应结构体
type ResultPublisherCpsGeneral struct {
	BaseResp struct {
		ErrMsg string `json:"err_msg"`
		Ret    int    `json:"ret"`
	} `json:"base_resp"`
	List []struct {
		Date            string  `json:"date"`
		ExposureCount   int     `json:"exposure_count"`
		ClickCount      int     `json:"click_count"`
		ClickRate       float64 `json:"click_rate"`
		OrderCount      int     `json:"order_count"`
		OrderRate       float64 `json:"order_rate"`
		TotalFee        int     `json:"total_fee"`
		TotalCommission int     `json:"total_commission"`
	} `json:"list"`
	Summary struct {
		ExposureCount   int     `json:"exposure_count"`
		ClickCount      int     `json:"click_count"`
		ClickRate       float64 `json:"click_rate"`
		OrderCount      int     `json:"order_count"`
		OrderRate       float64 `json:"order_rate"`
		TotalFee        int     `json:"total_fee"`
		TotalCommission int     `json:"total_commission"`
	} `json:"summary"`
	TotalNum int `json:"total_num"`
}

// GetPublisherCpsGeneral 获取公众号返佣商品数据
func GetPublisherCpsGeneral(params *ParamsPublisherCpsGeneral, result *ResultPublisherCpsGeneral) wx.Action {
	return wx.NewGetAction(urls.PublisherCpsGeneral,
		wx.WithQuery("page", strconv.Itoa(params.Page)),
		wx.WithQuery("page_size", strconv.Itoa(params.PageSize)),
		wx.WithQuery("start_date", params.StartDate),
		wx.WithQuery("end_date", params.EndDate),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ParamsPublisherSettlement 表示获取公众号结算收入数据及结算主体信息的请求参数
type ParamsPublisherSettlement struct {
	Page      int    `url:"page"`
	PageSize  int    `url:"page_size"`
	StartDate string `url:"start_date"` // 获取数据的开始时间 yyyy-mm-dd
	EndDate   string `url:"end_date"`   // 获取数据的结束时间 yyyy-mm-dd
}

// ResultPublisherSettlement 表示获取公众号结算收入数据及结算主体信息的响应结构体
type ResultPublisherSettlement struct {
	BaseResp struct {
		ErrMsg string `json:"err_msg"`
		Ret    int    `json:"ret"`
	} `json:"base_resp"`
	Body              string `json:"body"`
	RevenueAll        int    `json:"revenue_all"`
	PenaltyAll        int    `json:"penalty_all"`
	SettledRevenueAll int    `json:"settled_revenue_all"`
	SettlementList    []struct {
		Date           string `json:"date"`
		Zone           string `json:"zone"`
		Month          string `json:"month"`
		Order          int    `json:"order"`
		SettStatus     int    `json:"sett_status"`
		SettledRevenue int    `json:"settled_revenue"`
		SettNo         string `json:"sett_no"`
		MailSendCnt    int    `json:"mail_send_cnt"`
		SlotRevenue    []struct {
			SlotID             int64 `json:"slot_id"`
			SlotSettledRevenue int   `json:"slot_settled_revenue"`
		} `json:"slot_revenue"`
	} `json:"settlement_list"`
	TotalNum int `json:"total_num"`
}

// GetPublisherSettlement 获取公众号结算收入数据及结算主体信息
func GetPublisherSettlement(params *ParamsPublisherSettlement, result *ResultPublisherSettlement) wx.Action {
	return wx.NewGetAction(urls.PublisherSettlement,
		wx.WithQuery("page", strconv.Itoa(params.Page)),
		wx.WithQuery("page_size", strconv.Itoa(params.PageSize)),
		wx.WithQuery("start_date", params.StartDate),
		wx.WithQuery("end_date", params.EndDate),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
