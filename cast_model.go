package cast

import (
	"fmt"
	"strings"
)

func castObjectiveModelFile(model DataModel) *FileModelCast {
	fi := new(FileModelCast)
	a, im, con := attributesFormat(model.Attributes)
	fi.H = fmt.Sprintf(temp.ModelH, im, model.Name, model.ParentName, a)
	fi.M = fmt.Sprintf(temp.ModelM, model.Name, model.Name, con)
	return fi
}

func attributesFormat(list []Attribute) (string, string, string) {
	str := ""
	impStr := ""
	containerStr := ""
	for _, v := range list {
		str += attributeProperty(v.Name, v.Type, v.ModelName)
		str += "\n"
		if !defaultClassType(v.Type, v.ModelName) {
			impStr += fmt.Sprintf("#import \"%s.h\"", v.ModelName)
			impStr += "\n"
			containerStr += fmt.Sprintf("@\"%s\" : [%s class],\n", v.Name, v.ModelName)
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
	case "array":
		if modelName != "" {
			return defaultClassType(modelName, "")
		}
	}
	return false
}
