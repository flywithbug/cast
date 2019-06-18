//
//  JYMapiNetworkClient+ride_getBikesByLoction.h
//  JYCastKit
//
//  Created by Flywithbug on 2019-06-18.
//  Copyright © 2019 hellobike. All rights reserved.
//
	
#import <JYCastKit/JYCastKit.h>
#import "JYParaQueryPara.h"
#import "JYRespGetBikesByLocationBikes.h"

	
NS_ASSUME_NONNULL_BEGIN
	
@interface JYMapiNetworkClient (ride_getBikesByLoction)
- (nullable NSURLSessionDataTask *)mapi_ride_getBikesByLoction_withmoduleType:(JYModuleType)type
		para:(JYParaQueryPara * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, JYRespGetBikesByLocationBikes * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure;

@end
	
NS_ASSUME_NONNULL_END	
	
	
	
