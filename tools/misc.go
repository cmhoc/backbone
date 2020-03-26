//created as a separate file for easy use in other files
package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

//creating a new log format
var Log = logrus.New()
var Debug bool
var Conf = viper.New()

func init() {
	//setting config file path
	Conf.AddConfigPath(".")
	//setting defaults
	Conf.SetDefault("debug", true)
	//Importing Config File
	viper.SetConfigType("toml")
	Conf.SetConfigName("config")
	err := Conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal Error, Config File: %s \n", err))
	}

	//formatting the log
	Log.Formatter = new(logrus.TextFormatter)
	Log.Formatter.(*logrus.TextFormatter).DisableColors = false
	Log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	if Conf.GetBool("debug") {
		Log.Level = logrus.TraceLevel
		Log.Out = os.Stdout
		Log.Info("Debug Logging Mode On")
	} else {
		output, err := os.OpenFile("backbonelog", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Panic("Error Opening File")
		}
		Log.Level = logrus.InfoLevel
		Log.SetOutput(output)
		Log.Info("Debug Logging Mode Off")
	}
}
