package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load("../../.env.example"); err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../infra/database/gorm/query",
		ModelPkgPath: "../../infra/database/gorm/orm",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	g.UseDB(db)

	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if tableName == "schema_migrations" { //Just return an empty string and the table will be ignored.
			return ""
		}
		return tableName
	})

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
