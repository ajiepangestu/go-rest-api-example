package database

func InitDatabase() {
	CreateUserTable()
}

func CreateUserTable() {
	DB.Query(`CREATE TABLE IF NOT EXIST users (
		id SERIAL PRIMARY KEY,
		first_name text,
		last_name text,
		email text UNIQUE,
		password text
	)`)
}
