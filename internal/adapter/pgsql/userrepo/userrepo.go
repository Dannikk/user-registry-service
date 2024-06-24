package userrepo

import (
	"context"
	"database/sql"
	"log"
	"user_registry/internal/entity"
	"user_registry/internal/service/userregistry"
)

type Repository struct {
	pgsql *sql.DB
}

var _ userregistry.Repository = &Repository{}

func New(pgsqlConnect *sql.DB) *Repository {
	return &Repository{
		pgsql: pgsqlConnect,
	}
}

func (repo *Repository) CreateUser(_ context.Context, user *entity.User) (int64, error) {
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
