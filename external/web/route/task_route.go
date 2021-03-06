package route

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_T_Clean/external/web/controller/task"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/model"
	utask "github.com/ArtefactGitHub/Go_T_Clean/usecase/task"
)

type TaskRoute struct {
	controller *task.TaskController
}

func NewTaskRoute(interactor utask.TaskInteractor) TaskRoute {
	c := task.NewTaskController(interactor)
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
