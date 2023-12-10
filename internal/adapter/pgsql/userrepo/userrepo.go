package userrepo

import (
	"context"
	"database/sql"
	"log"
	"user_registry/internal/entity"
	"user_registry/internal/service/user_reg"
)

type Repository struct {
	pgsql *sql.DB
}

var _ user_reg.Repository = &Repository{}

func New(pgsql_connect *sql.DB) *Repository {
	return &Repository{
		pgsql: pgsql_connect,
	}
}

func (repo *Repository) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	sqlStatement := `INSERT INTO userstor (name, age) VALUES ($1, $2) RETURNING id`
	var id int64

	err := repo.pgsql.QueryRow(sqlStatement, user.Name, user.Age).Scan(&id)

	if err != nil {
		log.Printf("Unable to execute the query. %v\n", err)
		return 0, err
	}

	log.Printf("Inserted a single record %v\n", id)
	return id, nil
}
