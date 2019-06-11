package cast

type ApiMethodType int

const (
	_ ApiMethodType = iota
	ApiMethodTypePOST
	ApiMethodTypeGET
	ApiMethodTypePut
	ApiMethodTypeDelete
	ApiMethodTypeOptions
)

type ApiInfo struct {
	Path          string        //请求路径
	Method        ApiMethodType //请求类型  POST <GET <PUT < DELETE < OPTIONS < HEADER
	Name          string        //接口名称
	Alias         string        //别名(中文名)
	StartVersion  string        //起始版本
	EndVersion    string        //结束版本
	Version       string        //接口版本号
	Desc          string        //接口说明
	UpdateTime    int           //最后更新时间
	CreateTime    int           //创建时间
	RequestHeader DataModel     //请求参数体
	RequestBody   DataModel     //请求参数体
	ResponseBody  DataModel     //返回数据体

	Status bool //接口状态

}
