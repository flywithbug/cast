package utils

import "fmt"

func Test()  {



	testBool()
	testMd5()

}

func testBool()  {
	fmt.Println(TOBool("abc"))

	fmt.Println(TOBool("true"))

	fmt.Println(TOBool("false"))

	fmt.Println(TOBool("0"))

	fmt.Println(TOBool("1"))
}


func testMd5(){
	Md5("abc")

}

