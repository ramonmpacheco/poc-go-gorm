package gormedge

import (
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
