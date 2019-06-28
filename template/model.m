#import "{*para*}.h"

@implementation {*para*}

+ (NSDictionary *)modelContainerPropertyGenericClass {
	// value should be Class or Class name.
	return @{{*para*}
	};
}

- (void)encodeWithCoder:(nonnull NSCoder *)aCoder {
    [self jy_modelEncodeWithCoder:aCoder];
}

- (nullable instancetype)initWithCoder:(nonnull NSCoder *)aDecoder {
    return [self jy_modelInitWithCoder:aDecoder];
}


@end