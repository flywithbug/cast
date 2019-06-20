package cast

import (
	"fmt"
	"testing"
)

func TestCastObjectiveApiFile(t *testing.T) {

	api := API{
		Id:             100,
		Name:           "checkBlackUser",
		Action:         "user.ev.checkBlackUser",
		Method:         "post",
		Path:           "api/",
		ParameterName:  "JYParaCheckBlackUser",
		ResponseName:   "JYRespCheckBlackUser",
		ParaOriginName: "CheckBlackUser",
		RespOriginName: "CheckBlackUser",
		NoResponse:     false,
		NoParameter:    false,
		CategoryClass:  "JYMapiNetworkClient",
		AddTime:        1561013457,
		UpDateTime:     1561013457,
		Author:         "me",
		Status:         "done",
	}
	fi1 := castObjectiveApiFile(api)
	fmt.Println(fi1.H)
	fmt.Println(fi1.M)

}
