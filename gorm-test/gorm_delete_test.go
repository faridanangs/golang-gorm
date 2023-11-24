package gormtest

import (
	"fmt"
	"golang-gorm/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftDelete(t *testing.T) {
	todo := model.Todo{
		UserId:      "1",
		Title:       "todos app 1",
		Description: "hello world this is todos app number one",
	}
	err := DB.Create(&todo).Error
	assert.Nil(t, err)

	err = DB.Delete(&todo).Error
	assert.Nil(t, err)

	var todos []model.Todo
	err = DB.Find(&todos).Error
	assert.Nil(t, err)

}
func TestUnscoped(t *testing.T) {
	var todos []model.Todo

	// kita menggunakan Unscoped untuk memastikan bahwa datanya sudah benar" terhapus
	// di dalam database secara permanen jika kita mempunyai data dngan primatykey atau unix indek maka
	// kita harus benar benar memastukan datanya suoaya tidak terjadi dupkikat dan terjadi error
	err := DB.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
	fmt.Println(todos)

	// jika kita mendelete tanpa menggunakan Unscoped maka datanya akan tersimpan di database belum
	// di apus secara permanen hanya mengeluarkan tanggal di hapusnya saja 2023-11-20 18:36:27
	err = DB.Delete(&todos, "id = ?", 5).Error
	assert.Nil(t, err)
}
