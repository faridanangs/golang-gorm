package gormtest

import (
	"golang-gorm/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	var users model.User
	err := DB.Take(&users, "id = ?", "1").Error
	assert.Nil(t, err)

	users.Name.FirstName = "wagas"
	users.Name.MiddleName = "udin"
	users.Name.LastName = "samud"

	err = DB.Save(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, "wagas", users.Name.FirstName)

}

type UserResponseUpdated struct {
	Id         string
	FirstName  string
	MiddleName string
	LastName   string
}

func TestUpdateSelectColumn(t *testing.T) {
	err := DB.Model(&model.User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"first_name":  "farid anang samudra",
		"middle_name": "",
		"last_name":   "hello world",
	}).Error
	assert.Nil(t, err)

	err = DB.Model(&model.User{}).Where("id = ?", "1").Update("first_name", "saipul").Error
	assert.Nil(t, err)

	err = DB.Model(&model.User{}).Where("id = ?", "1").Updates(model.User{
		Name: model.Name{FirstName: "raika", MiddleName: "aisyah", LastName: "asgoriah"},
	}).Error

	assert.Nil(t, err)

}
