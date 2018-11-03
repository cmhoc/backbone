//originally this was going to be in the main file but I wanted to be able to use it in other files too.

package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

//creating a new log format
var Log = logrus.New()

func init() {
	Log.Formatter = new(logrus.TextFormatter)
	Log.Formatter.(*logrus.TextFormatter).DisableColors = false
	Log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	Log.Level = logrus.TraceLevel
	Log.Out = os.Stdout
}