package config

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"go-wmb/manager"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	path           string
	name           string
}

func (c *Config) readConfigFile() *viper.Viper {
	viper.SetConfigName(c.name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config File Not Found")
		} else {
			panic("Config File Error")
		}
	}
	return viper.GetViper()
}

func (c Config) Get(key string) string {
	return c.readConfigFile().GetString(key)
}

func NewConfig(path, configFileName string) *Config {
	config := new(Config)
	config.name = configFileName
	config.path = path
	dbHost := config.Get("wmbahari.postgres.host")
	dbPort := config.Get("wmbahari.postgres.port")
	dbName := config.Get("wmbahari.postgres.name")
	dbUser := config.Get("wmbahari.postgres.user")
	dbPassword := config.Get("wmbahari.postgres.password")
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager
	return config
}
