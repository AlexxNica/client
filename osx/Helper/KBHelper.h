//
//  KBHelper.h
//  Keybase
//
//  Created by Gabriel on 4/20/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface KBHelper : NSObject

- (void)listen:(xpc_connection_t)service;

@end
