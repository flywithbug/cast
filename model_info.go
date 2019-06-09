package cast

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/flywithbug/utils"
)

const ()

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
	Type         string //Model , Parameter (返回数据模型，请求参数模型)
	Alias        string
	Desc         string
	CreateTime   string
	StartVersion string
	EndVersion   string
	Owner        string
	Attributes   []Attribute
	ParentName   string //父类类名
	ParentHash   string //父类hash值 md5
	Hash         string //属性值的Hash md5
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

func (d *DataModel) HashId() string {
	d.SortAttributesByName()
	origin := d.Name + ":["
	for i, v := range d.Attributes {
		origin += v.Name
		if i < len(d.Attributes)-1 {
			origin += ","
		}
	}
	origin += "]"
	fmt.Println(origin)
	return utils.Md5(origin)
}

func (d DataModel) Valid() bool {
	if d.Name == "" {
		return false
	}
	if len(d.Attributes) == 0 {
		return false
	}
	return true
}
