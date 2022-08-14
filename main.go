package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teodoroau/go-gorn-restapi/routes"
)

func main() {
	IP := "50.116.46.121"

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	fmt.Println("Servidor iniciado")
	http.ListenAndServe(IP + ":8080", r)
}
