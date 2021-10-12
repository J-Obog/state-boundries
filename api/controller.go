package api

import "net/http"

func GetStateBorders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello!</h1>"))
}