package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const file string = "dev.sqlite"

type SqliteClient struct {
	Db *sql.DB
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

	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	c.Db = db
	return nil
}

func (c *SqliteClient) SetUpDB() error {
	if c.Db == nil {
		return fmt.Errorf("db connection not set up")
	}

	sqlScript := userTable + budgetTable + transactionsTable

	_, err := c.Db.Exec(sqlScript)
	if err != nil {
		return err
	}

	return nil
}
