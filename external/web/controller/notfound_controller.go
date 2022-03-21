package controller

import (
	"net/http"
	"text/template"

	"github.com/ArtefactGitHub/Go_T_Clean/external/web/config"
)

const commonViewFilePath string = "common/"

func Notfound(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(config.LayoutFile, config.ToPath(commonViewFilePath, "notfound"))
	t.ExecuteTemplate(w, config.LayoutName, nil)
}
