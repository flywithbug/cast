package cast

import (
	"fmt"

	"strings"

	"github.com/flywithbug/utils"
)

func castObjectiveApiFile(a API) FileModelCast {
	fi := FileModelCast{}
	fi.H = fmt.Sprintf(temp.ApiH, a.ParameterName,
		a.ResponseName, a.Name, a.Name,
		a.ParameterName, a.ResponseName)
	fi.M = fmt.Sprintf(temp.ApiM, a.Name, a.Name, a.Name, a.ParameterName, a.ResponseName, a.Action, a.ResponseName)

	fi.Name = fmt.Sprintf("%s+%s", temp.BaseApiClass, a.Name)
	fi.Md5 = utils.Md5(fi.H + fi.M)
	fi.MImport = fmt.Sprintf("#import \"%s.h\"", fi.Name)
	fi.MContent = strings.ReplaceAll(fi.M, fi.MImport, "")

	fi.Type = "api"
	return fi
}
