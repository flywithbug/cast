package cast

type ObjectiveAPIFileModel struct {
	Api Api

	DataModel DataModel
	ParaModel DataModel
	Import    string

	ParaModelName string //模型名字
	ApiFullName   string //JYMapiNetworkClient+user_ev_ride_getBikesByLocation
	ApiCategory   string

	Header         string            //文件头
	ContainerModel map[string]string //h文件
	Interface      string
	Attribute      string
	NullBEGIN      string
	NullEnd        string
	ImportM        string
	Implementation string
	ImpContent     string

	H string
	M string
}

func FormObjectiveAPIFileModel(aM Api) (apiFileM *ObjectiveAPIFileModel) {
	apiFileM = new(ObjectiveAPIFileModel)
	apiFileM.Api = aM
	i, ok := Cache().get(aM.ResponseModelName)
	if !ok {
		return nil
	}
	dM, ok := i.(DataModel)
	if !ok {
		return nil
	}
	apiFileM.DataModel = dM

	i, ok = Cache().get(aM.ParameterName)
	if !ok {
		return nil
	}
	dM, ok = i.(DataModel)
	if !ok {
		return nil
	}
	apiFileM.ParaModel = dM
	apiFileM.formHeader()

	return nil
}

func (a *ObjectiveAPIFileModel) formHeader() {
	a.ApiCategory = formatApiCategoryName(a.Api.Action)

	//mName := formatModelName(a.Api.Name)
	//ob.ModelName = mName
	//now := time.Now().Format("2006-01-02")
	//ob.Header = fmt.Sprintf(header, mName, now)
}
