package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 500, "Something went wrong ")
}
