package config

import (
	"os"
	"sync/atomic"
)

type Config struct {
	TesAPIPort       string
	DevMode       bool
	MachineName   string
	Version       string
	DisableLog    bool
	AppURL        string
	MysqlURL      string
	MysqlUser     string
	MysqlPassword string
	MysqlSchema   string
}

var cnf Config
//var scnf sync.Once
var running uint64
//var mutex sync.Mutex

func Install() {
	atomic.StoreUint64(&running, 0)
	hostName := getHostName()

	cnf = Config{
		MachineName:   hostName,
		TesAPIPort:       ":" + os.Getenv("API_PORT"),
		AppURL:        os.Getenv("APP_URL"),
		MysqlURL:      os.Getenv("DB_ADDR"),
		MysqlUser:     os.Getenv("DB_USER"),
		MysqlPassword: os.Getenv("DB_PASSWORD"),
		MysqlSchema:   os.Getenv("DB_SCHEMA"),
		DisableLog:    false,
	}
}

func DBDSN() string {
	return cnf.MysqlUser + ":" + cnf.MysqlPassword + "@tcp(" + cnf.MysqlURL + ":3306)/"+cnf.MysqlSchema+"?parseTime=true"
}

//Get retorna o objeto de configurações da aplicação
func Get() Config {
	return cnf
}

//IsRunning verifica se a aplicação tem que aceitar requisições
func IsRunning() bool {
	return atomic.LoadUint64(&running) > 0
}

//IsNotProduction returns true if application is running in DevMode or MockMode
func IsNotProduction() bool {
	return cnf.DevMode
}

//Stop faz a aplicação parar de receber requisições
func Stop() {
	atomic.StoreUint64(&running, 1)
}

func getHostName() string {
	machineName, err := os.Hostname()
	if err != nil {
		return ""
	}
	return machineName
}
