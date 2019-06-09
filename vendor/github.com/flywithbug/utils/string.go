package utils



import (
	"fmt"
	"math/rand"
	"strconv"

)

const DefaultCharset = Alphanumeric + "!%@#"

// RandString 产生随机字符串，可用于密码等
func RandString(n int, defaultCharsets ...string) string {
	charset := DefaultCharset
	if len(defaultCharsets) > 0 {
		charset = defaultCharsets[0]
	}

	result, err := String(n, charset)
	if err != nil {
		fmt.Println("goutils RandString error:", err)
		return Md5(strconv.Itoa(rand.Intn(999999)))
	}

	return result
}

func RandSMSString(n int, defaultCharsets ...string) string {
	charset := Alphanumeric
	if len(defaultCharsets) > 0 {
		charset = defaultCharsets[0]
	}

	result, err := String(n, charset)
	if err != nil {
		fmt.Println("goutils RandString error:", err)
		return Md5(strconv.Itoa(rand.Intn(999999)))
	}

	return result
}
//func ConvertInt2String(v interface{})string  {
//	switch v.(type) {
//	case int:
//		return strconv.Itoa(int(v))
//	case string:
//		return string(v)
//	case int64:
//		return strconv.FormatInt(int64(v),10)
//	default:
//		return ""
//	}
//}

