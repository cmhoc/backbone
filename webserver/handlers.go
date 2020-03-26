package webserver

import (
	"backbone/sql-interface"
	"backbone/tools"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

//middleware for logging requests
func Logging(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.RawQuery != "" {
			ReturnParse(w, r)
		}

		handler.ServeHTTP(w, r)

		tools.Log.WithFields(logrus.Fields{
			"Method": r.Method,
			"URL":    r.URL,
			"IP":     r.RemoteAddr,
			"Host":   r.Host,
		}).Trace("Webserver Request")
	})
}

func Billsjson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //setting the content type
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")    //Allowing frame to be loaded from the origin
	w.Header().Set("Connection", "Keep-Alive")         //setting the connection type
	//setting keep alive settings
	w.Header().Set("Keep-Alive", "timeout=5")
	w.Header().Add("Keep-Alive", "max=100")

	//Convert the bills to json
	billsjson, err := json.Marshal(database.Bills)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		tools.Log.Fatal("Error creating bills as json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//write the bills list
	_, err = w.Write(billsjson)
	if err != nil {
		tools.Log.WithField("Error", err).Warn("Error Writing to JSON")
		return
	}
}

func Votejson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting the content type
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")    //Allowing frame to be loaded from the origin
	w.Header().Set("Connection", "Keep-Alive")         //setting the connection type
	//setting keep alive settings
	w.Header().Set("Keep-Alive", "timeout=5")
	w.Header().Add("Keep-Alive", "max=100")

	//Convert the votes to json
	votesjson, err := json.Marshal(database.Votes)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		tools.Log.Fatal("Error creating parties as json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//write the vote list
	_, err = w.Write(votesjson)
	if err != nil {
		tools.Log.WithField("Error", err).Warn("Error Writing to JSON")
		return
	}
}

func Partyjson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //setting the content type
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")    //Allowing frame to be loaded from the origin
	w.Header().Set("Connection", "Keep-Alive")         //setting the connection type
	//setting keep alive settings
	w.Header().Set("Keep-Alive", "timeout=5")
	w.Header().Add("Keep-Alive", "max=100")

	//Convert the parties to json
	partyjson, err := json.Marshal(database.Parties)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		tools.Log.Fatal("Error creating parties as json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//write the bills list
	_, err = w.Write(partyjson)
	if err != nil {
		tools.Log.WithField("Error", err).Warn("Error Writing to JSON")
		return
	}
}