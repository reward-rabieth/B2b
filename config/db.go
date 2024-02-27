package config

// DatabaseConfig holds configuration data.
type DatabaseConfig struct {
	Address      string
	DbName       string
	User         string
	Password     string
	DbArgs       string
	MigrationURl string
	dbSource     string
}

func (cfg *DatabaseConfig) URL() string {
	//return fmt.Sprintf("postgres://%s:%s@%s?dbname=%s&%s", cfg.User, cfg.Password, cfg.Address, cfg.DbName, cfg.DbArgs)
	return cfg.dbSource
}
