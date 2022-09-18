package db

import (
	"fmt"
	"os"
)

type Config struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

func (cnfg *Config) bootCnfg() *Config {
	return &Config{
		Host:         os.Getenv("POSTGRES_HOST"),
		Port:         os.Getenv("POSTGRES_PORT"),
		User:         os.Getenv("POSTGRES_USERNAME"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName: os.Getenv("POSTGRES_NAME"),
	}
}

func Resolve(cnfg *Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		cnfg.Host, 
		cnfg.User,
		cnfg.Password,
		cnfg.DatabaseName,
		cnfg.Port,
	)
}