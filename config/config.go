package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
	Timeout  string `json:"timeout"`
	RedisIP  string `json:"redisIP"`
}

var Config config

func init() {
	configData, err := os.Open("../../config/config.json")
	if err != nil {
		log.Panicln(err)
		return
	}
	jsonData, err := io.ReadAll(configData)
	if err != nil {
		log.Panicln(err)
		return
	}

	err = json.Unmarshal(jsonData, &Config)
	if err != nil {
		log.Panicln(err)
		return
	}
	return
}
