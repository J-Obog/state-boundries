package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStateBorders(w http.ResponseWriter, r *http.Request) {
	state := mux.Vars(r)["state"]

	if v, in := mapper[state]; in { //check if key in map
		res, err:= json.Marshal(v)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		
		w.WriteHeader(200)
		w.Write(res)
		return
	} 

	w.WriteHeader(404)
}

func GetAllStateBorders(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(mapper)

	if err != nil { //check for marshalling error
		w.WriteHeader(500)
		return
	} 

	w.WriteHeader(200)
	w.Write(res)
	return
}