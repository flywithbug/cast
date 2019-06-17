//
//  JYRespGetBikesByLocationBikes.h
//  JYCastKit
//
//  Created by Flywithbug on 2019-06-17.
//  Copyright © 2019 hellobike. All rights reserved.
//
	
#import <JYCastKit/JYCastKit.h>
#import "JYRespGetBikesByLocation.h"

	
NS_ASSUME_NONNULL_BEGIN
	
@interface JYRespGetBikesByLocationBikes : JYResponseModel
@property (nonatomic, copy)  	NSString *bikeNo;
@property (nonatomic, strong)	NSNumber *lng;
@property (nonatomic, strong)	NSNumber *lat;
@property (nonatomic, assign)	BOOL bikeNo3;
@property (nonatomic, strong)	JYRespGetBikesByLocation *QueryPara;
@property (nonatomic, copy)  	NSArray <JYRespGetBikesByLocation *> *list;
@end
	
NS_ASSUME_NONNULL_END	
	
	
	
