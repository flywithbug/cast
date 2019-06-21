package cast

import (
	"fmt"

	"strings"

	"github.com/flywithbug/utils"
)

func castObjectiveModelFile(model DataModel) FileModelCast {
	fi := FileModelCast{}
	a, im, con := attributesFormat(model.Attributes)
	fi.H = fmt.Sprintf(temp.ModelH, im, model.Name, model.ParentName, a)
	fi.M = fmt.Sprintf(temp.ModelM, model.Name, model.Name, con)
	fi.Name = model.Name
	fi.Type = "model"
	fi.Md5 = utils.Md5(fi.H + fi.M)
	return fi
}

func attributesFormat(list []Attribute) (string, string, string) {
	str := ""
	impStr := ""
	containerStr := ""
	impMap := make(map[string]bool)
	for _, v := range list {
		str += attributeProperty(v.Name, v.Type, v.ModelName)
		str += "\n"
		if !defaultClassType(v.Type, v.ModelName) {
			impMap[v.ModelName] = true
			if v.Type == "array" {
				containerStr += fmt.Sprintf("@\"%s\" : [%s class],\n", v.Name, v.ModelName)
			}
		}
	}
	for v := range impMap {
		if !IsDefaultType(v) {
			impStr += fmt.Sprintf("#import \"%s.h\"", v)
			impStr += "\n"
		}
	}
	containerStr = strings.ReplaceAll(containerStr, "\n", "")
	return str, impStr, containerStr
}

func attributeProperty(name, aType, modelName string) string {
	switch strings.ToLower(aType) {
	case "string":
		return fmt.Sprintf("@property (nonatomic, copy)  \tNSString *%s;", name)
	case "bool", "boolean":
		return fmt.Sprintf("@property (nonatomic, assign)\tBOOL %s;", name)
	case "number", "float":
		return fmt.Sprintf("@property (nonatomic, strong)\tNSNumber *%s;", name)
	case "integer":
		return fmt.Sprintf("@property (nonatomic, assign)\tNSInteger %s;", name)
	case "object":
		return fmt.Sprintf("@property (nonatomic, strong)\t%s *%s;", modelName, name)
	case "array":
		modelName = modelNameTransform(modelName)
		return fmt.Sprintf("@property (nonatomic, copy)  \tNSArray <%s *> *%s;", modelName, name)
	}
	return ""
}

func modelNameTransform(name string) string {
	switch strings.ToLower(name) {
	case "string":
		return "NSString"
	case "bool", "boolean", "number", "float", "integer":
		return "NSNumber"
	case "array":
		return "NSArray"
	}
	return name
}

func defaultClassType(Type, modelName string) bool {
	switch strings.ToLower(Type) {
	case "string", "bool", "boolean", "number", "float", "integer":
		return true
	case "nsstring", "nsnumber", "nsinteger":
		return true
	case "array":
		if modelName != "" {
			return defaultClassType(modelName, "")
		}
	}
	return false
}

func defalutClassKitType(Type string) bool {
	switch strings.ToLower(Type) {
	case "string", "bool", "boolean", "number", "float", "integer":
		return true
	case "nsstring", "nsnumber", "nsinteger", "nsarray":
		return true
	}
	return false
}
