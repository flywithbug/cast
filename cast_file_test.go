package cast

import (
	"fmt"
	"testing"
)

func makeDaModel() DataModel {
	model0 := DataModel{
		Name: "GetBikesByLocation",
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
		Name: "GetBikesByLocationBikes",
		Type: ModelTypeTypeModel,
		Attributes: []Attribute{
			{Name: "bikeNo", Type: "string", Notes: "车辆编号"},
			{Name: "bikeNo2", Type: "int", Notes: "车辆编号"},
			{Name: "bikeNo3", Type: "bool", Notes: "车辆编号"},
			{Name: "QueryPara", Type: "Object", DataModel: model0},
		},
	}
	return model
}

func TestCastObjective_C_H(t *testing.T) {
	r, _ := CastModelObjective_C_H(makeDaModel())
	fmt.Println(r)
}
