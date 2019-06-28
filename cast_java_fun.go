package cast

import (
	"fmt"
	"strings"
)

func formatJavaProperty(note, name, aType, modelName string) string {
	switch strings.ToLower(aType) {
	case "string", "nsstring":
		return fmt.Sprintf("\tprivate String %s;//%s", name, note)
	case "bool", "boolean":
		return fmt.Sprintf("\tprivate boolean %s;//%s", name, note)
	case "number", "nsnumber":
		return fmt.Sprintf("\tprivate double %s;//%s", name, note)
	case "float":
		return fmt.Sprintf("\tprivate float %s;//%s", name, note)
	case "integer":
		return fmt.Sprintf("\tprivate int %s;//%s", name, note)
	case "object":
		return fmt.Sprintf("\tprivate %s %s;//%s", modelName, name, note)
	case "array":
		modelName = formatJavaPropertyType(modelName)
		return fmt.Sprintf("\tprivate List<%s> %s;//%s", modelName, name, note)
	}
	return ""
}

func formatJavaPropertyType(aType string) string {
	switch strings.ToLower(aType) {
	case "string", "nsstring":
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
