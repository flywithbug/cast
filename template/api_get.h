#import <JYCastKit/JYCastKit.h>
#import "{*para*}.h"
#import "{*para*}.h"

NS_ASSUME_NONNULL_BEGIN

@interface JYMapiNetworkClient ({*para*})

/**
{*para*}
 */
- (nullable NSURLSessionDataTask *)mapi_{*para*}_withModuleType:(JYModuleType)type
		para:({*para*} * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))downloadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, {*para*} * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;

@end

NS_ASSUME_NONNULL_END
