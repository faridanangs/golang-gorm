package gormtest

import (
	"fmt"
	"golang-gorm/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuerySingleObjectAndInlineCondition(t *testing.T) {
	user := model.User{}
	// first kita gunakan untuk mengambil urutan yang pertama
	// kita bisa menambahkan kndisi di dalam Querynya contoh
	// err := DB.First(&user).Error
	err := DB.First(&user, "id = ?", "3").Error
	assert.Nil(t, err)
	assert.Equal(t, "3", user.ID)

	user = model.User{}
	// last kita gunakan untuk mengambil urutan yang terakhir
	err = DB.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "4", user.ID)

	user = model.User{}
	// take kita gunakan untuk mengambil data sesuai dengan yang kita mau
	// ini lebih baik dari pada menggunakan first dan last karna dia tidak menggunkan
	// order
	err = DB.Take(&user, "id = ?", "2").Error
	assert.Nil(t, err)
	assert.Equal(t, "2", user.ID)
}

func TestQueryAllObjects(t *testing.T) {
	users := []model.User{}
	// kita use find untuk mengambil semua data dan kita juga bisa hanya memanggil sesuai idnya saja
	err := DB.Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []model.User{}
	err = DB.Find(&users, "id in ?", []string{"1", "2"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestQueryCondition(t *testing.T) {
	users := []model.User{}
	// taruh sebelum method find
	// whwre akan menggunakan oprator and untuk mencari data yang kita mau
	// SELECT * FROM `users` WHERE first_name like '%farid%' AND password='farid'
	err := DB.Where("first_name like ?", "%farid%").Where("password=?", "farid").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestOrOprator(t *testing.T) {
	users := []model.User{}
	// Or menggunakan oprator or untuk mencari data yang kita mau
	// SELECT * FROM `users` WHERE first_name like '%farid%' OR password='farid'
	err := DB.Where("first_name like ?", "%farid%").Or("password=?", "farid").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}
func TestNotOprator(t *testing.T) {
	users := []model.User{}
	// Not di gunakan untuk tidak memasukan data yang kita mau
	// SELECT * FROM `users` WHERE first_name like '%farid%' AND NOT password='farid'
	err := DB.Where("first_name like ?", "%farid%").Not("password=?", "farid").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestSelectOprator(t *testing.T) {
	users := []model.User{}
	// Select di gunakan untuk memilih data yang kita mau ambil
	err := DB.Select("id", "first_name").Not("password = ?", "farid").Find(&users).Error
	assert.Nil(t, err)

	for _, data := range users {
		assert.NotEqual(t, "", data.Name.FirstName)
	}

	assert.Equal(t, 3, len(users))
}

func TestStructCondition(t *testing.T) {
	userStruct := model.User{
		// jika kita mengirim string "" dan 0, maka dia akan di anggap sebagai nilai default
		// unutk mengatasi masalah ini kita bisa menggunakan map
		Name:     model.Name{FirstName: "farid", MiddleName: ""},
		Password: "farid",
	}
	users := []model.User{}
	// where akan mencari data yang kita kirim berupa struct jika cocok dia akan menagmbilnya
	err := DB.Where(userStruct).Find(&users).Error
	assert.Nil(t, err)

	assert.Equal(t, 1, len(users))
}
func TestMapCondition(t *testing.T) {
	userMap := map[string]interface{}{
		"middle_name": "",
		"first_name":  "farid",
		"password":    "farid",
	}
	users := []model.User{}
	// where akan mencari data yang kita kirim berupa map jika cocok dia akan menagmbilnya
	err := DB.Where(userMap).Find(&users).Error
	assert.Nil(t, err)

	assert.Equal(t, 0, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	users := []model.User{}
	// jika kita menggunakan offset dia mulai ngambil datanya mulai dari 2 ke atas
	// kita gunakan limit unutk membatasi sampai mana datanya bisa di ambil
	err := DB.Order("id asc, first_name desc").Limit(4).Offset(2).Find(&users).Error
	assert.Nil(t, err)

	fmt.Print(users)
	assert.Equal(t, 3, len(users))
}

type UserResponse struct {
	Id        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	// kita gunakan model untuk memilih filed yang ingin kita pangggil kemudian kita masukan ke dalam users
	err := DB.Model(&model.User{}).Select("id", "first_name", "last_name").Find(&users).Error

	assert.Nil(t, err)
	fmt.Println(users)
}
