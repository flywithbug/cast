package cast

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

func Valid() bool {
	for k, _ := range Cache().Apis {
		_, b := GetApi(k)
		if !b {
			return false
		}
	}
	for k, _ := range Cache().Models {
		_, b := GetDataModel(k)
		if !b {
			return false
		}
	}
	return true
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
