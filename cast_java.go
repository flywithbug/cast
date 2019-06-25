package cast

import (
	"fmt"

	"strings"

	"github.com/flywithbug/utils"
)

func castJavaApi(a API) FileJavaCast {
	fi := FileJavaCast{}
	notes := formatApiNotes(a)
	responseName := formatModelName(a.ResponseName)
	fi.Type = "api"
	fi.Name = formatApiName(a.Name)
	impStr, attStr := formatPackageImport(responseName, "api", a.ParameterModel)
	fi.Content = fmt.Sprintf(temp.JavaApi, impStr, notes, fi.Name, responseName, attStr, fi.Name,
		a.Action, responseName, responseName)
	fi.Md5 = utils.Md5(fi.Content)
	return fi
}

func formatPackageImport(respName, aType string, para DataModel) (impStr, attStr string) {
	impStr = fmt.Sprintf("package %s.%s;", temp.JavaPackage, aType)
	impStr += "\n"
	if respName != "" {
		impStr += fmt.Sprintf("import %s.bean.%s;\n", temp.JavaPackage, respName)
	}
	hasList := false
	for _, v := range para.Attributes {

		v.ModelName = formatModelName(v.ModelName)

		if !defaultClassType(v.Type, v.ModelName) {
			impStr += fmt.Sprintf("import %s.bean.%s;\n", temp.JavaPackage, v.ModelName)
		}
		if v.Type == "array" {
			hasList = true
		}
		attStr += formatJavaProperty(v.Notes, v.Name, v.Type, v.ModelName)
		attStr += "\n"
	}

	if hasList {
		impStr += "import java.util.List;\n"
	}
	return
}

func castJavaModel(model DataModel) FileJavaCast {
	fi := FileJavaCast{}
	model.Name = formatModelName(model.Name)
	fi.Name = model.Name
	fi.Type = "bean"

	impStr, attStr := formatPackageImport("", "bean", model)
	fi.Content = fmt.Sprintf(temp.JavaModel, impStr, model.Name, attStr)
	fi.Md5 = utils.Md5(fi.Content)
	return fi
}

func formatModelName(name string) string {
	name = strings.ReplaceAll(name, "JYResp", "EBikeResp")
	name = strings.ReplaceAll(name, "JYPara", "EBikePara")
	return name
}

func formatApiName(name string) string {
	return fmt.Sprintf("EBike%sRequest", strFirstToUpper(name))
}
