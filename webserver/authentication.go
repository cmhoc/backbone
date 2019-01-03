package webserver

import (
	"backbone/tools"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
)

var (
	state string //state given to reddit
	returnstate string //state returned from reddit
	code string //onetime use code reddit gives us to exchange for a bearer token
)

//the HTTP handler for authentication sending
func AuthSend(w http.ResponseWriter, r *http.Request) {
	authcon, err := authLink()
	if err != nil {
		tools.Log.Debug("Error Creating Auth link")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authcon, http.StatusOK)

}

func authLink() (string, error) {

	//using hash combined with random gen to create an even more random string
	h := sha256.New()
	h.Write([]byte(string(rand.Int())))

	var (
		state = hex.EncodeToString(h.Sum(nil)) //state for ensuring returned tokens are valid
		//link to the authorisation page
		authcon = "https://www.reddit.com/api/v1/authorize.compact?" + //the compact version is being used to make it more mobile friendly
		tools.Conf.GetString("redditclient") + "=CLIENT_ID&" +
		"response_type" + "=TYPE&" +
		state + "=RANDOM_STRING&" +
		tools.Conf.GetString("redirectURI") + "=URI&" +
		"permanent" + "=DURATION&" +
		"identity" + "=SCOPE_STRING"
	)

	return authcon, nil
}

//parses the Query paramaters returned from reddit
func ReturnParse(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	//running through possible errors
	for i := 0; i < len(query["error"]); i++ {
		if query["error"][i] == "access_denied" {
			tools.Log.Debug("User Denied Access")
			w.WriteHeader(http.StatusForbidden)
			return
		} else if query["error"][i] == "unsupported_response_type" {
			tools.Log.Debug("Unsupported Response Type")
			w.WriteHeader(http.StatusInternalServerError)
		} else if query["error"][i] == "invalid_scope" {
			tools.Log.Debug("Invalid Scope Parameter")
			w.WriteHeader(http.StatusInternalServerError)
		} else if query["error"][i] == "invalid_request" {
			tools.Log.Debug("Issue with Request sent")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if !stateCheck(query["state"][0], state) {
		tools.Log.Debug("Bad State Return")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//redirecting to the home page if parsed correctly
	http.Redirect(w, r, "cmhoc.ca", http.StatusOK)
}


//ensuring the return state is correct
func stateCheck(returnstate string, state string) (bool) {
	return returnstate == state
}

//retriveing access token
func retriveToken() (error) {
	var data = "grant_type" + "=authorization_code&" +
		code + "=CODE&" +
		tools.Conf.GetString("redirectURI") + "=URI"

	return nil
}