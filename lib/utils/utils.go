package utils

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

var kb keybd_event.KeyBonding

func init() {
	var err error
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
}

func KeyPress(key int) {
	kb.HasCTRL(true)
	kb.HasSHIFT(true)
	kb.SetKeys(key)

	kb.Press()
	time.Sleep(50 * time.Millisecond)
	kb.Release()

	kb.Clear()
}
