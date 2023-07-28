package dataprovider

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/entity"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDatabase interface {
	Create(value interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	UpdateWithAssociations(value interface{}) *gorm.DB
	Delete(value interface{}, args ...interface{}) *gorm.DB
}

type Database struct {
	DB *gorm.DB
}

func (dbi *Database) Create(value interface{}) *gorm.DB {
	return dbi.DB.Create(value)
}

func (dbi *Database) Model(value interface{}) *gorm.DB {
	return dbi.DB.Model(value)
}

func (dbi *Database) UpdateWithAssociations(value interface{}) *gorm.DB {
	return dbi.DB.Session(getFullSaveAssociations()).Updates(value)
}

func getFullSaveAssociations() *gorm.Session {
	return &gorm.Session{FullSaveAssociations: true}
}

func (dbi *Database) Delete(value interface{}, args ...interface{}) *gorm.DB {
	return dbi.DB.Select(clause.Associations).Delete(value)
}

func NewPostgres() *Database {
	fmt.Println("About to connect to database")
	dsn := "host=poc-go-gorm-postgres user=postgres password=postgres dbname=poc_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Pastel{}, &entity.Ingredient{})
	return &Database{
		DB: db,
	}
}

func NewSqlite() *Database {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.Pastel{}, &entity.Ingredient{})
	return &Database{
		DB: db,
	}
}
