package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// PostgresConfig is a struct that holds the information obtained from .env file
type PostgresConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

// Dialect returns string 'postgres'
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// ConnInfo returns the PostgresSQL connection url
func (c PostgresConfig) ConnInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

// GetPostgresConfig gets the required info from .env file
func GetPostgresConfig() PostgresConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Incorrect DB_PORT: %s", err.Error())
	}

	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
