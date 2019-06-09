package cast

import (
	"encoding/json"
	"sort"
)

/**
 * @author  ori
 * @date 2019-06-08 17:12
 * @description: 属性信息
 */
type Attribute struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Notes    string `json:"notes"`
	Alias    string `json:"alias"`
	Hash     string `json:"hash"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

/**
 * @author  ori
 * @date 2019-06-08 17:21
 * @description: TODO
 */
type DataModel struct {
	Name         string
	Alias        string
	Desc         string
	CreateTime   string
	StartVersion string
	EndVersion   string
	Owner        string
	Attributes   []Attribute
}

func (d DataModel) ToJson() string {
	js, _ := json.Marshal(d)
	return string(js)
}

func (d *DataModel) SortAttributesByName() {
	sort.Slice(d.Attributes, func(i, j int) bool {
		return d.Attributes[i].Name < d.Attributes[j].Name
	})
}
