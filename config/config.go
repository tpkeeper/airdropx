package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	EthApi                  string
	ExcelFilePath           string
	SheetName               string
	AirDropXContractAddress string
	TokenContractAddress    string
	StartIndex              uint64
	TransNumberLimit        uint64
	ChainId                 int64
	GasLimit                int64
	Seed                    string
	LogFilePath             string
}

func Load() (*Config, error) {
	configFilePath := flag.String("C", "conf.toml", "Config file path")
	flag.Parse()
	cfg := Config{}
	if err := loadSysConfig(*configFilePath, &cfg); err != nil {
		return nil, err
	}
	if cfg.LogFilePath == "" {
		cfg.LogFilePath = "./log_data"
	}

	return &cfg, nil
}

func loadSysConfig(path string, config *Config) error {
	_, err := os.Open(path)
	if err != nil {
		return err
	}
	if _, err := toml.DecodeFile(path, config); err != nil {
		return err
	}
	fmt.Println("load sysConfig success")
	return nil
}
