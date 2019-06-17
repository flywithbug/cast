package cast

import (
	"encoding/json"
	"strings"
)

type APIMethodType int
type APIStatusType int

const (
	_ APIMethodType = iota
	ApiMethodTypePOST
	ApiMethodTypeGET
	ApiMethodTypePut
	ApiMethodTypeDelete
	ApiMethodTypeOptions
	ApiMethodTypeHeader
)

const (
	_ APIStatusType = iota
	APIStatusTypeUnDone
	APIStatusTypeDone
	APIStatusTypeOffline
	APIStatusTypeUnDefine
)

type Api struct {
	Id                int
	Action            string
	Path              string        //请求路径
	Method            APIMethodType //请求类型  POST <GET <PUT < DELETE < OPTIONS < HEADER
	Name              string        //接口名称
	Alias             string        //别名(中文名)
	Notes             string        //接口说明注释
	ParameterName     string        //请求参数体 key
	ResponseModelName string        //返回数据体 key
	Status            APIStatusType //接口状态
	Author            string
	AddTime           int
	UpDateTime        int
	NoneResponse      bool
}

func (a Api) ParameterModel() (DataModel, bool) {
	in, b := Cache().get(a.ParameterName)
	if !b {
		return DataModel{}, false
	}
	if m, ok := in.(DataModel); ok {
		return m, ok
	}
	return DataModel{}, false
}

func (a Api) ResponseModel() (DataModel, bool) {
	in, b := Cache().get(a.ResponseModelName)
	if !b {
		return DataModel{}, false
	}
	if m, ok := in.(DataModel); ok {
		return m, ok
	}
	return DataModel{}, false
}

func (a Api) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

func MethodType(methodType APIMethodType) string {
	switch methodType {
	case ApiMethodTypePOST:
		return "POST"
	case ApiMethodTypeGET:
		return "GET"
	case ApiMethodTypePut:
		return "PUT"
	case ApiMethodTypeDelete:
		return "DELETE"
	case ApiMethodTypeOptions:
		return "OPTIONS"
	case ApiMethodTypeHeader:
		return "HEADER"
	}
	return "UNDEFINE"
}

func StatusStringType(statusType string) APIStatusType {
	statusType = strings.ToUpper(statusType)
	switch statusType {
	case "UNDONE":
		return APIStatusTypeUnDone
	case "DONE":
		return APIStatusTypeDone
	case "OFFLINE":
		return APIStatusTypeOffline
	}
	return APIStatusTypeUnDefine
}

func MethodStringType(methodType string) APIMethodType {
	methodType = strings.ToUpper(methodType)
	switch methodType {
	case "POST":
		return ApiMethodTypePOST
	case "GET":
		return ApiMethodTypeGET
	case "PUT":
		return ApiMethodTypePut
	case "DELETE":
		return ApiMethodTypeDelete
	case "OPTIONS":
		return ApiMethodTypeOptions
	case "HEADER":
		return ApiMethodTypeHeader
	}
	return ApiMethodTypeGET
}

func StatusType(statusType APIStatusType) string {
	switch statusType {
	case APIStatusTypeUnDone:
		return "未完成"
	case APIStatusTypeDone:
		return "已完成"
	case APIStatusTypeOffline:
		return "已下线"
	}
	return "UNDEFINE"
}

func makeObjective_cFiles(api Api) (h, m string, err error) {

	return
}
