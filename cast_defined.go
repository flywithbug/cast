package cast

import (
	"fmt"
)

var (
	header = "//\n//  %s.h\n//  JYCastKit\n//\n//  Created by Flywithbug on %s.\n//  " +
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

	categoryName      = "JYMapiNetworkClient+%s" //.h.m
	categoryInterface = "@interface JYMapiNetworkClient (%s)"

	categoryMethod = "- (nullable NSURLSessionDataTask *)mapi_%s_withmoduleType:(JYModuleType)type" +
		"\n\tpara:(%s * _Nullable)param" +
		"\n\tprogress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress" +
		"\n\tsuccess:(nullable void (^)(NSURLSessionDataTask *_Nullable task, %s * _Nullable responseObject))success" +
		"\n\tfailure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;"
)

//模型名称生产
func formatModelName(name string) string {
	return fmt.Sprintf("JYResp%s", name)
}

//模型名称生产
func formatParaName(name string) string {
	return fmt.Sprintf("JYPara%s", name)
}
