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

	importStr = "#import <JYCastKit/JYCastKit.h>"

	interF = "@interface %s : JYResponseModel"

	stringProperty = "@property (nonatomic, copy)  \tNSString *%s"
	boolProperty   = "@property (nonatomic, assign)\tBOOL %s"
	numberProperty = "@property (nonatomic, strong)\tNSNumber *%s"
	objectProperty = "@property (nonatomic, strong)\t%s *%s"
	arrayProperty  = "@property (nonatomic, copy)  \tNSArray <%s *> *%s"
)

func CastModelObjective_C_H(model DataModel) (str string, err error) {
	mName := formatModelName(model.Name)
	str = formHeader(model)

	str += "\t\n"

	str += importStr
	str += "\n"
	str += "\t\n"

	str += "NS_ASSUME_NONNULL_BEGIN"
	str += "\t\n"

	str = fmt.Sprintf(str+interF, mName)
	str += "\n"
	str += formatAttributes(model.Attributes)

	return
}

func formatAttributes(list []Attribute) (str string) {
	for _, v := range list {
		switch strings.ToLower(v.Type) {
		case "string":
			str += fmt.Sprintf(stringProperty, v.Name)
			str += "\n"
		case "bool", "boolean":
			str += fmt.Sprintf(boolProperty, v.Name)
			str += "\n"
		case "number":
			str += fmt.Sprintf(numberProperty, v.Name)
			str += "\n"
		case "object":
			str += fmt.Sprintf(objectProperty, formatModelName(v.ModelName), v.Name)
			str += "\n"
		case "array":
			str += fmt.Sprintf(arrayProperty, formatModelName(v.ModelName), v.Name)
			str += "\n"
		}
	}
	return str
}

func formHeader(model DataModel) string {
	mName := formatModelName(model.Name)
	now := time.Now().Format("2006-01-02")
	return fmt.Sprintf(header, mName, now)
}

//模型名称生产
func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}
