package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Data
// 用户数据分析
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
//getusersummary

//begin_date	是	获取数据的起始日期，begin_date和end_date的差值需小于“最大时间跨度”（比如最大时间跨度为1时，begin_date和end_date的差值只能为0，才能小于1），否则会报错
//end_date	是	获取数据的结束日期，end_date允许设置的最大值为昨日
type ParamsGetUserSummary struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type ResultGetUserSummary struct {
	List []*UserData `json:"list"` // 用户数据列表
}

// User数据 表示单个用户数据的响应结构体
type UserData struct {
	RefDate    string `json:"ref_date"`    // 数据的日期，格式为 "YYYY-MM-DD"
	UserSource int    `json:"user_source"` // 用户的渠道，数值代表的含义如下：
	// 0代表其他合计
	// 1代表公众号搜索
	// 17代表名片分享
	// 30代表扫描二维码
	// 57代表文章内账号名称
	// 100代表微信广告
	// 161代表他人转载
	// 149代表小程序关注
	// 200代表视频号
	// 201代表直播
	NewUser      int `json:"new_user"`      // 新增的用户数量
	CancelUser   int `json:"cancel_user"`   // 取消关注的用户数量
	CumulateUser int `json:"cumulate_user"` // 累计用户数量，new_user减去cancel_user即为净增用户数量
}

