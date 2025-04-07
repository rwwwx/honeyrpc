package config

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfig struct {
	NetworkType      string `json:"network_type"`
	Port             string `json:"port"`
	Host             string `json:"host"`
	LogLevel         string `json:"log_level"`
	SolanaRpcNodeUrl string `json:"solana_rpc_node_url"`
}

func LoadConfig(configPath string) *AppConfig {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var config AppConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
