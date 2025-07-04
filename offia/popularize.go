package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type QRCodeAction string

const (
	QRScene         QRCodeAction = "QR_SCENE"           // 临时的整型参数值
	QRStrScene      QRCodeAction = "QR_STR_SCENE"       // 临时的字符串参数值
	QRLimitScene    QRCodeAction = "QR_LIMIT_SCENE"     // 永久的整型参数值
	QRLimitStrScene QRCodeAction = "QR_LIMIT_STR_SCENE" // 永久的字符串参数值
)

type QRCodeScene struct {
	SceneID  int    `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type QRCodeActionInfo struct {
	Scene *QRCodeScene `json:"scene"`
}

type ParamsQRCodeCreate struct {
	ActionName    QRCodeAction      `json:"action_name"`
	ActionInfo    *QRCodeActionInfo `json:"action_info"`
	ExpireSeconds int               `json:"expire_seconds,omitempty"`
}

type ResultQRCodeCreate struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

// CreateQRCode 创建临时二维码（expireSeconds：二维码有效时间，最大不超过2592000秒（即30天），不填，则默认有效期为30秒。）
func CreateQRCode(params *ParamsQRCodeCreate, result *ResultQRCodeCreate) wx.Action {
	return wx.NewPostAction(urls.OffiaQRCodeCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsQRCodeShow struct {
	Ticket string `json:"ticket"`
}

type ResultQRCodeShow struct {
	Body []byte `json:"body"`
}

func QRCodeShow(params *ParamsQRCodeShow, result *ResultQRCodeShow) wx.Action {
	return wx.NewGetAction(urls.OffiaQRCodeShow,
		wx.WithQuery("ticket", string(params.Ticket)),
		wx.WithDecode(func(resp []byte) error {
			result.Body = resp
			return nil
		}),
	)
}

type ParamsShortURL struct {
	Action  string `json:"action"`
	LongURL string `json:"long_url"`
}

// ResultShortURL 短链接
type ResultShortURL struct {
	ShortURL string `json:"short_url"`
}

// ShortURL 长链接转短链接（长链接支持http://、https://、weixin://wxpay格式的url）
func ShortURL(longURL string, result *ResultShortURL) wx.Action {
	params := &ParamsShortURL{
		Action:  "long2short",
		LongURL: longURL,
	}

	return wx.NewPostAction(urls.OffiaShortURLGenerate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
