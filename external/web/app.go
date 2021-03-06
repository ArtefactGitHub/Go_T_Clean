package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/external/common"
	inmemory "github.com/ArtefactGitHub/Go_T_Clean/external/infurastructure/persistence/inmemory/task"
	setting "github.com/ArtefactGitHub/Go_T_Clean/external/infurastructure/persistence/rdb/mysql/setting"
	rdb "github.com/ArtefactGitHub/Go_T_Clean/external/infurastructure/persistence/rdb/mysql/task"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/config"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/controller"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/middleware"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/model"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/route"
	"github.com/ArtefactGitHub/Go_T_Clean/usecase/task"
	_ "github.com/go-sql-driver/mysql"
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
	// 設定ファイルの取得
	cfg, err := config.LoadConfig("./external/web/config/config.yml")
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	routes := app.getRoutes(cfg)
	for _, r := range routes {
		switch r.Method {
		case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
			router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
		default:
		}
	}
	router.NotFoundHandler = http.HandlerFunc(controller.Notfound)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.Notfound)

	address := fmt.Sprintf("%s:%s", cfg.Url, cfg.Port)
	log.Printf("running on %s", address)
	return http.ListenAndServe(address, middleware.MethodOverride(router))
}

func (app *webApp) getRoutes(cfg config.MyConfig) []model.Route {
	if app.storeType.IsMySql() {
		repository, err := rdb.NewMySqlTaskRepository(
			setting.NewMySqlSetting(
				cfg.SqlDriver, cfg.User, cfg.Password, cfg.Protocol, cfg.Address, cfg.DataBase,
			))
		if err != nil {
			log.Fatalf("NewMySqlTaskRepository() error: %s", err.Error())
		}

		interactor := task.NewTaskInteractor(repository)
		return route.NewTaskRoute(interactor).GetRoutes()
	} else {
		repository, err := inmemory.NewInMemoryTaskRepository()
		if err != nil {
			log.Fatalf("NewInMemoryTaskRepository() error: %s", err.Error())
		}

		task := task.NewTaskInteractor(repository)
		return route.NewTaskRoute(task).GetRoutes()
	}
}
