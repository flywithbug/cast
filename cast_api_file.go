package cast

import (
	"fmt"
	"time"
)

type ObjectiveAPIFileModel struct {
	Api Api

	DataModel DataModel
	ParaModel DataModel
	Import    string

	ParaModelName string //模型名字
	RespModelName string //模型名字

	ApiFullName string //JYMapiNetworkClient+user_ev_ride_getBikesByLocation
	ApiCategory string

	Header string //文件头

	ContainerModel map[string]string //h文件

	Interface string

	NullBEGIN      string
	NullEnd        string
	ImportM        string
	Implementation string
	Action         string
	ImpContent     string

	HMethodContainer string

	MImpMethod string
	H          string
	M          string
}

func formObjectiveAPIFileModel(aM Api) (apiFileM ObjectiveAPIFileModel) {
	apiFileM.ContainerModel = make(map[string]string)
	apiFileM.Api = aM
	i, ok := Cache().get(aM.ResponseModelName)
	if !ok {
		return
	}
	dM, ok := i.(DataModel)
	if !ok {
		return
	}
	apiFileM.DataModel = dM

	i, ok = Cache().get(aM.ParameterName)
	if !ok {
		return
	}
	dM, ok = i.(DataModel)
	if !ok {
		return
	}
	apiFileM.ParaModel = dM
	apiFileM.formContainerString()

	return apiFileM
}

func (a *ObjectiveAPIFileModel) formContainerString() {

	a.NullBEGIN = "NS_ASSUME_NONNULL_BEGIN"
	a.NullEnd = "NS_ASSUME_NONNULL_END"

	a.ApiCategory = formatApiCategoryName(a.Api.Action)
	a.ApiFullName = fmt.Sprintf("%s+%s", Conf().FileModel.APIClientClass, a.ApiCategory)
	a.ParaModelName = formatParaName(a.Api.ParameterName)
	a.RespModelName = formatModelName(a.Api.ResponseModelName)

	a.ContainerModel[a.ParaModelName] = a.ParaModelName
	a.ContainerModel[a.RespModelName] = a.RespModelName

	a.Interface = fmt.Sprintf("@interface %s (%s)", Conf().FileModel.APIClientClass, a.ApiCategory)

	a.HMethodContainer = fmt.Sprintf(categoryMethod, a.ApiCategory, a.ParaModelName, a.RespModelName)
	//mName := formatModelName(a.Api.Name)
	//ob.ModelName = mName
	now := time.Now().Format("2006-01-02")
	a.Header = fmt.Sprintf(header, a.ApiFullName, now)

	a.Import = Conf().FileModel.APImport
	a.Import += "\n"
	for k := range a.ContainerModel {
		a.Import += fmt.Sprintf(importStr1, k)
		a.Import += "\n"
	}

	a.ImportM = fmt.Sprintf(importStr1, a.ApiFullName)
	a.Implementation = fmt.Sprintf("@implementation %s (%s)", Conf().FileModel.APIClientClass, a.ApiCategory)

	a.Action = fmt.Sprintf(apiAction, a.Api.Action)

	a.MImpMethod = fmt.Sprintf(categoryImpMethod, a.ApiCategory,
		a.ParaModelName, a.RespModelName, a.RespModelName)

	a.formApiHFileString()
	a.formApiMFileString()

}

func (a *ObjectiveAPIFileModel) formApiHFileString() {

	str := a.Header
	str += "\t\n"
	str += a.Import
	str += "\n"
	str += "\t\n"

	str += a.NullBEGIN
	str += "\n\t\n"

	str += a.Interface
	str += "\n"

	str += a.HMethodContainer
	str += "\n"
	str += "\n"
	str += "@end\n"
	str += "\t\n"

	str += a.NullEnd
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	a.H = str
}

func (a *ObjectiveAPIFileModel) formApiMFileString() {
	str := a.Header
	str += "\t\n"

	str += a.ImportM

	str += "\n"
	str += "\t\n"

	str += "\n\t\n"

	str += a.Implementation
	str += "\n"
	str += a.Action
	str += "\n"
	str += "\n"
	str += a.MImpMethod
	str += "\n"
	str += "\n"
	str += "@end\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"
	str += "\t\n"

	a.M = str
}

func CastApiObjective_C_H_M(model Api) ObjectiveAPIFileModel {
	obm := formObjectiveAPIFileModel(model)
	obm.formContainerString()
	return obm
}