// GetUserSummary
// 获取用户增减数据（getusersummary）
func GetUserSummary(params *ParamsGetUserSummary, result *ResultGetUserSummary) wx.Action {
	return wx.NewPostAction(urls.GetUserSummary,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// 累计用户数据
func GetUserCumulate(params *ParamsGetUserSummary, result *ResultGetUserSummary) wx.Action {
	return wx.NewPostAction(urls.GetUserCumulate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGetArticleSummary struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// DailyArticleData 表示图文群发每日数据的响应结构体
type ResultDailyArticleData struct {
	List []*ArticleDailyData `json:"list"` // 每日文章数据列表
}

// ArticleDailyData 表示单日文章数据的响应结构体
type ArticleDailyData struct {
	RefDate          string `json:"ref_date"`            // 数据的日期，格式为 "YYYY-MM-DD"
	MsgID            string `json:"msgid"`               // 图文消息ID，格式为 "msgid_index"
	Title            string `json:"title"`               // 图文消息的标题
	IntPageReadUser  int    `json:"int_page_read_user"`  // 图文页阅读人数
	IntPageReadCount int    `json:"int_page_read_count"` // 图文页阅读次数
	OriPageReadUser  int    `json:"ori_page_read_user"`  // 原文页阅读人数
	OriPageReadCount int    `json:"ori_page_read_count"` // 原文页阅读次数
	ShareUser        int    `json:"share_user"`          // 分享人数
	ShareCount       int    `json:"share_count"`         // 分享次数
	AddToFavUser     int    `json:"add_to_fav_user"`     // 收藏人数
	AddToFavCount    int    `json:"add_to_fav_count"`    // 收藏次数
}

//getarticlesummary
//https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func GetArticleSummary(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetArticleSummary,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//getarticletotal
func GetArticleTotal(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetArticleTotal,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//获取图文统计数据（getuserread）	3	https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
func GetUserRead(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetUserRead,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//获取图文统计分时数据（getuserreadhour）	1	https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
func GetUserReadHour(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetUserReadHour,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//获取图文分享转发数据（getusershare）	7	https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
func GetUserShare(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetUserShare,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//获取图文分享转发分时数据（getusersharehour）	1	https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
func GetUserShareHour(params *ParamsGetArticleSummary, result *ResultDailyArticleData) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgHour,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ResultUpstreamMsg 表示获取消息发送概况数据的响应结构体
type ResultUpstreamMsg struct {
	List []*UpstreamMsgData `json:"list"`
}

// ResultUpstreamMsgHour 表示获取消息发送分时数据的响应结构体
type ResultUpstreamMsgHour struct {
	List []*UpstreamMsgHourData `json:"list"`
}

// ResultUpstreamMsgWeek 表示获取消息发送周数据的响应结构体
type ResultUpstreamMsgWeek struct {
	List []*UpstreamMsgData `json:"list"`
}

// ResultUpstreamMsgMonth 表示获取消息发送月数据的响应结构体
type ResultUpstreamMsgMonth struct {
	List []*UpstreamMsgData `json:"list"`
}

// ResultUpstreamMsgDist 表示获取消息发送分布数据的响应结构体
type ResultUpstreamMsgDist struct {
	List []*UpstreamMsgDistData `json:"list"`
}

// ResultUpstreamMsgDistWeek 表示获取消息发送分布周数据的响应结构体
type ResultUpstreamMsgDistWeek struct {
	List []*UpstreamMsgDistData `json:"list"`
}

// ResultUpstreamMsgDistMonth 表示获取消息发送分布月数据的响应结构体
type ResultUpstreamMsgDistMonth struct {
	List []*UpstreamMsgDistData `json:"list"`
}

// UpstreamMsgData 表示单条消息发送数据的响应结构体
type UpstreamMsgData struct {
	RefDate  string `json:"ref_date"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

// UpstreamMsgHourData 表示单条消息发送分时数据的响应结构体
type UpstreamMsgHourData struct {
	RefDate  string `json:"ref_date"`
	RefHour  int    `json:"ref_hour"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

// UpstreamMsgDistData 表示单条消息发送分布数据的响应结构体
type UpstreamMsgDistData struct {
	RefDate       string `json:"ref_date"`
	CountInterval int    `json:"count_interval"`
	MsgUser       int    `json:"msg_user"`
}

// ParamsGetUpstreamMsg 表示获取消息发送概况数据的请求参数
type ParamsGetUpstreamMsg struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgHour 表示获取消息发送分时数据的请求参数
type ParamsGetUpstreamMsgHour struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgWeek 表示获取消息发送周数据的请求参数
type ParamsGetUpstreamMsgWeek struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgMonth 表示获取消息发送月数据的请求参数
type ParamsGetUpstreamMsgMonth struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgDist 表示获取消息发送分布数据的请求参数
type ParamsGetUpstreamMsgDist struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgDistWeek 表示获取消息发送分布周数据的请求参数
type ParamsGetUpstreamMsgDistWeek struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// ParamsGetUpstreamMsgDistMonth 表示获取消息发送分布月数据的请求参数
type ParamsGetUpstreamMsgDistMonth struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// GetUpstreamMsg 获取消息发送概况数据
func GetUpstreamMsg(params *ParamsGetUpstreamMsg, result *ResultUpstreamMsg) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUpstreamMsgHour 获取消息发送分时数据
func GetUpstreamMsgHour(params *ParamsGetUpstreamMsgHour, result *ResultUpstreamMsgHour) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgHour,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUpstreamMsgWeek 获取消息发送周数据
func GetUpstreamMsgWeek(params *ParamsGetUpstreamMsgWeek, result *ResultUpstreamMsgWeek) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgWeek,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUpstreamMsgMonth 获取消息发送月数据
func GetUpstreamMsgMonth(params *ParamsGetUpstreamMsgMonth, result *ResultUpstreamMsgMonth) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgMonth,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUpstreamMsgDist 获取消息发送分布数据
func GetUpstreamMsgDist(params *ParamsGetUpstreamMsgDist, result *ResultUpstreamMsgDist) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgDist,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUpstreamMsgDistWeek 获取消息发送分布周数据
func GetUpstreamMsgDistWeek(params *ParamsGetUpstreamMsgDistWeek, result *ResultUpstreamMsgDistWeek) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgDistWeek,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//GetUpstreamMsgDistMonth
func GetUpstreamMsgDistMonth(params *ParamsGetUpstreamMsgDistMonth, result *ResultUpstreamMsgDistMonth) wx.Action {
	return wx.NewPostAction(urls.GetUpstreamMsgDistMonth,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
