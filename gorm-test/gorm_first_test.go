package gormtest

import (
	"golang-gorm/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	// setiap dns dari masing" database berbeda
	dsn := "root:@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// kita gunakan gorm.Open untuk melakukan koneksi ke database dan kita masukan driver dan gorm konfignya

	// kita berikan logger ke dalam configny supaya dia mencatat semua aktivitas yang sudah di lakukan
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// connection pool
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db
}

var DB = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, DB)
}

// execute data to database
func TestExecute(t *testing.T) {
	err := DB.Exec("insert into sample(id, name) values (?,?)", 1, "farid").Error
	assert.Nil(t, err)
	err = DB.Exec("insert into sample(id, name) values (?,?)", 2, "anang").Error
	assert.Nil(t, err)
	err = DB.Exec("insert into sample(id, name) values (?,?)", 3, "samudra").Error
	assert.Nil(t, err)
	err = DB.Exec("insert into sample(id, name) values (?,?)", 4, "tampan").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   int
	Name string
}

// RowSql from database
func TestRowSqlSample(t *testing.T) {
	var sample Sample
	err := DB.Raw("select * from sample where id = ?", 1).Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "farid", sample.Name)

	var samples []Sample
	err = DB.Raw("select * from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestRowSqlUsers(t *testing.T) {
	var user []model.User
	err := DB.Raw("select * from users").Scan(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, 99, len(user))
}
