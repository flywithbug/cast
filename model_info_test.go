package cast

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	m := DataModel{}
	list := make([]Attribute, 0)
	a := Attribute{Name: "b"}
	list = append(list, a)
	a = Attribute{Name: "d"}
	list = append(list, a)
	a = Attribute{Name: "a"}
	list = append(list, a)
	a = Attribute{Name: "c"}
	list = append(list, a)
	a = Attribute{Name: "k"}
	list = append(list, a)

	a = Attribute{Name: "z"}
	list = append(list, a)

	a = Attribute{Name: "l"}
	list = append(list, a)

	m.Attributes = list

	for _, v := range m.Attributes {
		fmt.Println(v.Name)
	}

	m.SortAttributesByName()

	fmt.Println("===-=-=-===")
	for _, v := range m.Attributes {
		fmt.Println(v.Name)
	}
}
