package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/reward-rabieth/b2b/config"
	"github.com/reward-rabieth/b2b/db"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"log/slog"
)

type App struct {
	DB                *pgxpool.Pool
	shutdownCallbacks []func()

	//logger
	logger *slog.Logger

	//db repositories
	repos users.Store
}

func NewApp() (app App, err error) {
	app = App{}
	app.logger = slog.Default()
	if app.DB, err = db.Connect(app.logger, config.GetDatabaseConfig()); err != nil {
		return
	}

	app.shutdownCallbacks = []func(){}

	//repositories initialization
	if app.repos = users.NewStore(app.DB); err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Shutdown() {
	for _, callBack := range app.shutdownCallbacks {
		callBack()
	}
}
