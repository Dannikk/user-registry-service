package config

import (
	"fmt"
	"os"
	"user_registry/pkg/types"

	"github.com/spf13/viper"
)

type RedisCfg struct {
	Address  string
	Password string
	DB       int
}

type PostgresCfg struct {
	Username string
	Password string
	DB       string
	Host     string
}

type PGAdminCfg struct {
	Mail     string
	Password string
}

type Config struct {
	Redis   RedisCfg
	Pgsql   PostgresCfg
	PGAdmin PGAdminCfg
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	redisHost, ok := os.LookupEnv("REDISHOST")
	if !ok {
		redisHost = types.LOCALHOST
	}

	viper.SetDefault("Redis_Addres", fmt.Sprintf("%s:6379", redisHost))
	viper.SetDefault("Redis_PW", "") // no password set
	viper.SetDefault("Redis_DB", 0)  // use default DB

	c.Redis = RedisCfg{
		Address:  viper.GetString("Redis_Addres"),
		Password: viper.GetString("Redis_PW"),
		DB:       viper.GetInt("Redis_DB"),
	}

	pghost, ok := os.LookupEnv("PGHOST")
	if !ok {
		pghost = "localhost"
	}
	c.Pgsql = PostgresCfg{
		Username: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PW"),
		DB:       viper.GetString("POSTGRES_DB"),
		Host:     pghost,
	}

	c.PGAdmin = PGAdminCfg{
		Password: viper.GetString("PGADMIN_PW"),
		Mail:     viper.GetString("PGADMIN_MAIL"),
	}

	return c, nil
}
