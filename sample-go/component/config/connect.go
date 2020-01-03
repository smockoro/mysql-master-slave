package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Register MySQL Driver
	"github.com/jmoiron/sqlx"
)

// ConnectDB : connect to mysql server
func ConnectDB(cfg *Config) (*sqlx.DB, error) {
	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DBUser is none")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DBPassword is none")
	}
	if cfg.DBHost == "" {
		return nil, fmt.Errorf("DBHost is none")
	}
	if cfg.DBSchema == "" {
		return nil, fmt.Errorf("DBSchema is none")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBSchema)
	return sqlx.Open("mysql", dsn)
}

// ConnectMasterDB : connect to mysql server
func ConnectMasterDB(cfg *MasterSlaveConfig) (*sqlx.DB, error) {
	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DBUser is none")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DBPassword is none")
	}
	if cfg.MasterDBHost == "" {
		return nil, fmt.Errorf("MasterDBHost is none")
	}
	if cfg.DBSchema == "" {
		return nil, fmt.Errorf("DBSchema is none")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.MasterDBHost,
		cfg.DBSchema)
	return sqlx.Open("mysql", dsn)
}

// ConnectSlaveDB : connect to mysql server
func ConnectSlaveDB(cfg *MasterSlaveConfig, slaveID int64) (*sqlx.DB, error) {
	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DBUser is none")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DBPassword is none")
	}
	if cfg.SlaveNum <= slaveID {
		return nil, fmt.Errorf("Slave ID is invalid")
	}
	if cfg.SlaveDBHosts[slaveID] == "" {
		return nil, fmt.Errorf("SlaveDBHost is none")
	}
	if cfg.DBSchema == "" {
		return nil, fmt.Errorf("DBSchema is none")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.MasterDBHost,
		cfg.DBSchema)
	return sqlx.Open("mysql", dsn)
}
