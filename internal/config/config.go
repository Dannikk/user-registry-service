package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type RedisCfg struct {
	Addres   string
	Password string
	DB       int
}

type PostgresCfg struct {
	POSTGRES_USER string
	POSTGRES_PW   string
	POSTGRES_DB   string
	PGHOST        string
}

type PGAdminCfg struct {
	PGADMIN_MAIL string
	PGADMIN_PW   string
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

	redis_host, ok := os.LookupEnv("REDISHOST")
	if !ok {
		redis_host = "localhost"
	}

	viper.SetDefault("Redis_Addres", fmt.Sprintf("%s:6379", redis_host))
	viper.SetDefault("Redis_PW", "") // no password set
	viper.SetDefault("Redis_DB", 0)  // use default DB

	c.Redis = RedisCfg{
		Addres:   viper.GetString("Redis_Addres"),
		Password: viper.GetString("Redis_PW"),
		DB:       viper.GetInt("Redis_DB"),
	}

	pghost, ok := os.LookupEnv("PGHOST")
	if !ok {
		pghost = "localhost"
	}
	c.Pgsql = PostgresCfg{
		POSTGRES_USER: viper.GetString("POSTGRES_USER"),
		POSTGRES_PW:   viper.GetString("POSTGRES_PW"),
		POSTGRES_DB:   viper.GetString("POSTGRES_DB"),
		PGHOST:        pghost,
	}

	c.PGAdmin = PGAdminCfg{
		PGADMIN_PW:   viper.GetString("PGADMIN_PW"),
		PGADMIN_MAIL: viper.GetString("PGADMIN_MAIL"),
	}

	return c, nil
}
