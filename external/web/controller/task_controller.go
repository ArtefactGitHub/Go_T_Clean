package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web/config"
	"github.com/gorilla/mux"
)

const taskViewFilePath string = "task/"

type TaskController struct {
	taskInteractor *interactor.TaskInteractor
}

func NewTaskController() TaskController {
	i := interactor.NewTaskInteractor()
	return TaskController{taskInteractor: &i}
}

func (c TaskController) Index(w http.ResponseWriter, r *http.Request) {
	response, err := c.taskInteractor.GetAll()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	t, _ := template.ParseFiles(config.LayoutFile, config.ToPath(taskViewFilePath, "index"))
	t.ExecuteTemplate(w, config.LayoutName, response)
}

func (c TaskController) New(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(config.LayoutFile, config.ToPath(taskViewFilePath, "new"))
	t.ExecuteTemplate(w, config.LayoutName, nil)
}

func (c TaskController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := model.NewTask(0, r.Form.Get("name"))

	id, err := c.taskInteractor.Create(task)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/task/%d", id), http.StatusMovedPermanently)
}

func (c TaskController) Show(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	task, err := c.taskInteractor.Get(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	t, _ := template.ParseFiles(config.LayoutFile, config.ToPath(taskViewFilePath, "show"))
	t.ExecuteTemplate(w, config.LayoutName, task)
}

func (c TaskController) Edit(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	task, err := c.taskInteractor.Get(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	t, _ := template.ParseFiles(config.LayoutFile, config.ToPath(taskViewFilePath, "edit"))
	t.ExecuteTemplate(w, config.LayoutName, task)
}

func (c TaskController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	r.ParseForm()
	task := model.NewTask(id, r.Form.Get("name"))

	_, err = c.taskInteractor.Update(task)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/task/%d", id), http.StatusMovedPermanently)
}

func (c TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	_, err = c.taskInteractor.Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/task", http.StatusMovedPermanently)
}
