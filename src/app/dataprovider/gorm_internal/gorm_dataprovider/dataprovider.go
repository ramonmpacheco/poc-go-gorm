package gormdataprovider

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	gormedge "github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/gorm_internal/gorm_edge"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPostgres() *gormedge.Database {
	fmt.Println("About to connect to database")
	dsn := "host=poc-go-gorm-postgres user=postgres password=postgres dbname=poc_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Pastel{}, &entity.Ingredient{})
	return &gormedge.Database{
		DB: db,
	}
}

func NewSqlite() *gormedge.Database {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Pastel{}, &entity.Ingredient{})
	return &gormedge.Database{
		DB: db,
	}
}
