package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nedpals/supabase-go"
	"github.com/reward-rabieth/b2b/config"
	"github.com/reward-rabieth/b2b/db"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	usersession "github.com/reward-rabieth/b2b/session"
	"log/slog"
)

type App struct {
	DB                *pgxpool.Pool
	shutdownCallbacks []func()

	//supabase
	supa *supabase.Client

	UserSessionComponent usersession.Component

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

	sbConfig := config.GetSupabaseConfig()
	client := supabase.CreateClient(sbConfig.BaseURL, sbConfig.ApiKey)

	//component initialization
	app.UserSessionComponent = usersession.NewComponent(client)
	return app, nil

}

func (app *App) Shutdown() {
	for _, callBack := range app.shutdownCallbacks {
		callBack()
	}
}
