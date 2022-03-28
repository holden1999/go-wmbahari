package config

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"go-wmb/manager"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	Server         struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"pass"`
		Name     string `yaml:"name"`
	}
}

//func (c *Config) readConfigFile(path string, name string) (config Config, err error) {
//	v := viper.New()
//	v.SetConfigName(name)
//	v.SetConfigType("yaml")
//	v.AddConfigPath(path)
//	v.AutomaticEnv()
//	err = v.ReadInConfig()
//	if err != nil {
//		return
//	}
//	err = viper.Unmarshal(&config)
//	return
//}

func NewConfig() *Config {
	dataSourceName := Config{Server}
	//fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	return config
}

//func New(path string, configFileName string) Config {
//	c := Config{}
//	config := c.readConfigFile(path, configFileName)
//	return config
//
//}
