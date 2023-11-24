package gormtest

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormConfig() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}

// ### cache prepared statement
// secara default gorm query tidak menyimpan prepare statementnya di dalam memory
// kita bisa mentimpan prepare statemennya di dalam memory sehingga ketika kita membutuhkannya
// dia tidak di buat ulang melainkan di ambil di dalam memori namun memorikita akan semakin besar
// kita bisa aktifin fitur ini di config dengan name PrepareStmt

//  ### SkipDefaultTransaction
// kita juga bisa mematikan auto transactionnya
// secara default ketika kita melakukan crud semua akan di jalankan dalam transaction,
// walaupu kita tidak menjalankannya

// ### Select()
// kita di darankan menggunakan select jika ingin mengambil datanya karna secara default gorm akan menggunakan
// * untuk mengambil datanya, bayangkan jika jumlah kolomnya 100 biji

//  ### Rows()
// saaat kita melakukan query find() ke slice, secara default gorm akan mengambil seluruh data dan menyimpannya di dalam slice
// baik slice yg kita mau atau tidak semuanya akan di ambil dan ini kadang tidak opataimal jika hasil querynya besar
// kita di saraankan menggunakan rows jika kita ingin melakukan query dgn jumlah yng besar, sehingga kita bisa mengambil data
// yang di butuhkan saja tanpa harus menyimpan semua datanya ke memori
