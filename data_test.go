package cast

import (
	"fmt"
	"testing"
)

func TestAssetInfo(t *testing.T) {
	b, err := Asset("template/api.h")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
