//
//  JYRespQueryPara.h
//  JYCastKit
//
//  Created by Flywithbug on 2019-06-18.
//  Copyright © 2019 hellobike. All rights reserved.
//
	
#import <JYCastKit/JYCastKit.h>

	
NS_ASSUME_NONNULL_BEGIN
	
@interface JYRespQueryPara : JYResponseModel
@property (nonatomic, copy)  	NSString *cityCode;
@property (nonatomic, strong)	NSNumber *lng;
@property (nonatomic, strong)	NSNumber *lat;
@property (nonatomic, copy)  	NSString *adCode;
@property (nonatomic, assign)	BOOL redPacketModel;
@end
	
NS_ASSUME_NONNULL_END	
	
	
	
