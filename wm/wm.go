package wm

import (
	"fmt"

	"salientwm/logger"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/BurntSushi/xgbutil/xevent"
)

var (
	X 		*xgbutil.XUtil
	Root	*xwindow.Window
)

func initializeRootEventHandlers() {
	var err error

	evMasks := xproto.EventMaskPropertyChange |
		xproto.EventMaskFocusChange |
		xproto.EventMaskButtonPress |
		xproto.EventMaskButtonRelease |
		xproto.EventMaskStructureNotify |
		xproto.EventMaskSubstructureNotify |
		xproto.EventMaskSubstructureRedirect
	err = xwindow.New(X, X.RootWin()).Listen(evMasks)
	
	if err != nil {
		logger.Error.Fatalf("Could not listen to Root window events: %s", err)
	}

	// Oblige map request events
	xevent.MapRequestFun(
		func(X *xgbutil.XUtil, ev xevent.MapRequestEvent) {
			//xclient.New(ev.Window)
			fmt.Println("new map event!")
	}).Connect(X, Root.Id)
	
}

func Initialize(x *xgbutil.XUtil) {
	var err error

	X = x
	Root = xwindow.New(X, X.RootWin())

	if _, err = Root.Geometry(); err != nil {
		logger.Error.Fatalf("Could not get ROOT window geometry: %s", err)
	} else {
		logger.Message.Println("Got root geometry")
	}

	initializeRootEventHandlers()
}