package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/xuender/grephub/app"
	"github.com/xuender/kit/los"
	"github.com/xuender/kit/oss"
)

//go:embed all:frontend/www
var assets embed.FS

const (
	_width  = 1024
	_height = 768
)

func main() {
	app := app.InitApp()

	los.Must0(wails.Run(&options.App{
		Title:       "Grep Hub " + oss.Version,
		Width:       _width,
		Height:      _height,
		Menu:        app.Menu,
		AssetServer: &assetserver.Options{Assets: assets},
		OnStartup:   app.Startup,
		Bind:        app.Bind,
	}))
}