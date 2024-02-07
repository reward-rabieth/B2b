package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"strings"
	"time"
)

const (
	localConfigName = "dev"
)

func setupAllFlags() {
	pflag.String("config-name", localConfigName, "name of the config file")
}

func ConfigureViperSettings() error {
	setupAllFlags()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		slog.Error("could not bind viper flags")
		return err
	}
	//map environmental variable  to viper config
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	slog.Info("Reading from config: " + viper.GetString("config-name"))
	viper.SetConfigName(viper.GetString("config-name"))
	viper.AddConfigPath(".")
	viper.AddConfigPath(".config")

	return nil
}

func readConfig() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := viper.ReadInConfig(); errors.As(err, &viper.ConfigFileNotFoundError{}) {
		slog.Warn("no config file found at path" + currentDir)
	} else if err != nil {
		return err
	}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			viper.WatchConfig()
			viper.OnConfigChange(func(in fsnotify.Event) {
				slog.Info("config file changed")
			})
		}
	}()

	return nil
}

func printActiveConfigToStdout() {
	if activeConfig := viper.GetViper().ConfigFileUsed(); activeConfig != "" {
		slog.Info("using config file " + activeConfig)

		for _, key := range viper.AllKeys() {
			fmt.Printf("%s: %v\n", key, viper.Get(key))
		}
	} else {
		slog.Info("No config file loaded")
	}
}

func ReadConfiguration() error {
	if err := ConfigureViperSettings(); err != nil {
		return err
	}

	if err := readConfig(); err != nil {
		return err
	}
	printActiveConfigToStdout()

	return nil
}
