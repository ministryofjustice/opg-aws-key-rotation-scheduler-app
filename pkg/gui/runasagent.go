package gui //#nosec

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

int
SetActivationPolicy(void) {
    [NSApp setActivationPolicy:NSApplicationActivationPolicyAccessory];
    return 0;
}
*/
import "C"
import "fmt"

func setActivationPolicy() {
	fmt.Println("Setting ActivationPolicy")
	C.SetActivationPolicy()
}
