package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/controller"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/middleware"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/model"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/route"
	"github.com/gorilla/mux"
)

type App struct{}

func (app *App) Run() {
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
	log.Fatal(http.ListenAndServe(address, middleware.MethodOverride(router)))
}

func getRoutes() []model.Route {
	interactor := interactor.NewTaskInteractor()
	return route.NewTaskRoute(interactor).GetRoutes()
}
