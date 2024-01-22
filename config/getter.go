package config

import "github.com/spf13/viper"

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Address:  viper.GetString("postgres-add"),
		DbName:   viper.GetString("db-name"),
		User:     viper.GetString("db-user"),
		Password: viper.GetString("password"),
		DbArgs:   viper.GetString("db-args"),
	}

}
