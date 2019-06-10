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
	Title   string
	Method  ApiMethodType //方法请求类型
	Path    string
	Service string
}
