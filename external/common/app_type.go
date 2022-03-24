package common

type AppType string

const (
	Web     AppType = "web"
	Console AppType = "console"
)

func (t AppType) IsWeb() bool {
	return t == Web
}

func (t AppType) IsConsole() bool {
	return t == Console
}
