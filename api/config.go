package api

import "github.com/spf13/viper"

type config struct {
	Api      apiConfig
	Database databaseConfig
	Redis    redisConfig
}

type apiConfig struct {
	Host    string
	Port    int
	Prefork bool
}

type databaseConfig struct {
	DatabaseName string
	PrepareStmt  bool
}

type redisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var Config config

func InitConfigs() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("Prefork", false)
	viper.SetDefault("PrepareStmt", false)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Config)
}
