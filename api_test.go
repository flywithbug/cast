package cast

import (
	"fmt"
	"testing"
)

func TestApiModel(t *testing.T) {

	model := DataModel{
		Name: "JYRespUser_ev_ride_getBikesByLocationBikesModel",
		Type: "Object",
		Attributes: []Attribute{
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
		},
	}
	model1 := DataModel{
		Name: "JYQueryParaUser_ev_ride_getBikesByLocation",
		Type: "Object",
		Attributes: []Attribute{
			{Name: "cityCode", Type: "string", Notes: "城市码"},
			{Name: "lng", Type: "number", Notes: "经度"},
			{Name: "lat", Type: "number", Notes: "纬度"},
			{Name: "adCode", Type: "string", Notes: "县区代码"},
			{Name: "redPacketModel", Type: "boolean", Notes: "是否红包车"},
		},
	}

	api := Api{
		Id:            993,
		Name:          "user.ev.ride.getBikesByLoction",
		Alias:         "获取附近电动车",
		Notes:         "根据位置获取附近电动车，目前获取附近500m内的200辆以内车辆",
		Method:        ApiMethodTypePOST,
		Path:          "/api",
		Author:        "小哈",
		AddTime:       1531222996,
		UpDateTime:    1558078875,
		Status:        APIStatusTypeDone,
		ResponseModel: model,
		Parameter:     model1,
	}

	fmt.Println(api.ToJson())

}
