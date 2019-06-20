package cast

import (
	"fmt"
	"strings"
)

var (
	temp = NewTemplate()
)

func addModel(model DataModel) error {
	if Cache().isModelExist(model.Name) {
		return fmt.Errorf("model:%s e", model.Name)
	}
	Cache().setModel(model.Name)
	Cache().set(model.Name, model)
	return nil
}

func addApi(api API) error {
	if Cache().isApiExist(api.Name) {
		return fmt.Errorf("model:%s e", api.Name)
	}
	Cache().setApi(api.Name)
	Cache().set(api.Name, api)
	return nil
}

func Cast(apis []API, models []DataModel) ([]FileModelCast, error) {
	err := Valid(apis, models)
	if err != nil {
		return nil, err
	}
	return cast2FileModel(apis, models), nil
}

func cast2FileModel(apis []API, models []DataModel) []FileModelCast {
	list := make([]FileModelCast, 0)
	for _, v := range apis {
		fm := castObjectiveApiFile(v)
		list = append(list, fm)
	}
	for _, v := range models {
		fm := castObjectiveModelFile(v)
		list = append(list, fm)
	}
	return list
}

func Valid(apis []API, models []DataModel) error {
	Cache().flush()
	for _, a := range apis {
		addApi(a)
	}
	for _, a := range models {
		addModel(a)
	}
	for k := range Cache().Apis {
		a, ok := getApi(k)
		if !ok {
			return fmt.Errorf("api:%s not exist", k)
		}
		if !a.NoParameter && !Cache().isModelExist(a.ParameterName) {
			return fmt.Errorf("action:%s,Para:%s", a.Action, a.ParameterName)
		}
		if !a.NoResponse && !Cache().isModelExist(a.ResponseName) {
			return fmt.Errorf("action:%s,Resp:%s", a.Action, a.ParameterName)
		}
	}
	for k := range Cache().Models {
		m, ok := getDataModel(k)
		if !ok {
			return fmt.Errorf("model:%s not exist", k)
		}
		if m.ParentName != "" && m.ParentName != temp.BaseParaClass &&
			m.ParentName != temp.BaseModelClass &&
			!Cache().isModelExist(m.ParentName) {
			return fmt.Errorf("model:%s :paren:%s' not exist", m.Name, m.ParentName)
		}
		for _, a := range m.Attributes {
			if a.Type == "object" || a.Type == "array" {
				if !IsDefaultType(a.ModelName) && !Cache().isModelExist(a.ModelName) {
					return fmt.Errorf("model:%s  AttModel:%s", m.Name, a.ModelName)
				}
			}
		}
	}
	return nil
}

func getApi(key string) (API, bool) {
	i, b := ca.Cache.Get(key)
	if b {
		if a, ok := i.(API); ok {
			return a, ok
		}
	}
	return API{}, false
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

func IsDefaultType(typeS string) bool {
	switch strings.ToLower(typeS) {
	case "string", "float", "number", "integer", "boolean":
		return true
	}
	return false
}
