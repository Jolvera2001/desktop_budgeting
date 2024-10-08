package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	m "desktop_budgeting/internal/models"
	"desktop_budgeting/internal/repository"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// db setup
	repo, err := repository.ConnectToDB()
	if err != nil {
		panic("issue establishing local db")
	}
	repo.AutoMigrate(&m.User{}, &m.Budget{}, &m.Category{}, &m.Income{}, &m.Transaction{})

	// creating crud services

	// creating services
	

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Budgeting",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
