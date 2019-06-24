package cast

//Api模型
type API struct {
	Id            int
	Title         string
	Name          string //api名字，独一无二
	Action        string //有可能重复
	Path          string
	Method        string
	Alias         string //别名(中文名)
	Notes         string
	ParameterName string //参数模型名字
	ResponseName  string //返回体模型名字
	//ParaOriginName string
	//RespOriginName string

	Status        string //接口状态
	Author        string //创建人
	AddTime       int
	UpDateTime    int
	NoResponse    bool
	NoParameter   bool
	CategoryClass string

	DocUrl string //api 地址
}

//数据模型
type DataModel struct {
	Id   int    //对于的ApiId
	Name string //添加前缀后的名字
	//OriginName string //原始名字 for java
	Type       string //model parameter
	Alias      string //后续使用
	Desc       string //模型描述
	ParentName string //父类名称
	Attributes []Attribute
	OriginType string
}

//数据属性模型
type Attribute struct {
	Name      string
	ModelName string //Objective-C中 string，boolean,number,array 之类转换类型时 NSArray NSString
	Type      string //array, object, string, boolean, float , number,integer
	Alias     string //别名
	Required  bool
	Default   string
	Notes     string
	FatherAtt bool //父类中是否已存在（针对iOS中）
}

type FileModelCast struct {
	H    string
	M    string
	Name string //文件名
	Type string //model api
	Md5  string

	//.m文件聚合时使用
	MContent string
	MImport  string
}
