package cast

import (
	"fmt"

	"strings"

	"github.com/flywithbug/utils"
)

func castJavaApi(a API) FileJavaCast {
	fi := FileJavaCast{}
	notes := formatApiNotes(a)

	fi.Name = fmt.Sprintf("EBike%sRequest", strFirstToUpper(a.Name))
	fi.Type = "api"
	impStr, attStr := formatPackageImport(a.ResponseName, "api", a.ParameterModel)
	fi.Content = fmt.Sprintf(temp.JavaApi, impStr, notes, a.Name, a.ResponseName, attStr, a.Name,
		a.Action, a.ResponseName, a.ResponseName)
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
	model.Name = strings.ReplaceAll(model.Name, "JYResp", "EBikeResp")
	model.Name = strings.ReplaceAll(model.Name, "JYPara", "EBikePara")
	fi.Name = model.Name
	fi.Type = "bean"

	impStr, attStr := formatPackageImport("", "bean", model)
	fi.Content = fmt.Sprintf(temp.JavaModel, impStr, model.Name, attStr)
	fi.Md5 = utils.Md5(fi.Content)
	return fi
}
