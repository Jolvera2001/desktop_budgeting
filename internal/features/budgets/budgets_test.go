package budgets

import (
	db "desktop_budgeting/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDB db.SqliteClient

func TestImplementInterface(t *testing.T) {
	var _ IBudgetService = &BudgetService{}
}

func TestCreateBudget(t *testing.T) {
	userId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID: userId,
		Name: "DeleteME!",
		Category: "Entertainment",
		Amount: 10495.25,
	}
	service := BudgetService{Client: testDB.Db}

	id, err := service.CreateBudget(budgetToDelete)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEqual(t, 0, id, "id should not be 0")

	tearDown(id, userId)
}

func TestGetbudget(t *testing.T) {
	userId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID: userId,
		Name: "DeleteME!",
		Category: "Entertainment",
		Amount: 10495.25,
	}
	service := BudgetService{Client: testDB.Db}
	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")

	budget, err := service.GetBudget(id)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEmpty(t, budget, "budget struct should not be empty")

	tearDown(id, userId)
}

func TestGetbudgets(t *testing.T) {
	userId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID: userId,
		Name: "DeleteME!",
		Category: "Entertainment",
		Amount: 10495.25,
	}
	service := BudgetService{Client: testDB.Db}
	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")

	budgetList, err := service.GetBudgets(userId)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEmpty(t, budgetList, "budget list should not be empty")

	tearDown(id, userId)
}

func TestUpdatebudget(t *testing.T) {
	// userId, err := setup()
	// assert.Nil(t, err)
}

func TestDeletebudget(t *testing.T) {
	// userId, err := setup()
	// assert.Nil(t, err)
}


func setup() (int64, error) {
	testDB.ConnectToDB()
	testDB.SetUpDB()

	res, err := testDB.Db.Exec("INSERT INTO users (email, name, budgetPeriod) VALUES (?, ?, ?)",
		"something@test.com", "BudgetuserTest", 1)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func tearDown(budgetId int64, userId int64) {
	if budgetId != 0 {
		testDB.Db.Exec("DELETE FROM budgets WHERE id = ?", budgetId)
	}

	if userId != 0 {
		testDB.Db.Exec("DELETE FROM users WHERE id = ?", userId)
	} 

	testDB.Db.Close()
}
