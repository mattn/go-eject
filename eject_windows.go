package eject

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

func Eject() {
	ole.CoInitialize(0)

	unk, _ := oleutil.CreateObject("WMPlayer.OCX")
	wmp, _ := unk.QueryInterface(ole.IID_IDispatch)
	drives := oleutil.MustGetProperty(wmp, "cdromCollection").ToIDispatch()
	for i := 0; i < int(oleutil.MustGetProperty(drives, "Count").Val); i++ {
		drive := oleutil.MustCallMethod(drives, "Item", i).ToIDispatch()
		oleutil.MustCallMethod(drive, "Eject")
	}
}
