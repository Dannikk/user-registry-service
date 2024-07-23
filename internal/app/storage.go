package app

import (
	"database/sql"
	"fmt"
	"log"
	"user_registry/config"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/lib/pq" // postgres
	"github.com/redis/go-redis/v9"
)

func newRedisClient(cfg config.RedisCfg) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return client
}

func newPostgresqlConnection(cfg config.PostgresCfg) (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.DB,
		cfg.Username,
		cfg.Password,
	)

	// Open the connection
	db, err := sql.Open("postgres", pgConnString)
	if err != nil {
		log.Printf("error opening connection: %v\n", err)
		return nil, err
	}

	var attemptNum int
	// check the connection
	err = backoff.Retry(
		func() error {
			err := db.Ping()
			attemptNum++
			log.Printf("Failed to ping db. Attempt=%v: %v\n", attemptNum, err)

			return err
		},
		backoff.NewExponentialBackOff(),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
