#import "JYMapiNetworkClient+{*para*}.h"


@implementation JYMapiNetworkClient ({*para*})

- (nullable NSURLSessionDataTask *)mapi_{*para*}_withModuleType:(JYModuleType)type
		para:({*para*} * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))downloadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, {*para*} * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure{
	param.action = @"{*para*}";
	return [self mapi_get:[self moduleApiServer:type] parameters:param repClass:[{*para*} class] progress:downloadProgress success:success failure:failure];
}

@end





