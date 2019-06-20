package cast

//Api模型
type API struct {
	Id     int
	Name   string
	Action string
	Path   string
	Method string
	Alias  string //别名(中文名)
	Notes  string
	//ParameterId   int    //参数模型Id
	//ResponseId    int    //返回模型Id
	ParameterName string //参数模型名字
	ResponseName  string //返回体模型名字
	Status        string //接口状态
	Author        string //创建人
	AddTime       int
	UpDateTime    int
	NoResponse    bool
	NoParameter   bool
	CategoryClass string
}

//数据模型
type DataModel struct {
	Name       string //添加前缀后的名字
	OriginName string //原始名字 for java
	Type       string //Model Parameter
	Alias      string //后续使用
	Desc       string //模型描述
	ParentName string //父类名称
	Attributes []Attribute
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
}

type FileModelCast struct {
	H    string
	M    string
	Name string //文件名
	Type string //model api
	Md5  string
}
