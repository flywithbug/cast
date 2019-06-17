package cast

import (
	"fmt"
	"strings"
	"time"
)

//const (
//	AttributeTypeString = "string"  //String类型
//	AttributeTypeFloat  = "float"   //浮点数
//	AttributeTypeNumber = "number"  //数据Number类型  注意转换时值变化
//	AttributeTypeInt    = "integer" //Int类型
//	AttributeTypeBool   = "boolean" //布尔类型
//	AttributeTypeArray  = "array"   //数组 （只能是其他的类型模型 必须要有ModelId）
//	AttributeTypeObject = "object"  //模型
//)

var (
	header = "//\n//  %s.h\n//  JYCastKit\n//\n//  Created by Flywithbug on %s.\n//  " +
		"Copyright © 2019 hellobike." +
		" All rights reserved.\n//\n"

	importStr  = "#import <JYCastKit/JYCastKit.h>"
	importStr1 = "#import \"%s.h\""
	interF     = "@interface %s : JYResponseModel"

	implementation = "@implementation %s"

	stringProperty    = "@property (nonatomic, copy)  \tNSString *%s"
	boolProperty      = "@property (nonatomic, assign)\tBOOL %s"
	numberProperty    = "@property (nonatomic, strong)\tNSNumber *%s"
	objectProperty    = "@property (nonatomic, strong)\t%s *%s"
	arrayProperty     = "@property (nonatomic, copy)  \tNSArray <%s *> *%s"
	containerProperty = "@\"%s\" : [%s class]"
)

type ObjectiveFileModel struct {
	DataModel      DataModel
	ModelName      string //模型名字
	Import         string
	Header         string            //文件头
	ContainerModel map[string]string //h文件
	Interface      string
	Attribute      string
	NullBEGIN      string
	NullEnd        string
	ImportM        string
	Implementation string
	ImpContent     string
}

func (ob *ObjectiveFileModel) formatHeader() {
	mName := formatModelName(ob.DataModel.Name)
	ob.ModelName = mName
	now := time.Now().Format("2006-01-02")
	ob.Header = fmt.Sprintf(header, mName, now)
}

func (ob *ObjectiveFileModel) formModelContainerProperty() {
	str := "+ (NSDictionary *)modelContainerPropertyGenericClass {\n\t// value should be Class or Class name.\n"
	str += "\treturn @{"
	for k, v := range ob.ContainerModel {
		str += fmt.Sprintf(containerProperty, v, k)
		str += ","
	}
	str += "};\n}"

	str = strings.Replace(str, "class],};", "class]};", -1)
	ob.ImpContent = str
}

func (ob *ObjectiveFileModel) formatAttributesAndImport() {
	var str string
	containerModel := make(map[string]string)
	for _, v := range ob.DataModel.Attributes {
		switch strings.ToLower(v.Type) {
		case "string":
			str += fmt.Sprintf(stringProperty, v.Name)
			str += "\n"
		case "bool", "boolean":
			str += fmt.Sprintf(boolProperty, v.Name)
			str += "\n"
		case "number", "float":
			str += fmt.Sprintf(numberProperty, v.Name)
			str += "\n"
		case "object":
			mName := formatModelName(v.ModelName)
			str += fmt.Sprintf(objectProperty, mName, v.Name)
			str += "\n"
			containerModel[mName] = v.Name
		case "array":
			mName := formatModelName(v.ModelName)
			str += fmt.Sprintf(arrayProperty, mName, v.Name)
			str += "\n"
			containerModel[mName] = v.Name
		}
	}
	ob.Attribute = str
	ob.ContainerModel = containerModel
}

func formObjectiveFileModel(model DataModel) *ObjectiveFileModel {
	obM := new(ObjectiveFileModel)
	obM.DataModel = model
	obM.formatHeader()
	obM.formatAttributesAndImport()
	obM.Import = importStr
	obM.Import += "\n"
	for k := range obM.ContainerModel {
		obM.Import += fmt.Sprintf(importStr1, k)
		obM.Import += "\n"
	}
	obM.NullBEGIN = "NS_ASSUME_NONNULL_BEGIN"
	obM.NullEnd = "NS_ASSUME_NONNULL_END"
	obM.Interface = fmt.Sprintf(interF, obM.ModelName)
	obM.ImportM = fmt.Sprintf(importStr1, obM.ModelName)
	obM.Implementation = fmt.Sprintf(implementation, obM.ModelName)
	obM.formModelContainerProperty()
	return obM
}

func CastModelObjective_C_H_M(model DataModel) (h, m string, err error) {
	obm := formObjectiveFileModel(model)
	h = obm.castObjective_C_H()
	m = obm.castObjective_C_M()
	return
}

func (ob *ObjectiveFileModel) castObjective_C_H() (str string) {
	str = ob.Header
	str += "\t\n"
	str += ob.Import
	str += "\n"
	str += "\t\n"

	str += ob.NullBEGIN
	str += "\n\t\n"

	str += ob.Interface
	str += "\n"
	str += ob.Attribute
	str += "@end\n"
	str += "\t\n"
	str += ob.NullEnd
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	return str
}

func (ob *ObjectiveFileModel) castObjective_C_M() (str string) {
	str = ob.Header
	str += "\t\n"
	str += ob.ImportM
	str += "\n"
	str += "\t\n"

	str += ob.Implementation
	str += "\n\t\n"

	str += ob.ImpContent
	str += "\n"
	str += "@end\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	return str
}

//模型名称生产
func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}
