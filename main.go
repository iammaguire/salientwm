package main

import (
	"fmt"
	"salientwm/wm"
	"salientwm/logger"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xprop"
	"github.com/BurntSushi/xgbutil/xevent"
)

type SError struct {
	Msg string
	Component string
}

func (e SError) Error() string {
	return e.Msg
}


// skeleton for implementing ownership function
func own(X *xgbutil.XUtil) error {
	atm, err := xprop.Atm(X, fmt.Sprintf("WM_S%d", X.Conn().DefaultScreen))
	
	if err != nil {
		return err
	}

	reply, err := xproto.GetSelectionOwner(X.Conn(), atm).Reply()
	
	if err != nil {
		return err
	}

	if reply.Owner != xproto.WindowNone {
		return SError { "Other WM already running, taking over not implemented yet...", "own" }
	} else {
		logger.Message.Println("Took ownership")
	}

	return nil
}

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer X.Conn().Close()

	if err := own(X); err != nil {
		fmt.Println(err)
		return
	}

	wm.Initialize(X)

	pingBefore, pingAfter, pingQuit := xevent.MainPing(X)

EVENTLOOP:
	for {
		select {
		case <-pingBefore:
			<-pingAfter
		case <-pingQuit:
			break EVENTLOOP
		}
	}
}