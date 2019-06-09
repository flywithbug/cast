package cast

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/flywithbug/utils"
)

const (
	AttributeTypeString = "String" //String类型
	AttributeTypeInt    = "Int"    //Int类型
	AttributeTypeFloat  = "Float"  //浮点数
	AttributeTypeBool   = "Bool"   //布尔类型
	AttributeTypeArray  = "Array"  //数组 （只能是其他的类型模型 必须要有ModelId）
	AttributeTypeObject = "Object" //模型
)

/**
 * @author  ori
 * @date 2019-06-08 17:12
 * @description: 属性信息
 */
type Attribute struct {
	Name      string `json:"name"`       //属性名
	ModelName string `json:"model_name"` //属性type 为 Array时 数组元素类型
	Type      string `json:"type"`       //属性类型
	Notes     string `json:"notes"`      //注释
	Alias     string `json:"alias"`      //别名：Json 捏取值Key
	Hash      string `json:"hash"`       //Hash key
	Required  bool   `json:"required"`   //是否必须
	Default   string `json:"default"`    //默认值
}

/**
 * @author  ori
 * @date 2019-06-08 17:21
 * @description: TODO
 */
type DataModel struct {
	Name         string
	Type         string //Model , Parameter (返回数据模型，请求参数模型: 默认空值为数据模型)
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
	HashOrigin   string //Hash原值
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
	d.HashOrigin = d.Name + ":["
	for i, v := range d.Attributes {
		d.HashOrigin += fmt.Sprintf("(%s,%s)", v.Type, v.Name)
		if i < len(d.Attributes)-1 {
			d.HashOrigin += ","
		}
	}
	d.HashOrigin += "]"
	d.Hash = utils.Md5(d.HashOrigin)
	return d.Hash
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
