package main

import (
	"fmt"

	"github.com/flywithbug/cast"
)

func main() {
	cast.Cache().Set("a", "b")
	fmt.Println(cast.Cache().Get("a"))
}
