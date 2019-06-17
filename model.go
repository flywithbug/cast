package cast

import (
	"encoding/json"
	"sort"
)

const (
	AttributeTypeString = "string"  //String类型
	AttributeTypeFloat  = "float"   //浮点数
	AttributeTypeNumber = "number"  //数据Number类型  注意转换时值变化
	AttributeTypeInt    = "integer" //Int类型
	AttributeTypeBool   = "boolean" //布尔类型
	AttributeTypeArray  = "array"   //数组 （只能是其他的类型模型 必须要有ModelId）
	AttributeTypeObject = "object"  //模型
)

type ModelTypeType int

const (
	_ ModelTypeType = iota
	ModelTypeTypeParameter
	ModelTypeTypeModel
)

/**
 * @author  ori
 * @date 2019-06-08 17:12
 * @description: 属性信息
 */
type Attribute struct {
	Name      string `json:"name"`       //属性名
	ModelName string `json:"model_name"` //属性type 为Array数组元素类型
	Type      string `json:"type"`       //属性类型
	Alias     string `json:"alias"`      //别名：Json 捏取值Key
	Hash      string `json:"hash"`       //Hash key
	Required  bool   `json:"required"`   //是否必须
	Default   string `json:"default"`    //默认值
	Notes     string `json:"notes"`      //注释,其他信息
	DataModel DataModel
}

/**
 * @author  ori
 * @date 2019-06-08 17:21
 * @description: TODO
 */
type DataModel struct {
	Name       string        //模型名称
	Type       ModelTypeType //Model , Parameter (返回数据模型，请求参数模型: 默认空值为数据模型)
	Alias      string
	Desc       string
	Attributes []Attribute
	ParentName string //父类类名
	//ParentHash string //父类hash值 md5
	//Hash       string //属性值的Hash md5
	//HashOrigin string //Hash原值
}

type ResponseModel struct {
	Code string    //返回Code
	Msg  string    //msg
	Data DataModel //data body
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

//func (d *DataModel) HashId() string {
//	d.SortAttributesByName()
//	d.HashOrigin = d.Name + ":["
//	for i, v := range d.Attributes {
//		d.HashOrigin += fmt.Sprintf("(%s,%s)", v.Type, v.Name)
//		if i < len(d.Attributes)-1 {
//			d.HashOrigin += ","
//		}
//	}
//	d.HashOrigin += "]"
//	d.Hash = utils.Md5(d.HashOrigin)
//	return d.Hash
//}

func (d DataModel) Valid() bool {
	if d.Name == "" {
		return false
	}
	if len(d.Attributes) == 0 {
		return false
	}
	return true
}
