package gormtest

import (
	"golang-gorm/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// transaction bisa terjaadi jika use db conecction is same
// connection pool di atur oleh gorm
// if we doing transaction, we can use method Transaction(callback), and in func callback we can using all code transaction exemple create dll.
func TestTransactionSuccess(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model.User{
			ID:       "1",
			Password: "farid",
			Name: model.Name{
				FirstName:  "farid",
				MiddleName: "anang",
				LastName:   "samudra",
			},
			Information: "hello world",
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.User{
			ID:       "2",
			Password: "anang",
			Name: model.Name{
				FirstName:  "farid",
				MiddleName: "anang",
				LastName:   "samudra",
			},
			Information: "hello world",
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.User{
			ID:       "3",
			Password: "samudra",
			Name: model.Name{
				FirstName:  "farid",
				MiddleName: "anang",
				LastName:   "samudra",
			},
			Information: "hello world",
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)

}
func TestTransactionError(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model.User{
			ID:       "4",
			Password: "farid 4",
			Name: model.Name{
				FirstName:  "farid",
				MiddleName: "anang",
				LastName:   "samudra",
			},
			Information: "hello world",
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.User{
			ID:       "2",
			Password: "anang",
			Name: model.Name{
				FirstName:  "farid",
				MiddleName: "anang",
				LastName:   "samudra",
			},
			Information: "hello world",
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := DB.Begin()
	defer tx.Rollback()

	err := tx.Create(&model.User{
		ID:       "4",
		Password: "farid 4",
		Name: model.Name{
			FirstName:  "farid",
			MiddleName: "anang",
			LastName:   "samudra",
		},
		Information: "hello world",
	}).Error
	if err == nil {
		tx.Commit()
	}

	assert.Nil(t, err)

}
func TestManualTransactionError(t *testing.T) {
	tx := DB.Begin()
	defer tx.Rollback()

	err := tx.Create(&model.User{
		ID:       "4",
		Password: "farid 4",
		Name: model.Name{
			FirstName:  "farid",
			MiddleName: "anang",
			LastName:   "samudra",
		},
		Information: "hello world",
	}).Error
	// if err dia aka menjalankan defer rollback
	if err == nil {
		tx.Commit()
	}

	assert.NotNil(t, err)

}

func TestLock(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var user model.User
		// kita menggunakan clauses untuk melakukan loocking supaya tidak terjadi trast pada database
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "2").Error
		if err != nil {
			return err
		}
		user.Name.FirstName = "unyuil"
		user.Name.LastName = "pocong"

		err = tx.Save(&user).Error
		return err
	})

	assert.Nil(t, err)
}
