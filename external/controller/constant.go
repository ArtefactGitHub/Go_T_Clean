package controller

import "fmt"

const (
	layoutName         string = "layout"
	layoutFile         string = "./external/view/layouts/layout.html"
	viewFileCommonPath string = "./external/view/views/"
)

func toPath(viewPath, viewName string) string {
	return fmt.Sprintf("%s%s%s.html", viewFileCommonPath, viewPath, viewName)
}
