package main

import (
	"fmt"

	"github.com/flywithbug/cast"
	"github.com/flywithbug/file_manager"
)

func makeDaModel() cast.DataModel {
	model0 := cast.DataModel{
		Name: "GetBikesByLocation",
		Type: cast.ModelTypeTypeModel,
		Attributes: []cast.Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}
	model := cast.DataModel{
		Name: "GetBikesByLocationBikes",
		Type: cast.ModelTypeTypeModel,
		Attributes: []cast.Attribute{
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo2", Type: "int", Notes: "车辆编号"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "bikeNo3", Type: "bool", Notes: "车辆编号"},
			{Name: "QueryPara", Type: "Object", DataModel: model0,
				ModelName: "GetBikesByLocation"},
			{Name: "list", Type: "array", DataModel: model0,
				ModelName: "GetBikesByLocation"},
		},
	}
	return model
}
func oldMakeObcFilllllll() {
	obm := cast.CastModelObjective_C_H_M(makeDaModel())
	fmt.Println(obm.H)
	fmt.Println("___________________________")

	fmt.Println(obm.M)
	hName := fmt.Sprintf("%s.h", obm.ModelName)
	err := file_manager.WriteFileString(hName, obm.H, true)
	if err != nil {
		fmt.Println(err)
	}
	mName := fmt.Sprintf("%s.m", obm.ModelName)
	err = file_manager.WriteFileString(mName, obm.M, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	makeFiles()
}

func makeFiles() {
	api := makeTestApi1()
	list := makeModelList1()
	err := cast.Cast([]cast.Api{api}, list)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	for _, v := range list {
		makeObcModelFile(v)
	}
	makeObcApiFile(api)
}

func makeObcApiFile(model cast.Api) {
	obm := cast.CastApiObjective_C_H_M(model)
	hName := fmt.Sprintf("%s.h", obm.ApiFullName)
	err := file_manager.WriteFileString(hName, obm.H, true)
	if err != nil {
		fmt.Println(err)
	}
	mName := fmt.Sprintf("%s.m", obm.ApiFullName)
	err = file_manager.WriteFileString(mName, obm.M, true)
	if err != nil {
		fmt.Println(err)
	}
}

func makeObcModelFile(model cast.DataModel) {
	obm := cast.CastModelObjective_C_H_M(model)
	hName := fmt.Sprintf("%s.h", obm.ModelName)
	err := file_manager.WriteFileString(hName, obm.H, true)
	if err != nil {
		fmt.Println(err)
	}
	mName := fmt.Sprintf("%s.m", obm.ModelName)
	err = file_manager.WriteFileString(mName, obm.M, true)
	if err != nil {
		fmt.Println(err)
	}
}

func makeModelList1() []cast.DataModel {
	model0 := cast.DataModel{
		Name: "GetBikesByLocation",
		Type: cast.ModelTypeTypeModel,
		Attributes: []cast.Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}
	model := cast.DataModel{
		Name: "GetBikesByLocationBikes",
		Type: cast.ModelTypeTypeModel,
		Attributes: []cast.Attribute{
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo2", Type: "int", Notes: "车辆编号"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "bikeNo3", Type: "bool", Notes: "车辆编号"},
			{Name: "QueryPara", Type: "Object", DataModel: model0,
				ModelName: "GetBikesByLocation"},
			{Name: "list", Type: "array", DataModel: model0,
				ModelName: "GetBikesByLocation"},
		},
	}
	model1 := cast.DataModel{
		Name: "QueryPara",
		Type: cast.ModelTypeTypeParameter,
		Attributes: []cast.Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}

	list := make([]cast.DataModel, 0)
	list = append(list, model0, model, model1)
	return list
}

func makeTestApi1() cast.Api {
	api := cast.Api{
		Id:                993,
		Name:              "user.ev.ride.getBikesByLoction",
		Action:            "user.ev.ride.getBikesByLoction",
		Alias:             "获取附近电动车",
		Notes:             "根据位置获取附近电动车，目前获取附近500m内的200辆以内车辆",
		Method:            cast.ApiMethodTypePOST,
		Path:              "/api",
		Author:            "小哈",
		AddTime:           1531222996,
		UpDateTime:        1558078875,
		Status:            cast.APIStatusTypeDone,
		ParameterName:     "QueryPara",
		ResponseModelName: "GetBikesByLocationBikes",
	}
	return api
}
