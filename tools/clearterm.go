package tools

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //storing functions for each OS (I prefer programming on linux but I deploy to windows)
var prevClearTime time.Time //The last time it was cleared

func init() {
	prevClearTime = time.Now()
	clear = make(map[string]func())
	clear["windows"] = func() {
		temp := exec.Command("cmd", "/c", "cls")
		temp.Stdout = os.Stdout
		temp.Run()
	}

	//TODO: Create the linux clear term commands
}

func clearScreen() {
	temp, ok := clear[runtime.GOOS]
	if ok {
		temp()
	} else {
		Log.Debug("Current Platform Not Supported! Contact the author if this is an error.")
		return
	}
}

func ClearLoop() { //Note: Always run on goroutine
	//infinite loop
	for true {
		if prevClearTime.AddDate(0, 0, Conf.GetInt("time")).Unix() < time.Now().Unix() {
			clearScreen()
			prevClearTime = time.Now()
			Log.WithField("Time", time.Now()).Trace("Cleared Terminal")
		}
	}
}
