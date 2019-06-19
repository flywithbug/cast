package cast

import (
	"fmt"
	"strings"
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

type ObjectiveModelFileModel struct {
	Type           string //model parameter
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

	H string
	M string
}

func (ob *ObjectiveModelFileModel) formatHeader() {
	mName := formatModelName(ob.DataModel.Name)
	if ob.DataModel.Type == ModelTypeTypeParameter {
		mName = formatParaName(ob.DataModel.Name)
	}
	ob.Header = fmt.Sprintf(header, mName)
	ob.ModelName = mName
}

func (ob *ObjectiveModelFileModel) formModelContainerProperty() {
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

func (ob *ObjectiveModelFileModel) formatAttributesAndImport() {
	var str string
	containerModel := make(map[string]string)
	for _, v := range ob.DataModel.Attributes {
		if Conf().PropertyFilter[v.Name] {
			continue
		}
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

func formObjectiveModelFileModel(model DataModel) *ObjectiveModelFileModel {
	obM := new(ObjectiveModelFileModel)
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

	if obM.DataModel.Type == ModelTypeTypeParameter {
		obM.Interface = fmt.Sprintf(interParaF, obM.ModelName)
	}
	obM.ImportM = fmt.Sprintf(importStr1, obM.ModelName)
	obM.Implementation = fmt.Sprintf(implementation, obM.ModelName)
	obM.formModelContainerProperty()
	return obM
}

func CastModelObjective_C_H_M(model DataModel) *ObjectiveModelFileModel {
	obm := formObjectiveModelFileModel(model)
	obm.H = obm.castObjective_C_H()
	obm.M = obm.castObjective_C_M()
	return obm
}

func (ob *ObjectiveModelFileModel) castObjective_C_H() (str string) {
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

func (ob *ObjectiveModelFileModel) castObjective_C_M() (str string) {
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
