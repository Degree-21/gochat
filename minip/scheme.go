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
