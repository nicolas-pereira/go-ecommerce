package config

import (
	"encoding/json"
	"io/ioutil"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

const fname = "config.json"

var database *DbConfig

func init() {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &database)
	if err != nil {
		database = nil
	}
}

func Database() *DbConfig {
	return database
}
