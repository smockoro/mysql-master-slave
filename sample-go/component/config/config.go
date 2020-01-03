package config

import (
	"os"
	"strings"
)

// Config : ...
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBSchema   string
}

// SEPALAER : ...
const (
	SEPALAER string = ","
)

// MasterSlaveConfig : ...
type MasterSlaveConfig struct {
	MasterDBHost string
	SlaveDBHosts []string
	SlaveNum     int64
	DBUser       string
	DBPassword   string
	DBSchema     string
}

// NewConfig : ...
func NewConfig() *Config {
	var cfg Config
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBSchema = os.Getenv("DB_SCHEMA")
	return &cfg
}

// NewMasterSlaveConfig : ...
func NewMasterSlaveConfig() *MasterSlaveConfig {
	var cfg MasterSlaveConfig
	cfg.MasterDBHost = os.Getenv("MASTER_DB_HOST")
	cfg.SlaveDBHosts = strings.Split(os.Getenv("SLAVE_DB_HOSTS"), SEPALAER)
	cfg.SlaveNum = int64(len(cfg.SlaveDBHosts))
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBSchema = os.Getenv("DB_SCHEMA")
	return &cfg
}
