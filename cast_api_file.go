package cast

type ObjectiveAPIFileModel struct {
	Api       Api
	DataModel DataModel
	Import    string

	ModelName      string            //模型名字
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

	return nil
}
