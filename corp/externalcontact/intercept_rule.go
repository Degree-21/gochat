package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ExtraRule struct {
	SemanticsList []int `json:"semantics_list,omitempty"`
}

type RuleApplicableRange struct {
	UserList      []string `json:"user_list,omitempty"`
	DeparmentList []int64  `json:"deparment_list,omitempty"`
}

type InterceptRule struct {
	RuleID          int64                `json:"rule_id"`
	RuleName        string               `json:"rule_name"`
	WordList        []string             `json:"word_list"`
	ExtraRule       *ExtraRule           `json:"extra_rule"`
	InterceptType   int                  `json:"intercept_type"`
	ApplicableRange *RuleApplicableRange `json:"applicable_range"`
}

type ParamsInterceptRuleAdd struct {
	RuleName        string               `json:"rule_name"`
	WordList        []string             `json:"word_list"`
	SemanticsList   []int                `json:"semantics_list,omitempty"`
	InterceptType   int                  `json:"intercept_type"`
	ApplicableRange *RuleApplicableRange `json:"applicable_range"`
}

type ResultInterceptRuleAdd struct {
	RuleID int64 `json:"rule_id"`
}

func AddInterceptRule(params *ParamsInterceptRuleAdd, result *ResultInterceptRuleAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleUpdate struct {
	RuleID                int64                `json:"rule_id"`
	RuleName              string               `json:"rule_name,omitempty"`
	WordList              []string             `json:"word_list,omitempty"`
	ExtraRule             *ExtraRule           `json:"extra_rule,omitempty"`
	InterceptType         int                  `json:"intercept_type,omitempty"`
	AddApplicableRange    *RuleApplicableRange `json:"add_applicable_range,omitempty"`
	RemoveApplicableRange *RuleApplicableRange `json:"remove_applicable_range,omitempty"`
}

func UpdateInterceptRule(params *ParamsInterceptRuleUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type RuleListData struct {
	RuleID     int64  `json:"rule_id"`
	RuleName   string `json:"rule_name"`
	CreateTime int64  `json:"create_time"`
}

type ResultInterceptRuleList struct {
	RuleList []*RuleListData `json:"rule_list"`
}

func ListInterceptRule(result *ResultInterceptRuleList) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactInterceptRuleList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleGet struct {
	RuleID int64 `json:"rule_id"`
}

type ResultInterceptRuleGet struct {
	Rule *InterceptRule `json:"rule"`
}

func GetInterceptRule(params *ParamsInterceptRuleGet, result *ResultInterceptRuleGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleDelete struct {
	RuleID int64 `json:"rule_id"`
}

func DeleteInterceptRule(params *ParamsInterceptRuleDelete) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
