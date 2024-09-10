package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	client := SqliteClient{}

	err := client.ConnectToDB()

	assert.Nil(t, err)
	assert.NotNil(t, client.Db)
}

func TestTableCreation(t *testing.T) {
	client := SqliteClient{}
	
	err := client.ConnectToDB()
	assert.Nil(t, err)

	err = client.SetUpDB()
	assert.Nil(t, err)
}
