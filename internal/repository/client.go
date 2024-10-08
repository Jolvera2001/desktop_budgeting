package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const file string = "dev.sqlite"

func ConnectToDB() (*gorm.DB, error) {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return nil, fmt.Errorf("unable to find appdata path")
	}

	appFolderPath := filepath.Join(appDataPath, "Desktop_Budgeting")
	err := os.MkdirAll(appFolderPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	dbFilePath := filepath.Join(appFolderPath, file)

	db, err := gorm.Open(sqlite.Open(dbFilePath))
	if err != nil {
		return nil, err
	}

	return db, nil
}
