package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/external/common"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/controller"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/middleware"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/model"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/route"
	"github.com/ArtefactGitHub/Go_T_Clean/infurastructure"
	ifmodel "github.com/ArtefactGitHub/Go_T_Clean/infurastructure/model"
	"github.com/gorilla/mux"
)

type webApp struct {
	deployType common.DeployType
	storeType  common.StoreType
}

func NewWebApp(deployType common.DeployType, storeType common.StoreType) common.App {
	app := webApp{deployType: deployType, storeType: storeType}
	return &app
}

func (app *webApp) Run() error {
	url := "localhost"
	port := "8080"

	router := mux.NewRouter()
	routes := app.getRoutes()
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
	return http.ListenAndServe(address, middleware.MethodOverride(router))
}

func (app *webApp) getRoutes() []model.Route {
	if app.storeType.IsMySql() {
		repository, err := infurastructure.NewMySqlTaskRepository(ifmodel.NewMySqlSetting(
			"", "", "", "", "", "",
		))
		if err != nil {
			log.Fatalf("NewMySqlTaskRepository() error: %s", err.Error())
		}

		interactor := interactor.NewTaskInteractor(repository)
		return route.NewTaskRoute(interactor).GetRoutes()
	} else {
		repository, err := infurastructure.NewInMemoryTaskRepository()
		if err != nil {
			log.Fatalf("NewInMemoryTaskRepository() error: %s", err.Error())
		}

		interactor := interactor.NewTaskInteractor(repository)
		return route.NewTaskRoute(interactor).GetRoutes()
	}
}
