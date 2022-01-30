package config

import (
	"os"
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

type Config struct{}

func (Config) ConfigParams() *DBConfig {
	dbConfig := &DBConfig{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_USER_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
	
	return dbConfig
}
