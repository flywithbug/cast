package cast

import (
	"fmt"

	"github.com/flywithbug/utils"
)

func castObjectiveApiFile(a API) *FileModelCast {
	fi := new(FileModelCast)
	fi.H = fmt.Sprintf(temp.ApiH, a.ParameterName,
		a.ResponseName, a.Name, a.Name,
		a.ParameterName, a.ResponseName)
	fi.M = fmt.Sprintf(temp.ApiM, a.Name, a.Name, a.ParameterName, a.ResponseName, a.Action, a.ResponseName)
	fi.Name = fmt.Sprintf("#import \"%s+%s.h\"", temp.BaseApiClass, a.Name)
	fi.Md5 = utils.Md5(fi.H + fi.M)
	fi.Type = "api"
	return fi
}
