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

var Database DbConfig

func GetConfig() error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Database)
}
