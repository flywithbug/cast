package cast

import (
	"fmt"
	"testing"
)

func TestCastObjectiveApiFile(t *testing.T) {

	api := API{
		Id:            100,
		Name:          "checkBlackUser",
		Action:        "user.ev.checkBlackUser",
		Method:        "post",
		Path:          "api/",
		ParameterName: "JYParaCheckBlackUser",
		ResponseName:  "JYRespCheckBlackUser",
		//ParaOriginName: "CheckBlackUser",
		//RespOriginName: "CheckBlackUser",
		NoResponse:    false,
		NoParameter:   false,
		CategoryClass: "JYMapiNetworkClient",
		AddTime:       1561013457,
		UpDateTime:    1561013457,
		Author:        "me",
		Status:        "done",
	}
	fi1 := castObjectiveApiFile(api)
	fmt.Println(fi1.H)
	fmt.Println(fi1.M)
}

func TestCastJavaApiFile(t *testing.T) {

	a1 := Attribute{
		Name:     "name",
		Type:     "string",
		Required: true,
	}
	a2 := Attribute{
		Name:     "age",
		Type:     "integer",
		Required: true,
	}
	a22 := Attribute{
		Name:      "features",
		Type:      "array",
		ModelName: "integer",
		Required:  true,
	}

	a3 := Attribute{
		Name:      "list",
		Type:      "array",
		ModelName: "integer",
		Required:  true,
	}
	a33 := Attribute{
		Name:      "people",
		Type:      "object",
		ModelName: "JYParaPeople111",
		Required:  true,
	}

	//md1 := DataModel{
	//	//OriginName: "people",
	//	Name: "JYParaPeople111",
	//	Type: "Model",
	//	Attributes: []Attribute{
	//		a1, a2, a3, a22,
	//	},
	//}
	md := DataModel{
		//OriginName: "people",
		Name: "JYParaPeople",
		Type: "Model",
		Attributes: []Attribute{
			a1, a2, a3, a22, a33,
		},
	}

	api := API{
		Id:            100,
		Name:          "checkBlackUser",
		Action:        "user.ev.checkBlackUser",
		Method:        "post",
		Path:          "api/",
		ParameterName: "JYParaPeople",
		ResponseName:  "JYRespPeople",
		//ParaOriginName: "CheckBlackUser",
		//RespOriginName: "CheckBlackUser",
		NoResponse:     false,
		NoParameter:    false,
		CategoryClass:  "JYMapiNetworkClient",
		AddTime:        1561013457,
		UpDateTime:     1561013457,
		Author:         "me",
		Status:         "done",
		ParameterModel: md,
	}
	fi := castJavaApi(api)
	fi2 := castJavaModel(md)
	fmt.Println(fi)
	fmt.Println(fi2)
}
