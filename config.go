package main

import (
	"fmt"
	"os"
	"strconv"
)

type DataBaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Dbname   string
	User     string
	Password string
}

type Config struct {
	DBConf DataBaseConfig
}

// NewConfig returns a new Config struct
func NewConfig() Config {
	return Config{
		DBConf: DataBaseConfig{
			Driver:   getEnv("driver", "pgx"),
			Host:     getEnv("host", "localhost"),
			Port:     getEnvAsInt("port", 5432),
			Dbname:   getEnv("dbname", ""),
			User:     getEnv("user", "postgres"),
			Password: getEnv("password", "changeme"),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// DSN return DataSourceName as string
func (dbConfig Config) FormDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		dbConfig.DBConf.Host,
		dbConfig.DBConf.Port,
		dbConfig.DBConf.Dbname,
		dbConfig.DBConf.User,
		dbConfig.DBConf.Password,
	)
}
