package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerConfig `json:"server"`
	DBСonfig     `json:"db"`
}

type ServerConfig struct {
	Addr         string `json:"addr"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

type DBСonfig struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DbName string `json:"db_name"`
}

func NewConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile("../../config/config.json")
	if err != nil {
		return config, fmt.Errorf("error while trying to read config file: %v", err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("error while trying to unmarshall config json: %v", err)
	}

	return config, nil
}

func (db DBСonfig) ToSourceString() string {
	return fmt.Sprintf("mongodb://%s:%d/", db.Host, db.Port)
}
