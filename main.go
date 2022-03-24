package main

import (
	"flag"
	"log"

	"github.com/ArtefactGitHub/Go_T_Clean/external/common"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console"
	"github.com/ArtefactGitHub/Go_T_Clean/external/web"
)

func main() {
	app := NewApp()
	log.Fatal(app.Run())
}

func NewApp() common.App {
	dep := flag.String("dep", string(common.Develop), "デプロイ環境")
	appType := flag.String("type", string(common.Web), "アプリ種別")
	store := flag.String("store", string(common.Memory), "永続化種別")

	flag.Parse()
	log.Printf("\n >>dep: %s\n >>type: %s\n >>store: %s", *dep, *appType, *store)

	if *appType == string(common.Web) {
		return web.NewWebApp(common.DeployType(*dep), common.StoreType(*store))
	} else {
		return console.NewConsoleApp(common.DeployType(*dep), common.StoreType(*store))
	}
}
