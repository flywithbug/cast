#import "JYRespPeopleData.h"

@implementation JYRespPeopleData

+ (NSDictionary *)modelContainerPropertyGenericClass {
	// value should be Class or Class name.
	return @{@"list" : [JYParaPeople class],
	};
}

@end