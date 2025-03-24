package config

import (
	"os"
)

type Postgresql struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewConfigPostgresql() Postgresql {
	return Postgresql{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
