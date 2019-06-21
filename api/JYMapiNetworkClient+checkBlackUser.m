#import "JYMapiNetworkClient+checkBlackUser.h"


@implementation JYMapiNetworkClient (checkBlackUser)

- (nullable NSURLSessionDataTask *)mapi_checkBlackUser_withModuleType:(JYModuleType)type
		para:(JYParaPeople * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, JYRespPeopleData * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure{
	param.action = @"user.ev.checkBlackUser";
	return [self mapi_post:[self moduleApiServer:type] parameters:param repClass:[JYRespPeopleData class] progress:uploadProgress success:success failure:failure];
}

@end





