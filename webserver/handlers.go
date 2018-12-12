package webserver

import (
	"discordbot/logging"
	"discordbot/sql-interface"
	"encoding/json"
	"net/http"
)

var Handlers map[string]func(http.ResponseWriter, *http.Request)

//this init is used to set all the handlers to the map
func init() {
	Handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	Handlers["/sheetdata"] = billsjson
	Handlers["/sheet"] = sheet
}

func billsjson(w http.ResponseWriter, r *http.Request) {
	//Convert the bills to json
	billsjson, err := json.Marshal(database.Bills[0])

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		logger.Log.Fatal("Error create bills as json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//write the bills list
	w.Write(billsjson)
}

func sheet(w http.ResponseWriter, r *http.Request) {
	staticFileDirectory := http.Dir("./webserver/static/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	http.Handle("/", staticFileHandler)
}
