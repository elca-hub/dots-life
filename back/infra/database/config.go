package database

import "os"

type MySQLConfig struct {
	host     string
	database string
	port     string
	user     string
	password string
}

func NewMySQLConfig() *MySQLConfig {
	return &MySQLConfig{
		host:     os.Getenv("MYSQL_HOST"),
		database: os.Getenv("MYSQL_DATABASE"),
		port:     os.Getenv("MYSQL_PORT"),
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
	}
}
