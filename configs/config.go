package configs

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Server Server
	Store  Store
}

type Server struct {
	Address string `json:"address"`
}

type Store struct {
	DB  PSQL
	RDB Redis
}

type PSQL struct {
	DSN    string `json:"dsn"`
	Driver string `json:"driver"`
}

type Redis struct {
	Address string `json:"address"`
}

func New() (Configs, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return Configs{}, err
	}

	cfg := Configs{
		Server: Server{
			Address: viper.GetString("server.address"),
		},
		Store: Store{
			DB: PSQL{
				DSN:    viper.GetString("psql.dsn"),
				Driver: viper.GetString("psql.driver"),
			},
			RDB: Redis{
				Address: viper.GetString("redis.address"),
			},
		},
	}

	return cfg, nil
}
