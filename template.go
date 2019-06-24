package cast

import (
	"fmt"
	"strings"
)

type Template struct {
	ApiH           string
	ApiM           string
	ApiGetH        string
	ApiGetM        string
	ModelH         string
	ModelM         string
	BaseModelClass string
	BaseParaClass  string
	BaseApiClass   string
}

func NewTemplate() Template {
	a, _ := Asset("template/api.h")
	b, _ := Asset("template/api.m")
	a1, _ := Asset("template/api_get.h")
	b1, _ := Asset("template/api_get.m")
	c, _ := Asset("template/model.h")
	d, _ := Asset("template/model.m")
	e, _ := Asset("template/const")
	list := strings.Split(string(e), ",")
	if len(list) != 3 {
		//JYResponseModel,JYQueryParameter,JYMapiNetworkClient
		//baseModel,basePara,baseApi
		panic(fmt.Errorf("const base model needed"))
	}
	return Template{
		ApiH:           strings.ReplaceAll(string(a), "{*para*}", "%s"),
		ApiM:           strings.ReplaceAll(string(b), "{*para*}", "%s"),
		ApiGetH:        strings.ReplaceAll(string(a1), "{*para*}", "%s"),
		ApiGetM:        strings.ReplaceAll(string(b1), "{*para*}", "%s"),
		ModelH:         strings.ReplaceAll(string(c), "{*para*}", "%s"),
		ModelM:         strings.ReplaceAll(string(d), "{*para*}", "%s"),
		BaseModelClass: list[0],
		BaseParaClass:  list[1],
		BaseApiClass:   list[2],
	}
}
