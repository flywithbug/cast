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
		BaseApiClass:   "JYMapiNetworkClient",
		BaseQueryClass: "JYQueryParameter",
		BaseModelClass: "JYResponseModel",
	}
	return config
}

type Config struct {
	BaseModelClass string //基类模型名称
	BaseQueryClass string //基类参数模型名称
	BaseApiClass   string // JYMapiNetworkClient
}
