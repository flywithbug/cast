#import "JYMapiNetworkClient+{*para*}.h"

@implementation {*para*}
- (nullable NSURLSessionDataTask *)mapi_{*para*}_withModuleType:(JYModuleType)type
		para:({*para*} * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, {*para*} * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure{
	param.action = @"{*para*}";
	return [self mapi_post:[self moduleApiServer:type] parameters:param repClass:[{*para*} class] progress:uploadProgress success:success failure:failure];
}

@end





