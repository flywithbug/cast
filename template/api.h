#import <JYCastKit/JYCastKit.h>

{*para*}

NS_ASSUME_NONNULL_BEGIN

@interface {*para*}
- (nullable NSURLSessionDataTask *)mapi_{*para*}_withModuleType:(JYModuleType)type
		para:({*para*} * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, {*para*} * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;

@end

NS_ASSUME_NONNULL_END
