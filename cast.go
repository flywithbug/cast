package cast

import (
	"fmt"
)

func AddModel(model DataModel) bool {
	if Cache().Exist(model.Name) {
		return false
	}
	Cache().Set(model.Name, model)
	Cache().SetModel(model.Name)
	return true
}

func AddApi(api Api) bool {
	if Cache().Exist(api.Name) {
		return false
	}
	Cache().Set(api.Name, api)
	Cache().SetApi(api.Name)
	return true
}

func exist(key string) bool {
	return Cache().Exist(key)
}

func Valid() error {
	for k, _ := range Cache().Apis {
		a, b := GetApi(k)
		if !b {
			return fmt.Errorf("api:%s not exist", k)
		}
		if !exist(a.ResponseModelName) {
			return fmt.Errorf("model:%s not exist", a.ResponseModelName)
		}
		if !exist((a.ParameterName)) {
			return fmt.Errorf("model:%s not exist", (a.ParameterName))

		}
	}
	for k, _ := range Cache().Models {
		m, b := GetDataModel(k)
		if !b {
			return fmt.Errorf("model:%s not exist", k)

		}
		if m.ParentName != "" && m.ParentName != Conf().FileModel.BaseModel && !exist(m.ParentName) {
			return fmt.Errorf("model:%s not exist", m.ParentName)
		}
		for _, a := range m.Attributes {
			if a.Type == "object" || a.Type == "array" {
				if !exist(a.Name) {
					return fmt.Errorf("model:%s not exist", a.Name)

				}
			}
		}
	}
	return nil
}

func GetApi(key string) (Api, bool) {
	i, b := ca.Cache.Get(key)
	if b {
		if a, ok := i.(Api); ok {
			return a, ok
		}
	}
	return Api{}, false
}

func GetDataModel(key string) (DataModel, bool) {
	i, b := ca.Cache.Get(key)
	if b {
		if a, ok := i.(DataModel); ok {
			return a, ok
		}

	}
	return DataModel{}, false
}
