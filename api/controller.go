package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStateBorders(w http.ResponseWriter, r *http.Request) {
	state := mux.Vars(r)["state"]

	if v, in := mapper[state]; in { //check if key in map
		res, _ := json.Marshal(v)
		w.WriteHeader(200)
		w.Write(res)
	} else {
		w.WriteHeader(404)
	}
}

func GetAllStateBorders(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(mapper)
	
	if err != nil { //check for marshalling error
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
		w.Write(res)
	}

}