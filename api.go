package cast

import (
	"encoding/json"
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
)

type Api struct {
	Id            int
	Path          string        //请求路径
	Method        APIMethodType //请求类型  POST <GET <PUT < DELETE < OPTIONS < HEADER
	Name          string        //接口名称
	Alias         string        //别名(中文名)
	Notes         string        //接口说明注释
	Parameter     DataModel     //请求参数体
	ResponseModel DataModel     //返回数据体
	Status        APIStatusType //接口状态
	Author        string
	AddTime       int
	UpDateTime    int
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
