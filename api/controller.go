package api

import (
	"encoding/json"
	"net/http"
)

func GetStateBorders(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(mapper)
	w.WriteHeader(200)
	w.Write(res)
}