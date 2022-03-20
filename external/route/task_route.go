package route

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/external/controller"
	"github.com/ArtefactGitHub/Go_T_Clean/external/model"
)

type TaskRoute struct {
	controller *controller.TaskController
}

func NewTaskRoute() TaskRoute {
	c := controller.NewTaskController()
	return TaskRoute{controller: &c}
}

func (r TaskRoute) GetRoutes() []model.Route {
	return []model.Route{
		{Path: "/", Method: http.MethodGet, Handler: r.controller.Index},
		{Path: "/task", Method: http.MethodGet, Handler: r.controller.Index},
		{Path: "/task/new", Method: http.MethodGet, Handler: r.controller.New},
		{Path: "/task", Method: http.MethodPost, Handler: r.controller.Create},
		{Path: "/task/{id}", Method: http.MethodGet, Handler: r.controller.Show},
		{Path: "/task/edit/{id}", Method: http.MethodGet, Handler: r.controller.Edit},
		{Path: "/task/{id}", Method: http.MethodPut, Handler: r.controller.Update},
		{Path: "/task/{id}", Method: http.MethodDelete, Handler: r.controller.Delete},
	}
}
