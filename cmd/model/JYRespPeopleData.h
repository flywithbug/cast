#import <JYCastKit/JYCastKit.h>

#import "JYParaPeople.h"


NS_ASSUME_NONNULL_BEGIN

@interface JYRespPeopleData:JYResponseModel
@property (nonatomic, copy)  	NSString *name;
@property (nonatomic, assign)	NSInteger age;
@property (nonatomic, strong)	JYParaPeople *people;
@property (nonatomic, copy)  	NSArray <JYParaPeople *> *list;

@end

NS_ASSUME_NONNULL_END



