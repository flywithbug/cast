package cast

import (
	"fmt"
	"testing"
)

func TestCastObjectiveModelFile(t *testing.T) {
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
		ModelName: "string",
		Required:  true,
	}

	a3 := Attribute{
		Name:      "list",
		Type:      "array",
		ModelName: "integer",
		Required:  true,
	}

	a4 := Attribute{
		Name:      "people",
		ModelName: "People",
		Type:      "object",
		Required:  true,
	}

	a5 := Attribute{
		Name:      "list",
		ModelName: "People",
		Type:      "array",
		Required:  true,
	}
	md := DataModel{
		Name:       "People",
		Type:       "Model",
		ParentName: "JYResponseModel",
		Attributes: []Attribute{
			a1, a2, a3, a22,
		},
	}
	md1 := DataModel{
		Name:       "People1",
		Type:       "Model",
		ParentName: "JYResponseModel",
		Attributes: []Attribute{
			a1, a2, a4, a5,
		},
	}

	fi1 := castObjectiveModelFile(md)
	fi2 := castObjectiveModelFile(md1)
	fmt.Println(fi1.H)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(fi1.M)
	fmt.Println("--------------------------------------------------------------------------------")

	fmt.Println(fi2.H)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(fi2.M)

}

func TestAttributeList(t *testing.T) {
	str1, str2, str3 := attributesFormat(makeAttlist())
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

}

func makeAttlist() []Attribute {
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

	a3 := Attribute{
		Name:     "gender",
		Type:     "integer",
		Required: true,
	}

	a4 := Attribute{
		Name:      "people",
		ModelName: "People",
		Type:      "object",
		Required:  true,
	}
	return []Attribute{a1, a2, a3, a4}
}
