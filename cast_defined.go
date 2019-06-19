package cast

import (
	"fmt"
	"strings"
)

var (
	header = "//\n//  %s.h\n//  JYCastKit\n//\n//  Created by Flywithbug on 2019.6.18.\n//  " +
		"Copyright © 2019 hellobike." +
		" All rights reserved.\n//\n"

	importStr  = "#import <JYCastKit/JYCastKit.h>"
	importStr1 = "#import \"%s.h\""

	interF = "@interface %s : JYResponseModel"

	interParaF = "@interface %s : JYQueryParameter"

	implementation = "@implementation %s"

	stringProperty    = "@property (nonatomic, copy)  \tNSString *%s;"
	boolProperty      = "@property (nonatomic, assign)\tBOOL %s;"
	numberProperty    = "@property (nonatomic, strong)\tNSNumber *%s;"
	objectProperty    = "@property (nonatomic, strong)\t%s *%s;"
	arrayProperty     = "@property (nonatomic, copy)  \tNSArray <%s *> *%s;"
	containerProperty = "@\"%s\" : [%s class]"

	//categoryName      = "JYMapiNetworkClient+%s" //.h.m
	//categoryInterface = "@interface JYMapiNetworkClient (%s)"

	categoryMethod = "- (nullable NSURLSessionDataTask *)mapi_%s_withmoduleType:(JYModuleType)type" +
		"\n\t\tpara:(%s * _Nullable)param" +
		"\n\t\tprogress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress" +
		"\n\t\tsuccess:(nullable void (^)(NSURLSessionDataTask *_Nullable task, %s * _Nullable responseObject))success" +
		"\n\t\tfailure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;"

	categoryImpMethod = "- (nullable NSURLSessionDataTask *)mapi_%s_withmoduleType:(JYModuleType)type" +
		"\n\t\tpara:(%s * _Nullable)param" +
		"\n\t\tprogress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress" +
		"\n\t\tsuccess:(nullable void (^)(NSURLSessionDataTask *_Nullable task, %s * _Nullable responseObject))success" +
		"\n\t\tfailure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure{" +
		"\n\tparam.action = @\"%s\";" +
		"\n\treturn [self mapi_post:[self moduleApiServer:type] parameters:param repClass:[%s class] progress:uploadProgress success:success failure:failure];" +
		"\n}"

	//apiAction = "- (NSString *)action{\n\treturn @\"%s\";\n}"
)

//模型名称生产
func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}

//模型名称生产
func formatParaName(name string) string {
	return fmt.Sprintf("JYPara%s", name)
}

//api名称生产
func formatApiCategoryName(action string) string {
	action = strings.Replace(action, "user.ev.", "", -1)
	action = strings.Replace(action, ".", "_", -1)
	return action
}

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func strHumpToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return capitalize(temp[0]) + upperStr
}

func capitalize(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func formatDefaultClass(typeee string) string {
	switch strings.ToLower(typeee) {
	case "string":
		return "NSString"
	case "bool", "boolean":
		return "NSNumber"
	case "number", "float":
		return "NSNumber"

	}
	return typeee
}
