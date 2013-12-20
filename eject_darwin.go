package eject
/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework DiscRecording

#import <Foundation/Foundation.h>
#import <DiscRecording/DiscRecording.h>

static int
_eject() {
    @autoreleasepool {
        NSArray* devices = [DRDevice devices];
        if ([devices count] < 1) {
            return -1;
        }

        DRDevice* firstDevice = devices[0];
        NSDictionary* status = [firstDevice status];

        BOOL ok;
        if ([(NSNumber*)status[DRDeviceIsTrayOpenKey] intValue] == 1) {
            ok = [firstDevice closeTray];
        }
        else {
            ok = [firstDevice openTray];
        }

        if (ok) {
            return 0;
        }
        else {
            return -2;
        }
    }
}

*/
import "C"
import "errors"

func Eject() error {
	if r := C._eject(); r != 0 {
		if r == -1 {
			return errors.New("drive not found")
		} else {
			return errors.New("failed to eject")
		}
	}
	return nil
}
