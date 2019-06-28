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
	if a.NoResponse {
		responseName = "Object"
	}
	fi.Type = "api"
	fi.Name = formatApiName(a.Name)
	apiClass := temp.JavaApiClass
	if a.RequiredToken {
		apiClass = temp.JavaApiReqLoginClass
	}
	impStr, attStr := formatPackageImport(responseName, "api", a.ParameterModel)
	fi.Content = fmt.Sprintf(temp.JavaApi, impStr, notes, fi.Name, apiClass, responseName, attStr, fi.Name,
		a.Action, responseName, responseName)
	fi.Md5 = utils.Md5(fi.Content)
	return fi
}

func formatPackageImport(respName, aType string, para DataModel) (impStr, attStr string) {
	impStr = fmt.Sprintf("package %s.%s;", temp.JavaPackage, aType)
	impStr += "\n"
	if respName != "" && respName != "Object" {
		impStr += fmt.Sprintf("import %s.bean.%s;\n", temp.JavaPackage, respName)
	}
	hasList := false
	for _, v := range para.Attributes {

		v.ModelName = formatModelName(v.ModelName)

		if !defaultClassType(v.Type, v.ModelName) && v.ModelName != "" {
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
	name = fmt.Sprintf("EBike%sRequest", strFirstToUpper(name))
	return strHumpToUpper(name)
}

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func strHumpToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return capitalize(temp[0]) + upperStr
}

func capitalize(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
