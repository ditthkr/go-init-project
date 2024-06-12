package config

import (
	"fmt"
	"github.com/spf13/viper"
	"project/internal/adapter/storage/postgres"
)

func Initial() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetAppAddr() string {
	return fmt.Sprintf(":%v", viper.GetString("app.port"))
}

func GetPostgresConfig() *postgres.Config {
	return &postgres.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

}
