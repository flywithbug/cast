package cast

import "strings"

type Template struct {
	ApiH   string
	ApiM   string
	ModelH string
	ModelM string
}

func NewTemplate() Template {
	a, _ := Asset("template/api.h")
	b, _ := Asset("template/api.m")
	c, _ := Asset("template/model.h")
	d, _ := Asset("template/model.m")
	return Template{
		ApiH:   strings.ReplaceAll(string(a), "{*para*}", "%s"),
		ApiM:   strings.ReplaceAll(string(b), "{*para*}", "%s"),
		ModelH: strings.ReplaceAll(string(c), "{*para*}", "%s"),
		ModelM: strings.ReplaceAll(string(d), "{*para*}", "%s"),
	}
}
