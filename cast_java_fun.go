package cast

import (
	"fmt"
	"strings"
)

func formatJavaProperty(note, name, aType, modelName string) string {
	switch strings.ToLower(aType) {
	case "string":
		return fmt.Sprintf("\tprivate\t\tString %s;//%s", name, note)
	case "bool", "boolean":
		return fmt.Sprintf("\tprivate\t\tboolean %s//%s;", name, note)
	case "number", "float":
		return fmt.Sprintf("\tprivate\t\tfloat %s//%s;", name, note)
	case "integer":
		return fmt.Sprintf("\tprivate\t\tint %s;//%s", name, note)
	case "object":
		return fmt.Sprintf("\tprivate\t\t%s %s;//%s", modelName, name, note)
	case "array":
		modelName = formatJavaPropertyType(modelName)
		return fmt.Sprintf("\tprivate\t\tList<%s> %s;//%s", modelName, name, note)
	}
	return ""
}

func formatJavaPropertyType(aType string) string {
	switch strings.ToLower(aType) {
	case "string":
		return "String"
	case "bool", "boolean":
		return "boolean"
	case "number", "integer", "int":
		return "int"
	case "float":
		return "float"
	}
	return aType
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
