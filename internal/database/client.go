package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

const file string = "dev.sqlite"

type SqliteClient struct {
	db *sql.DB
}

func (c *SqliteClient) ConnectToDB() error {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return fmt.Errorf("unable to find appdata path")
	}

	appFolderPath := filepath.Join(appDataPath, "Desktop_Budgeting")
	err := os.MkdirAll(appFolderPath, os.ModePerm)
	if err != nil {
		return err
	}

	dbFilePath := filepath.Join(appFolderPath, file)

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}

	c.db = db
	return nil
}
