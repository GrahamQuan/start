package main

import "net/http"

func handlerRead(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}