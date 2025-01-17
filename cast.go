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
	if Cache().isApiExist(api.Id) {
		return fmt.Errorf("model:%s e", api.Name)
	}
	Cache().setApi(api.Id)
	Cache().set(apiKey(api.Id), api)
	return nil
}

func apiKey(id int) string {
	return fmt.Sprintf("id:%d", id)
}

func Cast(apis []API, models []DataModel) ([]FileModelCast, []FileJavaCast, error) {
	err := Valid(apis, models)
	if err != nil {
		return nil, nil, err
	}
	list1, list2 := cast2FileModel(apis, models)
	return list1, list2, nil
}

func cast2FileModel(apis []API, models []DataModel) ([]FileModelCast, []FileJavaCast) {
	list := make([]FileModelCast, 0)
	list1 := make([]FileJavaCast, 0)
	for _, v := range apis {
		fm := castObjectiveApiFile(v)
		list = append(list, fm)
		list1 = append(list1, castJavaApi(v))
	}
	for _, v := range models {
		fm := castObjectiveModelFile(v)
		list = append(list, fm)
		if v.Type == "model" || v.Son {
			list1 = append(list1, castJavaModel(v))
		}
	}
	return list, list1
}

func Valid(apis []API, models []DataModel) error {
	Cache().flush()
	for _, a := range apis {
		if Cache().isApiExist(a.Id) {

		}
		addApi(a)
	}
	for _, a := range models {
		addModel(a)
	}
	for k := range Cache().Apis {
		a, ok := getApi(k)
		if !ok {
			return fmt.Errorf("apiId:%d not exist", k)
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

func getApi(id int) (API, bool) {
	i, b := ca.Cache.Get(apiKey(id))
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
	if typeS == "" {
		return true
	}
	switch strings.ToLower(typeS) {
	case "string", "float", "number", "integer", "boolean":
		return true
	case "nsstring", "nsuumber", "nsinteger", "bool":
		return true
	}
	return false
}

func Cast2JavaFileModel(apis []API, models []DataModel) ([]FileJavaCast, error) {
	err := Valid(apis, models)
	if err != nil {
		return nil, err
	}
	return cast2JavaModel(apis, models), nil
}

func cast2JavaModel(apis []API, models []DataModel) []FileJavaCast {
	list := make([]FileJavaCast, 0)
	for _, v := range apis {
		list = append(list, castJavaApi(v))
	}
	for _, v := range models {
		if v.Son {
			list = append(list, castJavaModel(v))
		}
	}

	return list
}
