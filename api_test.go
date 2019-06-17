package cast

import (
	"testing"
)

func makeLocationBikeApi() Api {
	//model := DataModel{
	//	Name: "JYRespUser_ev_ride_getBikesByLocationBikesModel",
	//	Type: "Object",
	//	Attributes: []Attribute{
	//		{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
	//		{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
	//		{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
	//	},
	//}
	//model1 := DataModel{
	//	Name: "JYQueryParaUser_ev_ride_getBikesByLocation",
	//	Type: "Object",
	//	Attributes: []Attribute{
	//		{Name: "cityCode", Type: "string", Notes: "城市码"},
	//		{Name: "lng", Type: "number", Notes: "经度"},
	//		{Name: "lat", Type: "number", Notes: "纬度"},
	//		{Name: "adCode", Type: "string", Notes: "县区代码"},
	//		{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
	//	},
	//}

	api := Api{
		Id:                993,
		Name:              "user.ev.ride.getBikesByLoction",
		Alias:             "获取附近电动车",
		Notes:             "根据位置获取附近电动车，目前获取附近500m内的200辆以内车辆",
		Method:            ApiMethodTypePOST,
		Path:              "/api",
		Author:            "小哈",
		AddTime:           1531222996,
		UpDateTime:        1558078875,
		Status:            APIStatusTypeDone,
		ResponseModelName: "",
		ParameterName:     "",
	}
	return api
}

func makeModelList() []DataModel {
	model0 := DataModel{
		Name: "getBikesByLocation",
		Type: ModelTypeTypeModel,
		Attributes: []Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}
	model := DataModel{
		Name: "getBikesByLocationBikesModel",
		Type: ModelTypeTypeModel,
		Attributes: []Attribute{
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "QueryPara", Type: "Object", DataModel: model0},
		},
	}
	model1 := DataModel{
		Name: "QueryPara",
		Type: ModelTypeTypeParameter,
		Attributes: []Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}

	list := make([]DataModel, 0)
	list = append(list, model0, model, model1)
	return list
}

func makeTestApi() Api {
	api := Api{
		Id:                993,
		Name:              "user.ev.ride.getBikesByLoction",
		Alias:             "获取附近电动车",
		Notes:             "根据位置获取附近电动车，目前获取附近500m内的200辆以内车辆",
		Method:            ApiMethodTypePOST,
		Path:              "/api",
		Author:            "小哈",
		AddTime:           1531222996,
		UpDateTime:        1558078875,
		Status:            APIStatusTypeDone,
		ParameterName:     "JYQueryParaUser_ev_ride_getBikesByLocation",
		ResponseModelName: "getBikesByLocation",
	}
	return api
}

func TestApiModel(t *testing.T) {

}

func makeFiles(api Api) (h, m string, err error) {
	//timeTemplate3 := "2006-01-02" //其他类型
	//now := time.Now().Format(timeTemplate3)
	//modelHeader := fmt.Sprintf("//\n//  %s.h\n//  JYCastKit\n//\n//  Created by Flywithbug on %s.\n//  Copyright © 2019 hellobike. All rights reserved.\n//", api.ResponseModelName, now)
	//fmt.Println(modelHeader)
	//
	//file_manager.WriteFileString("./test.h", modelHeader, true)
	return
}

func modleList(model DataModel) []DataModel {
	list := make([]DataModel, 0)
	list = append(list, model)
	for _, v := range model.Attributes {
		if v.DataModel.Name != "" {
			l := modleList(v.DataModel)
			list = append(list, l...)
		}
	}
	return list
}
