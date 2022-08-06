package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teodoroau/go-gorn-restapi/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	fmt.Println("Servidor iniciado")
	http.ListenAndServe("50.116.46.121:8080", r)

}
