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
		err := temp.Run()
		if err != nil {
			Log.WithField("Error", err).Warn("Error Running Command")
			return
		}
	}

	clear["linux"] = func() {
		temp := exec.Command("clear")
		temp.Stdout = os.Stdout
		err := temp.Run()
		if err != nil {
			Log.WithField("Error", err).Warn("Error Running Command")
			return
		}
	}
}

func clearScreen() {
	temp, ok := clear[runtime.GOOS]
	if ok {
		temp()
	} else {
		Log.Debug("Current Platform Not Supported! Contact thehowlinggreywolf if this is an error.")
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
