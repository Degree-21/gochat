package minip

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type JumpWxa struct {
	Path       string `json:"path"`
	Query      string `json:"query"`
	EnvVersion string `json:"env_version"`
}

type ParamsGenerateScheme struct {
	JumpWxa        *JumpWxa `json:"jump_wxa"`
	IsExpire       bool     `json:"is_expire"`
	ExpireType     int      `json:"expire_type"`
	ExpireInterval int      `json:"expire_interval"`
	ExpireTime     int      `json:"expire_time"`
}

// URLLinkRequest 定义生成小程序 URL Link 的请求结构体
type URLLinkRequest struct {
	Path           string     `json:"path,omitempty"`            // 否 通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Query          string     `json:"query,omitempty"`           // 否 通过 URL Link 进入小程序时的 query，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	ExpireType     int        `json:"expire_type,omitempty"`     // 否 默认值 0。小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
	ExpireTime     int64      `json:"expire_time,omitempty"`     // 否 到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为 30 天。expire_type 为 0 必填
	ExpireInterval int        `json:"expire_interval,omitempty"` // 否 到期失效的 URL Link 的失效间隔天数。生成的到期失效 URL Link 在该间隔时间到达前有效。最长间隔天数为 30 天。expire_type 为 1 必填
	CloudBase      *CloudBase `json:"cloud_base,omitempty"`      // 否 云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
	EnvVersion     string     `json:"env_version,omitempty"`     // 否 默认值 "release"。要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"，仅在微信外打开时生效。
}

// CloudBase 云开发静态网站自定义 H5 配置参数
type CloudBase struct {
	Env           string `json:"env"`                      // 是 云开发环境
	Domain        string `json:"domain,omitempty"`         // 否 静态网站自定义域名，不填则使用默认域名
	Path          string `json:"path,omitempty"`           // 否 云开发静态网站 H5 页面路径，不可携带 query
	Query         string `json:"query,omitempty"`          // 否 云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	ResourceAppID string `json:"resource_appid,omitempty"` // 否 第三方批量代云开发时必填，表示创建该 env 的 appid （小程序/第三方平台）
	EnvVersion    string `json:"env_version,omitempty"`    // 否 默认值 "release"。要打开的小程序版本。
}

// URLLinkResponse 定义生成小程序 URL Link 的返回结构体
type URLLinkResponse struct {
	ErrCode int    `json:"errcode"`  // 错误码
	ErrMsg  string `json:"errmsg"`   // 错误信息
	URLLink string `json:"url_link"` // 生成的小程序 URL Link
}

type ResultGenerateScheme struct {
	Openlink string `json:"openlink"`
}

func GenerateScheme(params *ParamsGenerateScheme, result *ResultGenerateScheme) wx.Action {
	return wx.NewPostAction(urls.MinipSchemeGenerateScheme,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func GenerateURLLink(params *URLLinkRequest, result *URLLinkResponse) wx.Action {
	return wx.NewPostAction(urls.MinipGenerateUrllink,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
