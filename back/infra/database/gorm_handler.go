package database

import (
	"back/domain/repo/db"
	"back/infra/database/gorm/repo"
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
		"%s:%s@tcp(%s:%s)/%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		return nil, err
	}

	return &GormHandler{db: db}, nil
}

func (h *GormHandler) UserRepository() db.InterUserRepository {
	return repo.NewGormUserRepository(h.db)
}
