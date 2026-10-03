package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type GormHandler struct {
	db *gorm.DB
}

func NewGormHandler(c *MySQLConfig) (*GormHandler, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	engine, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := engine.Ping(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &GormHandler{db: db}, nil
}
