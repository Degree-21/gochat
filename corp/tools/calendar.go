package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Calendar struct {
	CalID       string           `json:"cal_id"`
	Organizer   string           `json:"organizer"`
	ReadOnly    int              `json:"readonly"`
	Summary     string           `json:"summary"`
	Color       string           `json:"color"`
	Description string           `json:"description"`
	Shares      []*CalendarShare `json:"shares"`
}

type CalendarShare struct {
	UserID   string `json:"userid"`
	ReadOnly int    `json:"readonly,omitempty"`
}

type ParamsCalendarAdd struct {
	Calendar *CalendarAddData `json:"calendar"`
	AgentID  int64            `json:"agentid"`
}

type CalendarAddData struct {
	Organizer   string           `json:"organizer"`
	ReadOnly    int              `json:"readonly,omitempty"`
	SetAsDfault int              `json:"set_as_dfault"`
	Summary     string           `json:"summary"`
	Color       string           `json:"color"`
	Description string           `json:"description,omitempty"`
	Shares      []*CalendarShare `json:"shares,omitempty"`
}

type ResultCalendarAdd struct {
	CalID string `json:"cal_id"`
}

func AddCalendar(params *ParamsCalendarAdd, result *ResultCalendarAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsCalendarAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCalendarUpdate struct {
	Calendar *CalendarUpdateData `json:"calendar"`
}

type CalendarUpdateData struct {
	CalID       string           `json:"cal_id"`
	ReadOnly    int              `json:"read_only"`
	Summary     string           `json:"summary"`
	Color       string           `json:"color"`
	Description string           `json:"description"`
	Shares      []*CalendarShare `json:"shares"`
}

func UpdateCalendar(params *ParamsCalendarUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsCalendarUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsCalendarGet struct {
	CalIDList []string `json:"cal_id_list"`
}

type ResultCalendarGet struct {
	CalendarList []*Calendar `json:"calendar_list"`
}

func GetCalendar(params *ParamsCalendarGet, result *ResultCalendarGet) wx.Action {
	return wx.NewPostAction(urls.CorpToolsCalendarGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCalendarDelete struct {
	CalID string `json:"cal_id"`
}

func DeleteCalendar(params *ParamsCalendarDelete) wx.Action {
	return wx.NewPostAction(urls.CorpToolsCalendarDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
