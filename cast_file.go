package cast

import (
	"fmt"
	"strings"
	"time"
)

/*
//
//  JYRespUser_ev_ride_getBikesByLocationModel.h
//  JYCastKit
//
//  Created by Flywithbug on 2019/6/14.
//  Copyright © 2019 hellobike. All rights reserved.
//
*/
var (
	header = "//\n//  %s.h\n//  JYCastKit\n//\n//  Created by Flywithbug on %s.\n//  Copyright © 2019 hellobike. All rights reserved.\n//\n"

	interF = "@interface %s : JYResponseModel"

	stringProperty = "@property (nonatomic, copy)  \tNSString *%s"
	boolProperty   = "@property (nonatomic, assign)\tBOOL %s"
)

func CastModelObjective_C_H(model DataModel) (str string, err error) {
	mName := formatModelName(model.Name)
	now := time.Now().Format("2006-01-02")
	str = fmt.Sprintf(header, mName, now)
	str += "\t\n"
	str = fmt.Sprintf(str+interF, mName)
	str += "\n"
	str += formatAttributes(model.Attributes)
	return
}

func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}

func formatAttributes(list []Attribute) (str string) {
	for _, v := range list {
		switch strings.ToLower(v.Type) {
		case "string":
			str += fmt.Sprintf(stringProperty, v.Name)
			str += "\n"
		case "bool", "boolean":
			str += fmt.Sprintf(boolProperty, v.Name)
			str += "\n"
		}
	}
	return str
}
