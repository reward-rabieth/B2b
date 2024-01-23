package server

import (
	"github.com/reward-rabieth/b2b/config"
	"github.com/reward-rabieth/b2b/core/components/Procurer"
	"github.com/reward-rabieth/b2b/db"
	"gorm.io/gorm"
	"log/slog"
)

type App struct {
	DB                *gorm.DB
	shutdownCallbacks []func()

	//logger
	logger *slog.Logger

	//db repositories
	ProcurerRepo Procurer.Repo

	//component handler
	procurerComponent Procurer.Component
}

func NewApp() (app *App, err error) {
	app = &App{}
	app.logger = slog.Default()
	if app.DB, err = db.Connect(app.logger, config.GetDatabaseConfig()); err != nil {
		return
	}

	app.shutdownCallbacks = []func(){}

	//repositories initialization
	if app.ProcurerRepo, err = Procurer.NewRepo(app.DB); err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Shutdown() {
	for _, callBack := range app.shutdownCallbacks {
		callBack()
	}
}
