package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	db "desktop_budgeting/internal/database"
	"desktop_budgeting/internal/features/users"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// establishing db
	sqliteClient := &db.SqliteClient{}
	err := sqliteClient.ConnectToDB()
	if err != nil {
		panic(err)
	}

	// passing on to other structs
	userService := &users.UserService{Client: sqliteClient.Db}

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
			userService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
