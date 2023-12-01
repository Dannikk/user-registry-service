package app

import (
	"database/sql"
	"fmt"
	"log"
	"user_registry/internal/config"
	"github.com/cenkalti/backoff/v4"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func newRedisClient(cfg config.RedisCfg) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addres,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	return client
}

func newPostgresqlConnection(cfg config.PostgresCfg) (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.PGHOST,
		cfg.POSTGRES_DB,
		cfg.POSTGRES_USER,
		cfg.POSTGRES_PW,
	)

	// Open the connection
	db, err := sql.Open("postgres", pgConnString)

	if err != nil {
		log.Printf("error opening connection: %v\n", err)
		return nil, err
	}

	var attemptnum int
	// check the connection
	err = backoff.Retry(
		func () error {
			err := db.Ping()
			attemptnum += 1
			log.Printf("Failed to ping db. Attempt=%v: %v\n", attemptnum, err)
			return err
		},
		backoff.NewExponentialBackOff(),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB(db *sql.DB) error {
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS userstor (id SERIAL PRIMARY KEY, name VARCHAR, age INT)"); err != nil {
		return fmt.Errorf("creating table is failed! %v", err)
	}

	return nil
}

func CloseConnection(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("db pointer is nil")
	}

	log.Println("Closing connection...")
	if err := db.Close(); err != nil {
		return err
	}
	log.Println("Connection closed!")
	return nil
}