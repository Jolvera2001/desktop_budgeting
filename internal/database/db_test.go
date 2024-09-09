package database

import(
	"testing"
)

func TestConnection(t *testing.T) {
	client := SqliteClient{}
	err := client.ConnectToDB()

	if err != nil {
		t.Fatalf("Expected no error, go %v", err)
	}

	if client.db == nil {
		t.Fatalf("Expected a valid DB connection, go nil")
	}
}