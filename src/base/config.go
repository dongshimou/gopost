package base

import (
	"encoding/json"
	"gopost/src/logger"
	"io/ioutil"
)

type Configs struct {
	Debug    bool     `json:"debug"`
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
	Pass        string `json:"pass"`
	Email       string `json:"email"`
	Port        string `json:"port"`
	TLS         bool   `json:"tls"`
	CertFile    string `json:"cert_file"`
	KeyFile     string `json:"key_file"`
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
	logger.Print("=== read product config ===")
	file, err := ioutil.ReadFile("./product.json")
	if err != nil {
		logger.Debug("can't find config : product.json")
	} else {
		return file, err
	}
	logger.Print("=== read develop config ===")
	return ioutil.ReadFile("./develop.json")
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
