package cast

import (
	"fmt"
	"testing"
)

func TestFormObjectiveAPIFileModel(t *testing.T) {
	api := makeTestApi1()
	list := makeModelList1()
	_, err := Cast([]Api{api}, list)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	a, _ := formObjectiveAPIFileModel(api)

	fmt.Println(a.H)
	fmt.Println("-----------------------------------")
	fmt.Println(a.M)

}
