package main

import (
	"fmt"

	"errors"
	"os"

	"github.com/flywithbug/cast"
	"github.com/flywithbug/file_manager"
)

func main() {

	CastModelFile()

}

func CastModelFile() {
	a1 := cast.Attribute{
		Name:     "name",
		Type:     "string",
		Required: true,
	}
	a2 := cast.Attribute{
		Name:     "age",
		Type:     "integer",
		Required: true,
	}

	a3 := cast.Attribute{
		Name:     "gender",
		Type:     "integer",
		Required: true,
	}

	a4 := cast.Attribute{
		Name:      "people",
		ModelName: "JYParaPeople",
		Type:      "object",
		Required:  true,
	}
	a5 := cast.Attribute{
		Name:      "list",
		ModelName: "JYParaPeople",
		Type:      "array",
		Required:  true,
	}

	md := cast.DataModel{
		Name:       "JYParaPeople",
		Type:       "Model",
		ParentName: "JYQueryParameter",
		Attributes: []cast.Attribute{
			a1, a2, a3,
		},
	}
	md1 := cast.DataModel{
		Name:       "JYRespPeopleData",
		Type:       "Model",
		ParentName: "JYResponseModel",
		Attributes: []cast.Attribute{
			a1, a2, a4, a5,
		},
	}

	api := cast.API{
		Id:             100,
		Name:           "checkBlackUser",
		Action:         "user.ev.checkBlackUser",
		Method:         "post",
		Path:           "api/",
		ParameterName:  "JYParaPeople",
		ResponseName:   "JYRespPeopleData",
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
	list, err := cast.Cast([]cast.API{api}, []cast.DataModel{md, md1})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range list {
		temPath := ""
		if v.Type == "model" {
			temPath += "model/"
		} else {
			temPath += "api/"
		}
		hName := fmt.Sprintf("%s%s.h", temPath, v.Name)
		mName := fmt.Sprintf("%s%s.m", temPath, v.Name)
		err := WriteFileString(hName, v.H, true)
		if err != nil {
			panic(err)
		}
		err = WriteFileString(mName, v.M, true)
		if err != nil {
			panic(err)
		}
	}

}

func WriteFileString(path, content string, cover bool) error {
	if !cover {
		ex, _ := file_manager.PathExists(path)
		if ex {
			return errors.New("PathExists")
		}
	}

	fil, err := os.Create(path)
	if err != nil {
		err = file_manager.CreatePath(path)
		if err != nil {
			return err
		}
		return WriteFileString(path, content, cover)
	}
	_, err = fil.WriteString(content)
	return err
}
