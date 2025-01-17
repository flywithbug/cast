package cast

import (
	"fmt"

	"strings"

	"time"

	"github.com/flywithbug/utils"
)

func castObjectiveApiFile(a API) FileModelCast {
	fi := FileModelCast{}
	notes := formatApiNotes(a)
	if strings.ToLower(a.Method) == "post" {
		fi.H = fmt.Sprintf(temp.ApiH, a.ParameterName,
			a.ResponseName, a.Name, notes, a.Name,
			a.ParameterName, a.ResponseName)
		fi.M = fmt.Sprintf(temp.ApiM, a.Name, a.Name, a.Name, a.ParameterName, a.ResponseName, a.Action, a.ResponseName)
	} else {
		fi.H = fmt.Sprintf(temp.ApiGetH, a.ParameterName,
			a.ResponseName, a.Name, notes, a.Name,
			a.ParameterName, a.ResponseName)
		fi.M = fmt.Sprintf(temp.ApiGetM, a.Name, a.Name, a.Name, a.ParameterName, a.ResponseName, a.Action, a.ResponseName)
	}

	fi.Name = fmt.Sprintf("%s+%s", temp.BaseApiClass, a.Name)
	fi.Md5 = utils.Md5(fi.H + fi.M)
	fi.MImport = fmt.Sprintf("#import \"%s.h\"", fi.Name)
	fi.MContent = strings.ReplaceAll(fi.M, fi.MImport, "")

	fi.Type = "api"
	return fi
}

func formatApiNotes(a API) string {
	str := fmt.Sprintf(" * Url:\t\t%s\n * Dever:\t\t%s\n * AddTime:\t%s\n * UpdTime:\t%s\n * Status:\t\t%s\n * Notes:\t\t%s",
		a.DocUrl, a.Author, formatUninxTims(a.AddTime), formatUninxTims(a.UpDateTime), a.Status, a.Notes)
	return str
}

func formatUninxTims(timesamtp int) string {
	tm := time.Unix(int64(timesamtp), 0)
	//fmt.Println(timesamtp, tm)
	return fmt.Sprintf(tm.Format("2006-01-02 15:04:05"))
}
