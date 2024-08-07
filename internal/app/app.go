package app

import (
	"log"
	"user_registry/config"

	"user_registry/internal/handler/http/api"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg       *config.Config
	container *Container
	logger    *log.Logger
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

	app := &App{
		cfg:       cfg,
		container: NewContainer(redisClient, pgsqlConn),
	}

	return app, nil
}

func (app *App) StartHTTPServer() (err error) {
	gengine := gin.Default()

	handler := api.NewHandler(app.container.GetUseCase())

	handler.AddRoutes(gengine)

	err = gengine.Run(":8080")
	app.logger.Println(err)

	return
}

func (app *App) Shutdown() (err error) {
	return app.container.Shutdown()
}
