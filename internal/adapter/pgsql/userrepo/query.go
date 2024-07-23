package userrepo

const (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS userstor (id SERIAL PRIMARY KEY, name VARCHAR, age INT)`
	queryCreateUser  = `INSERT INTO userstor (name, age) VALUES ($1, $2) RETURNING id`
)
