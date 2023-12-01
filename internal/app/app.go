package app

import (
	"log"
	"user_registry/internal/handler/http/api"

	"github.com/gin-gonic/gin"
	"user_registry/internal/config"
)

type App struct {
	cfg    *config.Config
	c      *Container
	logger *log.Logger
}

func NewApp(cfgPath string) (*App, error) {
	cfg, err := config.NewConfig(cfgPath)

	if err != nil {
		log.Printf("config reading error: %v\n", err)
		return nil, err
	}

	redisClient := newRedisClient(cfg.Redis)
	pgsqlConn, err := newPostgresqlConnection(cfg.Pgsql)

	if err != nil {
		log.Printf("pgsql connection error: %v\n", err)
		return nil, err
	}

	if err := initDB(pgsqlConn); err != nil {
		log.Printf("initDB error: %v\n", err)
		return nil, err
	}

	app := &App{
		cfg: cfg,
		c:   NewContainer(redisClient, pgsqlConn),
	}

	return app, nil
}

func (app *App) StartHTTPServer() (err error) {
	gengine := gin.Default()

	handler := api.NewHandler(app.c.GetUseCase())

	handler.AddRoutes(gengine)

	err = gengine.Run(":8080")
	app.logger.Println(err)

	return
}
