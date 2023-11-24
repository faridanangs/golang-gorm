package model

import (
	"time"

	"gorm.io/gorm"
)

// defaul convention di gorm
// jika kita membuat struct dengna nama
// user => maka gorm akan membuatkan kita menjadi users #jamak
// OrderDetail => order_details

// kita juga bisa menggunakan tag untuk memberi tanda pada setiap field
// yang kita punya walaupun secara default gorm akan mengubah field kita menjadi small_case
// contoh CreatedAt => created_at oleh gorm
// kita bisa mengcustomnya filednya sesuai dengan column yang ada di db
// contoh CreatedAt => gorm:"column:dibuat_pada"
// tanda ini <-:create maksudnya hanya bisa di create tapi tidak bisa di update
// tanda ini <-:update maksudnya hanya bisa di update tapi tidak bisa di create
// tanda ini <- maksudnya bisa di update dan create
// tanda ini - maksudnya tidak bisa di read/write
type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
	Address     []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProduct []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
	// foreignKey:id ngambil dari user ID
	// joinForeignKey:user_id negambil dari dalam table user_like_product
	// references:id ngambil dari struct []Product
	// joinReferences:product_id ngambil dari dalam table user_like_product
}

// jika kita menggunakan manual nama table, kita bisa menggunakan interface
// Tabler, yang wajib memiliki method dengan nama TableName()
// kode ini berbentuk statis hanya di jalankan di awal saja tidak bisa di ubah"
func (u *User) TableName() string {
	return "users"
}

// urutan eksekusi hoknya untuk create
// begin transaction
// BeforeSave()
// BeforeCreate()
// save before associations
// insert into database
// save after association
// AfterSave()
// AfterCreate()
// commit or rollback transaction

// untuk bisa menjalankan urutan di atas kita harus membuat methodnya cntoh
// Hook ini di jalankan sebelum datanya di create
// sebenarnya bukan hanya create ada juga update,delete,find
func (user *User) BeforeCreate(db *gorm.DB) error {
	if user.ID == "" {
		user.ID = "user-" + time.Now().Format("20060102150405")
	}
	return nil
}

// tag embedded di gunakan untuk menggroup struct supaya tidak terlalu panjang struct mainnya
type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
