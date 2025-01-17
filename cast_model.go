package cast

import (
	"fmt"

	"strings"

	"sort"

	"github.com/flywithbug/utils"
)

func castObjectiveModelFile(model DataModel) FileModelCast {
	fi := FileModelCast{}
	model.sortAttributes()
	a, im, con := attributesFormat(model.Attributes)
	fi.H = fmt.Sprintf(temp.ModelH, im, model.Name, model.ParentName, a)
	fi.M = fmt.Sprintf(temp.ModelM, model.Name, model.Name, con)
	fi.Name = model.Name

	fi.MImport = fmt.Sprintf("#import \"%s.h\"", fi.Name)
	fi.MContent = strings.ReplaceAll(fi.M, fi.MImport, "")

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
		if v.FatherAtt {
			continue
		}
		notes := formatNotes(v.Required, v.Notes)
		str += attributeProperty(notes, v.Name, v.Type, v.ModelName)
		str += "\n"
		if !defaultClassType(v.Type, v.ModelName) {
			impMap[v.ModelName] = true
			if v.Type == "array" {
				containerStr += fmt.Sprintf("@\"%s\" : [%s class],\n", v.Name, v.ModelName)
			}
		}
	}

	list1 := make([]string, 0)
	for v := range impMap {
		if !IsDefaultType(v) {
			list1 = append(list1, v)
			//impStr += fmt.Sprintf("#import \"%s.h\"", v)
			//impStr += "\n"
		}
	}
	sort.Strings(list1)
	for _, v := range list1 {
		impStr += fmt.Sprintf("#import \"%s.h\"", v)
		impStr += "\n"
	}

	containerStr = strings.ReplaceAll(containerStr, "\n", "")
	return str, impStr, containerStr
}

func attributeProperty(note, name, aType, modelName string) string {
	switch strings.ToLower(aType) {
	case "string":
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, copy)  \tNSString *%s;", note, name)
	case "bool", "boolean":
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, assign)\tBOOL %s;", note, name)
	case "number", "float":
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, strong)\tNSNumber *%s;", note, name)
	case "integer":
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, assign)\tNSInteger %s;", note, name)
	case "object":
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, strong)\t%s *%s;", note, modelName, name)
	case "array":
		modelName = modelNameTransform(modelName)
		return fmt.Sprintf("\n/**\n%s\n*/\n@property (nonatomic, copy)  \tNSArray <%s *> *%s;", note, modelName, name)
	}
	return ""
}

func modelNameTransform(name string) string {
	switch strings.ToLower(name) {
	case "string", "nsstring":
		return "NSString"
	case "bool", "boolean", "number", "float", "integer", "nsnumber":
		return "NSNumber"
	case "array", "nsarray":
		return "NSArray"
	}
	return name
}

func defaultClassType(Type, modelName string) bool {
	switch strings.ToLower(Type) {
	case "string", "bool", "boolean", "number", "float", "integer":
		return true
	case "nsstring", "nsnumber":
		return true
	case "array":
		if modelName != "" {
			return defaultClassType(modelName, "")
		}
	}
	return false
}

func formatNotes(required bool, notes string) string {
	return fmt.Sprintf("Required:\t%v\nNotes:\t\t%s", required, notes)
}
