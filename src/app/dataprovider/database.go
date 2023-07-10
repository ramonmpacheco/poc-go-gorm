package dataprovider

import (
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/postgres/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=poc-go-gorm-postgres user=postgres password=postgres dbname=poc_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Pastel{}, &entity.Ingredient{})
	return db
}
