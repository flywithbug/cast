package cast

type FileModel struct {
	ModelImport   string //#import <JYCastKit/JYResponseModel.h>
	BaseModel     string //JYResponseModel
	ModelPrefix   string //JYResp
	ModelSuffix   string //Model
	APImport      string //#import <JYCastKit/JYCastKit.h>
	ParaBaseModel string //JYQueryParameter
	ParaPrefix    string //JYQueryPara
	ParaSuffix    string //空
	APIPrefix     string //mapi_
	APISuffix     string //_withModuleType:
}

type ObjectiveFileModel struct {
	Header         string //文件头
	ImportH        string //h文件
	ImportM        string //m文件
	ClassImport    string
	Interface      string
	Attribute      string
	Implementation string
	Content        string
}
