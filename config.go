package cast

var config *Config

func Conf() *Config {
	if config == nil {
		config = confInit()
	}
	return config
}

func confInit() *Config {
	config = &Config{
		FileModel: FileModel{
			ModelImport:    "#import <JYCastKit/JYResponseModel.h>",
			BaseModel:      "JYResponseModel",
			ModelPrefix:    "JYResp",
			ModelSuffix:    "Model",
			APImport:       "#import <JYCastKit/JYCastKit.h>",
			ParaBaseModel:  "JYQueryParameter",
			ParaPrefix:     "JYQueryPara",
			ParaSuffix:     "",
			APIClientClass: "JYMapiNetworkClient",
			APIPrefix:      "mapi_",
			APISuffix:      "_withModuleType",
		},
		PropertyFilter: map[string]bool{
			"action":     true,
			"sourceId":   true,
			"systemCode": true,
			"ticket":     true,
			"token":      true,
			"version":    true,
			"platform":   true,
		},
	}
	return config
}

func SetFileModel(model FileModel) {
	Conf().FileModel = model
}

func SetPropertyFilter(filter map[string]bool) {
	Conf().PropertyFilter = filter
}

type Config struct {
	FileModel      FileModel
	PropertyFilter map[string]bool
}
