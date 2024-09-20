package database

const (
	userTable string = `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	email TEXT,
	name TEXT,
	budgetPeriod INTEGER
	);`

	budgetTable string = `
	CREATE TABLE IF NOT EXISTS budgets(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userId INT NOT NULL,
	categoryId INT NOT NULL,
	name TEXT,
	amount REAL,
	FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	FOREIGN KEY (categoryId) REFERENCES categories(id) ON DELETE CASCADE
	);`

	transactionsTable string = `
	CREATE TABLE IF NOT EXISTS transactions(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userId INT NOT NULL,
	budgetId INT NOT NULL,
	description TEXT, 
	amount REAL,
	category TEXT,
	date DATETIME,
	FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	FOREIGN KEY (budgetId) REFERENCES budgets(id) ON DELETE CASCADE
	);`

	incomeTable string = `
	CREATE TABLE IF NOT EXISTS income(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userId INT NOT NULL,
	categoryId INT NOT NULL,
	amount REAL,
	isRegular INT,
	date TEXT,
	FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	FOREIGN KEY (budgetId) REFERENCES budgets(id) ON DELETE CASCADE
	);`
	categoryTable string = `
	CREATE TABLE IF NOT EXISTS categories(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userId INT NOT NULL,
	name TEXT,
	FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	);`
)
