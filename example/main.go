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

func main() {
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
