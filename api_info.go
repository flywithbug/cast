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
	ApiId         int           //Api服务端Id
	Path          string        //请求路径
	Method        ApiMethodType //请求类型  POST <GET <PUT < DELETE < OPTIONS < HEADER
	Title         string        //原名字段 /*user.ev.ride.canParkBike(能否还车)*/
	Name          string        //接口名称  从Title中分离出来 /* user.ev.ride.canParkBike */
	Alias         string        //别名(中文名)  从Title中分离出来 /* 能否还车 */
	Service       string        //关联服务
	Version       string        //接口版本号
	Desc          string        //接口说明
	UpdateTime    int           //最后更新时间
	CreateTime    int           //创建时间
	Author        string        //创建人
	AuthorId      int           //创建人
	RequestBody   DataModel     //请求参数体
	RequestHeader DataModel     //请求参数体
	ResponseBody  DataModel     //返回数据体
	StartVersion  string        //起始版本
	EndVersion    string        //结束版本
	Status        bool          //接口状态
	ProjectId     int           //所属项目组Id  例如助力车活动 http://yapi.hellobike.cn/project/1843/interface/api
	CatId         int           //业务线Id  例如骑行业务线  http://yapi.hellobike.cn/project/1843/interface/api/cat_7036
}
