package webserver

import (
	"backbone/tools"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
)

var (
	state       string //state given to reddit
	returnstate string //state returned from reddit
)

//the HTTP handler for authentication sending
func AuthSend(w http.ResponseWriter, r *http.Request) {
	authcon, _, err := authLink()
	if err != nil {
		tools.Log.Debug("Error Creating Auth link")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authcon, http.StatusOK)

}

func authLink() (string, string, error) {

	//using hash combined with random gen to create an even more random string
	h := sha256.New()
	h.Write([]byte(string(rand.Int())))
	state = hex.EncodeToString(h.Sum(nil)) //state for ensuring returned tokens are valid

	//link to the authorisation page
	//the compact version is not being used because it's super ugly
	var authcon = "https://www.reddit.com/api/v1/authorize?" +
			"client_id=" + tools.Conf.GetString("redditclient") +
			"&response_type=" + "code" +
			"&state=" + state +
			"&redirect_uri=" + tools.Conf.GetString("redirectURI") +
			"&duration=" + "permanent" +
			"&scope=" + "identity"


	return authcon, state, nil
}

//parses the Query paramaters returned from reddit
func ReturnParse(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	//running through possible errors
	if query["error"] != nil {
		for i := 0; i < len(query["error"]); i++ {
			if query["error"][i] == "access_denied" {
				tools.Log.Debug("User Denied Access")
				//w.WriteHeader(http.StatusForbidden)
				return
			} else if query["error"][i] == "unsupported_response_type" {
				tools.Log.Debug("Unsupported Response Type")
				//w.WriteHeader(http.StatusInternalServerError)
			} else if query["error"][i] == "invalid_scope" {
				tools.Log.Debug("Invalid Scope Parameter")
				//w.WriteHeader(http.StatusInternalServerError)
			} else if query["error"][i] == "invalid_request" {
				tools.Log.Debug("Issue with Request sent")
				//w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}

	if query["state"] != nil {
		if !stateCheck(query["state"][0], state) {
			tools.Log.Debug("Bad State Return")
			//w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	} else {
		tools.Log.Debug("No State Return")
		//w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	resp, err :=retriveToken(query["code"][0])
	if err != nil {
		tools.Log.WithField("Error", err).Debug("Error Retrieving Token")
		return
	}

	fmt.Println(resp.Body)

}

//ensuring the return state is correct
func stateCheck(returnstate string, state string) bool {
	return returnstate == state
}

//TODO: Fix state issue in access token
//retriveing access token

func retriveToken(code string) (*http.Response, error) {
	var (
		retrieveURL = "https://www.reddit.com/api/v1/access_token"
		data = "grant_type=authorization_code&code=" + code +
			"&redirect_uri=" + tools.Conf.GetString("redirectURI")
	)

	client := &http.Client{}

	req, err := http.NewRequest("POST", retrieveURL, bytes.NewBuffer([]byte(data)))
	if err != nil {return nil, err}

	//Setting up proper data
	req.SetBasicAuth(tools.Conf.GetString("redditclient"), tools.Conf.GetString("redditsecret"))
	req.Header.Set("User-Agent", tools.Conf.GetString("redditagent"))

	resp, err := client.Do(req)
	if err != nil {return nil, err}

	return resp, nil
}
