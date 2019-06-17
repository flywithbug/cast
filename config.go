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
			ModelImport:   "#import <JYCastKit/JYResponseModel.h>",
			BaseModel:     "JYResponseModel",
			ModelPrefix:   "JYResp",
			ModelSuffix:   "Model",
			APImport:      "#import <JYCastKit/JYCastKit.h>",
			ParaBaseModel: "JYQueryParameter",
			ParaPrefix:    "JYQueryPara",
			ParaSuffix:    "",
			APIPrefix:     "mapi_",
			APISuffix:     "_withModuleType",
		},
	}
	return config
}

func SetFileModel(model FileModel) {
	Conf().FileModel = model
}

type Config struct {
	FileModel FileModel
}
