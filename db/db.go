package db

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/reward-rabieth/b2b/config"
	"log"
	"log/slog"
)

// Connect  establishes a connection to the database using the provided configuration
func Connect(logger *slog.Logger, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, cfg.URL())
	if err != nil {
		return nil, err
	}

	logger.Info("connected to postgres on address\n" + cfg.URL())

	runDBMigration("file://db/migration", cfg.URL())
	return connPool, nil
}

func runDBMigration(migrationURL string, Url string) {
	m, err := migrate.New("file://db/migration", Url)
	if err != nil {
		log.Fatalln(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln(err)
	}

	slog.Info("db migrated successfully")
}
