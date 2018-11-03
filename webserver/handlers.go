package webserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})

var BillHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	payload, _ := json.Marshal(Votestest)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddVoteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var vote vote
	vars := mux.Vars(r)
	bill := vars["bill"]


	for _, p := range Votestest {
		if p.BillId == bill {
			vote = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if vote.BillId != "" {
		payload, _ := json.Marshal(vote)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})