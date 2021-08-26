package database

func InitDatabase() error {
	err := CreateUserTable()
	return err
}

func CreateUserTable() error {
	_, err := DB.Query(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name text,
		last_name text,
		email text UNIQUE,
		password text
	)`)
	return err
}
