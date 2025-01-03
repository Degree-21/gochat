package user

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsDepartmentCreate struct {
	Name     string `json:"name"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid"`
	Order    int64  `json:"order,omitempty"`
}

type ResultDepartmentCreate struct {
	ID int64 `json:"id"`
}

// CreateDepartment 创建部门
func CreateDepartment(params *ParamsDepartmentCreate, result *ResultDepartmentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpDepartmentCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsDepartmentUpdate struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid,omitempty"`
	Order    int64  `json:"order,omitempty"`
}

// UpdateDepartment 更新部门
func UpdateDepartment(params *ParamsDepartmentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpDepartmentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// DeleteDepartment 删除部门
func DeleteDepartment(id int64) wx.Action {
	return wx.NewGetAction(urls.CorpDepartmentDelete,
		wx.WithQuery("id", strconv.FormatInt(id, 10)),
	)
}

type Department struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	NameEN   string `json:"name_en"`
	ParentID int64  `json:"parentid"`
	Order    int64  `json:"order"`
}

type ResultDepartmentList struct {
	Department []*Department `json:"department"`
}

// GetDepartmentList 获取部门列表
func GetDepartmentList(id int64, result *ResultDepartmentList) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if id != 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id, 10)))
	}

	return wx.NewGetAction(urls.CorpDepartmentList, options...)
}
