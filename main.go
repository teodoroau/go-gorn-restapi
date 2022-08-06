package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	
	http.ListenAndServe("50.116.46.121:8080", r)
}
