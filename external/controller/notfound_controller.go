package controller

import (
	"net/http"
	"text/template"
)

const commonViewFilePath string = "common/"

func Notfound(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(layoutFile, toPath(commonViewFilePath, "notfound"))
	t.ExecuteTemplate(w, layoutName, nil)
}
