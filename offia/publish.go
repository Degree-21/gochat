package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"time"
)

type ReqPublishSubmit struct {
	MediaId string `json:"media_id"` //要发布的草稿的media_id
}

type ResultPublishSubmit struct {
	PublishId string `json:"publish_id"`
}

// 发布
func PublishSubmit(params *ReqPublishSubmit, result *ResultPublishSubmit) wx.Action {
	return wx.NewPostAction(urls.FreePublishSubmit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ReqPublishGet struct {
	PublishId string `json:"publish_id"`
}

type ResultPublishGet struct {
	PublishID     string         `json:"publish_id"`     // 发布任务id
	PublishStatus int            `json:"publish_status"` // 发布状态，0:成功, 1:发布中，2:原创失败, 3: 常规失败, 4:平台审核不通过, 5:成功后用户删除所有文章, 6: 成功后系统封禁所有文章
	ArticleID     string         `json:"article_id"`     // 当发布状态为0时（即成功）时，返回图文的 article_id，可用于“客服消息”场景
	ArticleDetail *ArticleDetail `json:"article_detail"` // 当发布状态为0时（即成功）时，返回文章详细信息
	FailIdx       []int          `json:"fail_idx"`       // 当发布状态为2或4时，返回不通过的文章编号，第一篇为 1；其他发布状态则为空
}

// ArticleDetail 结构体用于表示文章的详细信息
type ArticleDetail struct {
	Count int            `json:"count"` // 当发布状态为0时（即成功）时，返回文章数量
	Item  []*ArticleItem `json:"item"`  // 当发布状态为0时（即成功）时，返回文章列表
}

// ArticleItem 结构体用于表示单篇文章的信息
type ArticleItem struct {
	Idx        int    `json:"idx"`         // 当发布状态为0时（即成功）时，返回文章对应的编号
	ArticleURL string `json:"article_url"` // 当发布状态为0时（即成功）时，返回图文的永久链接
}

//发布状态轮询接口
func PublishGet(params *ReqPublishGet, result *ResultPublishGet) wx.Action {
	return wx.NewPostAction(urls.FreePublishGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ReqPublishDelete struct {
	ArticleId string `json:"article_id"`
	Index     int    `json:"index"`
}

type ResultPublishDelete struct {
	//ErrCode int    `json:"errcode"`
	//ErrMsg  string `json:"errmsg"`
}

// 删除草稿
func PublishDelete(params *ReqPublishDelete, result *ResultPublishDelete) wx.Action {
	return wx.NewPostAction(urls.FreePublishDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ReqPublishGetArticle struct {
	ArticleId string `json:"article_id"`
}

type ResultPublishGetArticle struct {
	NewsItem []*NewsItem `json:"news_item"` // 多图文消息应有多段 news_item 结构
}

//通过 article_id 获取已发布文章
func PublishGetArticle(params *ReqPublishGetArticle, result *ResultPublishGetArticle) wx.Action {
	return wx.NewPostAction(urls.FreePublishGetArticle,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ReqPublishBatchGetReq struct {
	Offset    int `json:"offset"`     //从全部素材的该偏移位置开始返回，0表示从第一个素材返回
	Count     int `json:"count"`      //返回素材的数量，取值在1到20之间
	NoContent int `json:"no_content"` //1 表示不返回 content 字段，0 表示正常返回，默认为 0
}

type ResultPublishBatchGet struct {
	TotalCount int                   `json:"total_count"` // 成功发布素材的总数
	ItemCount  int                   `json:"item_count"`  // 本次调用获取的素材的数量
	Item       []*PublishArticleItem `json:"item"`        // 可能有多个图文消息item结构
}
type ArticleContent struct {
	NewsItem []*NewsItem `json:"news_item"` // 多图文消息会在此处有多篇文章
}

// ArticleItem 结构体用于表示单条图文消息的完整信息
type PublishArticleItem struct {
	ArticleID  string          `json:"article_id"`  // 成功发布的图文消息id
	Content    *ArticleContent `json:"content"`     // 图文消息的内容
	UpdateTime time.Time       `json:"update_time"` // 这篇图文消息素材的最后更新时间
}

//批量获取已发布图文素材
func PublishBatchGet(params *ReqPublishBatchGetReq, result *ResultPublishBatchGet) wx.Action {
	return wx.NewPostAction(urls.FreePublishBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
