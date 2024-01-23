package db

import (
	"github.com/reward-rabieth/b2b/config"
	procurerModels "github.com/reward-rabieth/b2b/core/components/Procurer/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

var registeredModels = []interface{}{
	&procurerModels.Requisition{},
}

// Connect  establishes a connection to the database using the provided configuration
func Connect(logger *slog.Logger, cfg config.DatabaseConfig) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(cfg.URL()), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to postgres on address" + cfg.URL())

	if err := db.AutoMigrate(registeredModels...); err != nil {
		return nil, err
	}
	return db, nil
}
