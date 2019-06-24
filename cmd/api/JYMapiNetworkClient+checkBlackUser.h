#import <JYCastKit/JYCastKit.h>
#import "JYParaPeople.h"
#import "JYRespPeopleData.h"

NS_ASSUME_NONNULL_BEGIN

@interface JYMapiNetworkClient (checkBlackUser)

/**
 note
 */
- (nullable NSURLSessionDataTask *)mapi_checkBlackUser_withModuleType:(JYModuleType)type
		para:(JYParaPeople * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, JYRespPeopleData * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;

@end

NS_ASSUME_NONNULL_END
