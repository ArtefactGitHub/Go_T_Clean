package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/external/controller"
	"github.com/ArtefactGitHub/Go_T_Clean/external/model"
	"github.com/ArtefactGitHub/Go_T_Clean/external/route"
	"github.com/gorilla/mux"
)

func main() {
	url := "localhost"
	port := "8080"

	router := mux.NewRouter()
	routes := getRoutes()
	for _, r := range routes {
		switch r.Method {
		case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
			router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
		default:
		}
	}
	router.NotFoundHandler = http.HandlerFunc(controller.Notfound)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.Notfound)

	address := fmt.Sprintf("%s:%s", url, port)
	log.Printf("running on %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func getRoutes() []model.Route {
	return route.NewTaskRoute().GetRoutes()
}
