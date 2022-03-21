package config

import "fmt"

const (
	LayoutName         string = "layout"
	LayoutFile         string = "./external/web/view/layouts/layout.html"
	ViewFileCommonPath string = "./external/web/view/views/"
)

func ToPath(viewPath, viewName string) string {
	return fmt.Sprintf("%s%s%s.html", ViewFileCommonPath, viewPath, viewName)
}
