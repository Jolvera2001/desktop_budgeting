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
	userId, categoryId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID:   userId,
		CategoryID: categoryId,
		Name:     "DeleteME!",
		Amount:   10495.25,
	}
	service := BudgetService{Client: testDB.Db}

	id, err := service.CreateBudget(budgetToDelete)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEqual(t, 0, id, "id should not be 0")

	tearDown(id, userId, categoryId)
}

func TestGetbudget(t *testing.T) {
	userId, categoryId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID:   userId,
		CategoryID: categoryId,
		Name:     "DeleteME!",
		Amount:   10495.25,
	}
	service := BudgetService{Client: testDB.Db}
	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")

	budget, err := service.GetBudget(id)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEmpty(t, budget, "budget struct should not be empty")

	tearDown(id, userId, categoryId)
}

func TestGetbudgets(t *testing.T) {
	userId, categoryId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID:   userId,
		CategoryID: categoryId,
		Name:     "DeleteME!",
		Amount:   10495.25,
	}
	service := BudgetService{Client: testDB.Db}
	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")

	budgetList, err := service.GetBudgets(userId)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEmpty(t, budgetList, "budget list should not be empty")

	tearDown(id, userId, categoryId)
}

func TestUpdatebudget(t *testing.T) {
	userId, categoryId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID:   userId,
		CategoryID: categoryId,
		Name:     "DeleteME!",
		Amount:   10495.25,
	}
	service := BudgetService{Client: testDB.Db}

	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")
	budgetToChange, err := service.GetBudget(id)
	assert.Nil(t, err, "process should not bring up an error")
	update := budgetToChange
	update.Name = "updated"
	err = service.UpdateBudget(update)
	assert.Nil(t, err, "process should not bring up an error")
	updatedBudget, err := service.GetBudget(id)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEqual(t, budgetToChange, updatedBudget)

	tearDown(id, userId, categoryId)

}

func TestDeletebudget(t *testing.T) {
	userId, categoryId, err := setup()
	assert.Nil(t, err, "process should not bring up an error")
	budgetToDelete := BudgetDto{
		UserID:   userId,
		CategoryID: categoryId,
		Name:     "DeleteME!",
		Amount:   10495.25,
	}
	service := BudgetService{Client: testDB.Db}

	id, err := service.CreateBudget(budgetToDelete)
	assert.Nil(t, err, "process should not bring up an error")
	err = service.DeleteBudget(id)
	assert.Nil(t, err, "process should not bring up an error")
	actualBudget, err := service.GetBudget(id)

	assert.NotNil(t, err, "there should be an error")
	assert.Equal(t, Budget{}, actualBudget, "returned budget should be empty")

	tearDown(0, userId, categoryId)
}

func setup() (int64, int64, error) {
	testDB.ConnectToDB()
	testDB.SetUpDB()

	res, err := testDB.Db.Exec("INSERT INTO users (email, name, budgetPeriod) VALUES (?, ?, ?)",
		"something@test.com", "BudgetuserTest", 1)
	if err != nil {
		return 0, 0, err
	}

	userId, err := res.LastInsertId()

	res2, err := testDB.Db.Exec("INSERT INTO categories (userId, name) VALUES (?, ?)",
		userId, "testCategory")

	categoryId, err := res2.LastInsertId()

	return userId, categoryId, err
}

func tearDown(budgetId int64, userId int64, categoryId int64) {
	if budgetId != 0 {
		testDB.Db.Exec("DELETE FROM budgets WHERE id = ?", budgetId)
	}

	if userId != 0 {
		testDB.Db.Exec("DELETE FROM users WHERE id = ?", userId)
	}

	if userId != 0 {
		testDB.Db.Exec("DELETE FROM categories WHERE id = ?", categoryId)
	}

	defer testDB.Db.Close()
}
