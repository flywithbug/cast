package cast

import (
	"fmt"
)

func addModel(model DataModel) bool {
	if Cache().exist(model.Name) {
		return false
	}
	Cache().set(model.Name, model)
	Cache().setModel(model.Name)
	return true
}

func addApi(api Api) bool {
	if Cache().exist(api.Name) {
		return false
	}
	Cache().set(api.Name, api)
	Cache().setApi(api.Name)
	return true
}

func exist(key string) bool {
	return Cache().exist(key)
}

func valid() error {
	for k, _ := range Cache().Apis {
		a, b := getApi(k)
		if !b {
			return fmt.Errorf("api:%s not exist", k)
		}
		if !a.NoResponse && !exist(a.ResponseModelName) {
			return fmt.Errorf("model:%s not exist", a.ResponseModelName)
		}
		if !a.NoParameter && !exist((a.ParameterName)) {
			return fmt.Errorf("model:%s not exist", (a.ParameterName))
		}
	}
	for k, _ := range Cache().Models {
		m, b := getDataModel(k)
		if !b {
			return fmt.Errorf("model:%s not exist", k)
		}
		if m.ParentName != "" && m.ParentName != Conf().FileModel.BaseModel && !exist(m.ParentName) {
			return fmt.Errorf("model:%s not exist", m.ParentName)
		}
		for _, a := range m.Attributes {
			if a.Type == "object" || a.Type == "array" {
				if !exist(a.ModelName) {
					return fmt.Errorf("model:%s not exist", a.ModelName)

				}
			}
		}
	}
	return nil
}

func getApi(key string) (Api, bool) {
	i, b := ca.Cache.Get(key)
	if b {
		if a, ok := i.(Api); ok {
			return a, ok
		}
	}
	return Api{}, false
}

func getDataModel(key string) (DataModel, bool) {
	i, b := ca.Cache.Get(key)
	if b {
		if a, ok := i.(DataModel); ok {
			return a, ok
		}
	}
	return DataModel{}, false
}

//------------------------------------------------------------------------------------------------------------//
//------------------------------------------------------------------------------------------------------------//

func Cast(apis []Api, models []DataModel) error {
	for _, a := range apis {
		addApi(a)
	}
	for _, a := range models {
		addModel(a)
	}
	if err := valid(); err != nil {
		return err
	}
	return nil
}

//func Cast2ObjectiveCFiles()[]ObjectiveCFileModel  {
//
//}
