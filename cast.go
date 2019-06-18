package cast

import (
	"fmt"
	"strings"

	"github.com/flywithbug/file_manager"
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

func Cast(apis []Api, models []DataModel) ([]ObjectiveCFileModel, error) {
	for _, a := range apis {
		addApi(a)
	}
	for _, a := range models {
		addModel(a)
	}
	if err := valid(); err != nil {
		return nil, err
	}
	list := cast2ObjectiveCFiles(apis, models)
	return list, nil
}

func cast2ObjectiveCFiles(apis []Api, models []DataModel) []ObjectiveCFileModel {
	list := make([]ObjectiveCFileModel, 0)

	for _, v := range apis {
		obm := CastApiObjective_C_H_M(v)
		if obm != nil {
			//fmt.Println(v.Action)
			continue
		}
		obF := ObjectiveCFileModel{
			Name: obm.ApiFullName,
			H:    obm.H,
			M:    obm.M,
		}
		list = append(list, obF)
	}

	for _, v := range models {
		obm := CastModelObjective_C_H_M(v)
		obF := ObjectiveCFileModel{
			Name: obm.ModelName,
			H:    obm.H,
			M:    obm.M,
		}
		list = append(list, obF)
	}

	return list
}

func Cast2Files(apis []Api, models []DataModel, path string) error {
	list, err := Cast(apis, models)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	for _, v := range list {
		hName := fmt.Sprintf("%s.h", v.Name)
		err := file_manager.WriteFileString(path+hName, v.H, true)
		if err != nil {
			return err
		}
		mName := fmt.Sprintf("%s.m", v.Name)
		err = file_manager.WriteFileString(path+mName, v.M, true)
		if err != nil {
			return err
		}
	}
	return nil
}
