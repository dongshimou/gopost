package base

import (
	"encoding/json"
	"io/ioutil"
	"logger"
)

type Configs struct {
	Debug bool `json:"debug"`
	Database Database `json:"database"`
	Server   Server   `json:"server"`
	Token    Token    `json:"token"`
}
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type Server struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Email       string `json:"email"`

	Port string `json:"port"`
}
type Token struct {
	SecretKey string `json:"secret_key"`
	TTL       int64  `json:"ttl"`
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
func read() ([]byte, error) {
	logger.Print("=== read develop config ===")
	file, err := ioutil.ReadFile("./develop.json")
	if err != nil {
		logger.Debug("can't find config : develop.json")
	} else {
		return file, err
	}
	logger.Print("=== read product config ===")
	return ioutil.ReadFile("./product.json")
}
func readConfig() {
	file, err := read()
	if err != nil {
		logger.Print(err)
		logger.Error("=== read config fail===")
		return
	}
	c := Configs{}
	err = json.Unmarshal(file, &c)
	if err != nil {
		logger.Print(err)
	}
	gConfig = &c
	logger.SetDebugStatus(gConfig.Debug)
}
