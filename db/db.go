package db

import (
	"context"
	"errors"
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
		logger.Error("cannot connect to db", err)
	}

	logger.Info("connected to postgres on address\n" + cfg.URL())

	runDBMigration(cfg.MigrationURl, cfg.URL())
	return connPool, nil
}

func runDBMigration(migrationURL string, dbSource string) {
	m, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot create new migration instance %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to run migrate up %v", err)
	}

	slog.Info("db migrated successfully")
}
