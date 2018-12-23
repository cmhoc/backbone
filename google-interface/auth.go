package google

import (
	"backbone/tools"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
)

//for use in other sheets related functions
var sheet *spreadsheet.Service

func GoogleAuth() error {
	data, err := ioutil.ReadFile(tools.Conf.GetString("googletoken"))
	if err != nil {
		return fmt.Errorf("google api credentials not loaded")
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return fmt.Errorf("google Scopes not connected properly")
	}

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	sheet = service

	return nil
}
