//
//  JYMapiNetworkClient+ride_getBikesByLoction.h
//  JYCastKit
//
//  Created by Flywithbug on 2019-06-18.
//  Copyright © 2019 hellobike. All rights reserved.
//
	
#import "JYMapiNetworkClient+ride_getBikesByLoction.h"
	

	
@implementation JYMapiNetworkClient (ride_getBikesByLoction)
- (NSString *)action{
	return @"user.ev.ride.getBikesByLoction";
}

- (nullable NSURLSessionDataTask *)mapi_ride_getBikesByLoction_withmoduleType:(JYModuleType)type
		para:(JYParaQueryPara * _Nullable)param
		progress:(nullable void (^)(NSProgress *_Nonnull  uploadProgress))uploadProgress
		success:(nullable void (^)(NSURLSessionDataTask *_Nullable task, JYRespGetBikesByLocationBikes * _Nullable responseObject))success
		failure:(nullable void (^)(NSURLSessionDataTask * _Nullable task,   JYNetError *_Nullable error))failure{
	param.action = [self action];
	return [self mapi_post:[self moduleApiServer:type] parameters:param repClass:[JYRespGetBikesByLocationBikes class] progress:uploadProgress success:success failure:failure];
}

@end
	
	
	
	
	
