package offia

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ReqDraftAdd struct {
	ArticleType        string       `json:"article_type"`          // 文章类型，分别有图文消息（news）、图片消息（newspic），不填默认为图文消息（news）
	Title              string       `json:"title"`                 // 标题
	Author             string       `json:"author"`                // 作者
	Digest             string       `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	Content            string       `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
	ContentSourceURL   string       `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string       `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久MediaID）
	NeedOpenComment    int          `json:"need_open_comment"`     // Uint32 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment int          `json:"only_fans_can_comment"` // Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
	PicCrop2351        string       `json:"pic_crop_235_1"`        // 封面裁剪为2.35:1规格的坐标字段
	PicCrop11          string       `json:"pic_crop_1_1"`          // 封面裁剪为1:1规格的坐标字段
	ImageInfo          *ImageInfo   `json:"image_info"`            // 图片消息里的图片相关信息
	CoverInfo          *CoverInfo   `json:"cover_info"`            // 封面裁剪信息
	ProductInfo        *ProductInfo `json:"product_info"`          // 商品相关信息
	Url                string       `json:"url"`
}

// ImageInfo 结构体用于表示图片消息里的图片相关信息
type ImageInfo struct {
	ImageList []*ImageMedia `json:"image_list"` // 图片列表，最多20张
}

// ImageMedia 结构体用于表示单张图片信息
type ImageMedia struct {
	ImageMediaID string `json:"image_media_id"` // 图片素材id（必须是永久MediaID）
}

// CoverInfo 结构体用于表示封面裁剪信息
type CoverInfo struct {
	CropPercentList []*CropPercent `json:"crop_percent_list"` // 裁剪比例列表
}

// CropPercent 结构体用于表示裁剪比例信息
type CropPercent struct {
	Ratio string `json:"ratio"` // 裁剪比例，支持：“1_1”，“16_9”,“2.35_1”
	X1    string `json:"x1"`    // 裁剪后图片左上角的x坐标
	Y1    string `json:"y1"`    // 裁剪后图片左上角的y坐标
	X2    string `json:"x2"`    // 裁剪后图片右下角的x坐标
	Y2    string `json:"y2"`    // 裁剪后图片右下角的y坐标
}

// ProductInfo 结构体用于表示商品相关信息
type ProductInfo struct {
	FooterProductInfo *FooterProductInfo `json:"footer_product_info"` // 文末插入商品相关信息
}

// FooterProductInfo 结构体用于表示文末插入商品相关信息
type FooterProductInfo struct {
	ProductKey string `json:"product_key"` // 商品key
}

// DraftAddResult 结构体用于表示新增草稿的返回结果
type DraftAddResult struct {
	MediaID string `json:"media_id"` // 上传后的获取标志，长度不固定，但不会超过128字符
}

//// AddNews 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
//func AddNews(params *ParamsNewsAdd, result *ResultMaterialAdd) wx.Action {
//	return wx.NewPostAction(urls.OffiaNewsAdd,
//		wx.WithBody(func() ([]byte, error) {
//			return json.Marshal(params)
//		}),
//		wx.WithDecode(func(resp []byte) error {
//			return json.Unmarshal(resp, result)
//		}),
//	)
//}

// 新增草稿
func DraftAdd(params *ReqDraftAdd, result *ReqDraftAdd) wx.Action {
	return wx.NewPostAction(urls.DraftAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ReqDraftGet struct {
	MediaID string `json:"media_id"`
}

type DraftGetResult struct {
	NewsItem []*DraftGetNewsItem `json:"news_item"`
}

type DraftGetNewsItem struct {
	ArticleType        string       `json:"article_type"`          // 文章类型，分别有图文消息（news）、图片消息（newspic），不填默认为图文消息（news）
	Title              string       `json:"title"`                 // 标题
	Author             string       `json:"author"`                // 作者
	Digest             string       `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	Content            string       `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
	ContentSourceURL   string       `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string       `json:"thumb_media_id"`        // 图文消息的封面图片素材id（一定是永久MediaID）
	ShowCoverPic       int          `json:"show_cover_pic"`        // 是否在正文显示封面。平台已不支持此功能，因此默认为0，即不展示
	NeedOpenComment    int          `json:"need_open_comment"`     // Uint32 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment int          `json:"only_fans_can_comment"` // Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
	URL                string       `json:"url"`                   // 草稿的临时链接
	ImageInfo          *ImageInfo   `json:"image_info"`            // 图片消息里的图片相关信息
	ProductInfo        *ProductInfo `json:"product_info"`          // 商品相关信息
}

// draftGet方法
func DraftGet(params *ReqDraftGet, result *DraftGetResult) wx.Action {
	return wx.NewPostAction(urls.DraftGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

//DraftDelete
type ReqDraftDelete struct {
	MediaID string `json:"media_id"` // 上传后的获取标志，长度不固定，但不会超过128字符
}

type DraftDeleteResult struct {
}

func DraftDelete(params *ReqDraftDelete, result *DraftDeleteResult) wx.Action {
	return wx.NewPostAction(urls.DraftDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
