package eject

/*
#include <linux/cdrom.h>
*/
import "C"
import (
	"os"
	"syscall"
)

func Eject() {
	f, _ := os.Open("/dev/cdrom")
	defer f.Close()
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(f.Fd()), C.CDROMEJECT, 0)
}
