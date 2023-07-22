package test

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (mdb *MockDB) Begin(opts ...*sql.TxOptions) *gorm.DB {
	args := mdb.Called(opts)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Commit() *gorm.DB {
	args := mdb.Called()
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Create(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Delete(value interface{}, args ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{value}, args...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) Exec(sql string, values ...interface{}) *gorm.DB {
	args := mdb.Called(sql, values)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{out}, where...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{out}, where...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) FirstOrInit(out interface{}, where ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{out}, where...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{out}, where...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) Limit(limit int) *gorm.DB {
	args := mdb.Called(limit)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Model(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Offset(offset int) *gorm.DB {
	args := mdb.Called(offset)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Preload(query string, args ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{query}, args...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) Raw(sql string, values ...interface{}) *gorm.DB {
	args := mdb.Called(sql, values)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Rollback() *gorm.DB {
	args := mdb.Called()
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Save(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Scan(dest interface{}) *gorm.DB {
	args := mdb.Called(dest)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Table(name string, args ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{name}, args...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	args := mdb.Called(fc, opts)
	return args.Error(0)
}

func (mdb *MockDB) Unscoped() *gorm.DB {
	args := mdb.Called()
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Update(column string, value interface{}) *gorm.DB {
	args := mdb.Called(column, value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	mockArgs := append([]interface{}{query}, args...)
	return mdb.Called(mockArgs...).Get(0).(*gorm.DB)
}

func (mdb *MockDB) UpdateWithAssociations(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}
