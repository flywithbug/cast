package cast

import (
	"fmt"
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
)

func CastModelObjective_C_H(model DataModel) (str string, err error) {
	mName := formatModelName(model.Name)
	now := time.Now().Format("2006-01-02")
	str = fmt.Sprintf(header, mName, now)
	str += "\t\n"
	str = fmt.Sprintf(str+interF, mName)

	return
}

func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}

func formatAttributes(list []Attribute) string {

	return ""
}
