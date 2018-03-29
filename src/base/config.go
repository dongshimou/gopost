package base

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configs struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
}
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type Server struct {
	Port string `json:"port"`
}

var (
	gConfig *Configs
)

func init() {
	readConfig()
}
func GetConfig() *Configs {
	return gConfig
}
func readConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Print(err)
		return
	}
	c := Configs{}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Print(err)
	}
	gConfig = &c
}
