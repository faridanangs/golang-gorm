package gormtest

import (
	"context"
	"fmt"
	"golang-gorm/model"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestCreate(t *testing.T) {
	user := model.User{
		ID:       "5",
		Password: "farid",
		Name: model.Name{
			FirstName:  "udin",
			MiddleName: "anang",
			LastName:   "samudra",
		},
		Information: "hello world",
	}

	response := DB.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)

}

// jika kita ingin memasukan data yang banyak dalam sekali insert kita bisa menggunakan batch insert
func TestBatchInsert(t *testing.T) {
	var users []model.User

	for i := 1; i < 100; i++ {
		users = append(users, model.User{
			ID:       strconv.Itoa(i),
			Password: "kepo lu ke " + strconv.Itoa(i),
			Name: model.Name{
				FirstName: "user ke " + strconv.Itoa(i),
			},
		})
	}

	response := DB.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(99), response.RowsAffected)
}

func TestGormAutoIncrement(t *testing.T) {
	for i := 1; i < 10; i++ {
		userLogs := model.UserLogs{
			UserId: "1",
			Action: "Action test",
		}
		err := DB.Create(&userLogs).Error
		assert.Nil(t, err)
		assert.NotEqual(t, 0, userLogs.ID)
		fmt.Println(userLogs.ID)
	}
}

func TestOneToOne(t *testing.T) {
	wallet := model.Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := DB.Create(&wallet).Error
	assert.Nil(t, err)
}
func TestRetriveRelation(t *testing.T) {
	var user model.User

	// untuk mendapatkan data walletnya juga kita panggil method preload juga
	err := DB.Model(&user).Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestRetriveRelationJoins(t *testing.T) {
	var user model.User
	// preload cocok untuk relasi one to many dll.
	// jika kita meggunakan preloadd maka querynya akan di panggil dua kali maka dari itu kita lebih baik
	// meggnuakan joins, joins hanya di panggil satu kali dan masuk ke dalam query yang sama, dan untuk memanggil join kita harus
	// mengkuery table database yang spesipik supaya tidak terjadi error
	err := DB.Model(&user).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := model.User{
		ID:       "10",
		Password: "farid anang s",
		Name: model.Name{
			FirstName: "farid",
		},
		Wallet: model.Wallet{
			ID:      "10",
			UserId:  "10",
			Balance: 100000000,
		},
	}
	// jika seperti ini maka relasinya juga akan ikut terinsert
	err := DB.Create(&user).Error
	assert.Nil(t, err)
}
func TestSkipAutoCreateUpdate(t *testing.T) {
	user := model.User{
		ID:       "11",
		Password: "farid anang s",
		Name: model.Name{
			FirstName: "farid",
		},
		Wallet: model.Wallet{
			ID:      "11",
			UserId:  "11",
			Balance: 100000000,
		},
	}
	// jika kita menambahkan Omit(clause.Associations) maka data relasinya tidak ikut di insert
	err := DB.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestOneToMany(t *testing.T) {
	user := model.User{
		ID:       "14",
		Password: "anang s",
		Name: model.Name{
			FirstName: "samsul",
		},
		Wallet: model.Wallet{
			UserId:  "14",
			Balance: 100121212,
		},
		Address: []model.Address{
			{
				UserId:  "14",
				Address: "jalan A",
			},
			{
				UserId:  "14",
				Address: "jalan B",
			},
			{
				UserId:  "14",
				Address: "jalan C",
			},
		},
	}
	err := DB.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var users []model.User
	err := DB.Model(&model.User{}).Preload("Address").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	fmt.Println(users)
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var users model.User
	err := DB.Model(&model.User{}).Preload("Address").Joins("Wallet").Take(&users, "users.id = ?", "13").Error
	assert.Nil(t, err)
	fmt.Println(users)
}

func TestBelongToAddress(t *testing.T) {
	fmt.Println("Preload")

	address := model.Address{}
	err := DB.Model(model.Address{}).Preload("User").Find(&address).Error
	assert.Nil(t, err)

	fmt.Println("Join")

	address = model.Address{}
	err = DB.Model(model.Address{}).Joins("User").Find(&address).Error
	assert.Nil(t, err)
}

func TestBelongToWallet(t *testing.T) {
	fmt.Println("Preload")

	wallet := model.Wallet{}
	err := DB.Model(model.Wallet{}).Preload("User").Find(&wallet).Error
	assert.Nil(t, err)

	fmt.Println("Join")

	wallet = model.Wallet{}
	err = DB.Model(model.Wallet{}).Joins("User").Find(&wallet).Error
	assert.Nil(t, err)
}

func TestCreateMany2Many(t *testing.T) {
	product := model.Product{
		ID:    "P001",
		Name:  "APEL",
		Price: 10000000,
	}
	err := DB.Create(&product).Error
	assert.Nil(t, err)

	err = DB.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "1",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = DB.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "2",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

}

func TestPreloadMany2Many(t *testing.T) {
	var product model.Product
	err := DB.Model(model.Product{}).Preload("LikedByUsers").Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}
func TestPreloadMany2ManyUser(t *testing.T) {
	var user model.User
	err := DB.Model(model.User{}).Preload("LikeProduct").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
}

func TestAsociationFind(t *testing.T) {
	var product model.Product
	err := DB.Model(&model.Product{}).Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []model.User
	err = DB.Model(&model.Product{}).Where("first_name like ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	fmt.Println(users)
}
func TestAsociationAppend(t *testing.T) {
	var user model.User
	err := DB.Take(&user, "id = ?", "12").Error
	assert.Nil(t, err)

	var product model.Product
	err = DB.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	// kita gunakan append untuk memasukan data ke dalam product secara otomatis
	err = DB.Model(&model.Product{}).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAsociationDelete(t *testing.T) {
	var user model.User
	err := DB.Take(&user, "id = ?", "10").Error
	assert.Nil(t, err)

	var product model.Product
	err = DB.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	// kita gunakan append untuk memasukan data ke dalam product secara otomatis
	err = DB.Model(&model.Product{}).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestPreloadingInlineCondition(t *testing.T) {
	var user model.User
	err := DB.Preload("Wallet", "balance > ?", 100).Take(&user, "id = ?", "10").Error
	assert.Nil(t, err)
	fmt.Println(user)
}
func TestPreloadingNested(t *testing.T) {
	var wallet model.Wallet
	// untuk melakukan nested kita cukup menggunakan . untuk measuk ke dalama nestednya
	err := DB.Preload("User.Address").Take(&wallet, "id = ?", "10").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Address)
}
func TestPreloadingAll(t *testing.T) {
	var wallet model.Wallet
	// clause.Associations ini kita gunakan untuk mengambil semua datanya
	err := DB.Preload(clause.Associations).Take(&wallet, "id = ?", "10").Error
	assert.Nil(t, err)

	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Address)
}

func TestJoinQuery(t *testing.T) {
	var users []model.User
	// jika kita join manual maka kiat harus menggunakan nama tablenya
	err := DB.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []model.User{}
	// jika kita joins maka kiat harus menggunakan nama fieldnya
	err = DB.Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 9, len(users))
}

func TestJoinQueryCondition(t *testing.T) {
	var users []model.User
	err := DB.Joins("join wallets on wallets.user_id = users.id and wallets.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []model.User{}
	err = DB.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestCount(t *testing.T) {
	var count int64
	err := DB.Model(&model.Wallet{}).Joins("Wallet").Where("Wallet.balance > ?", 50000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(3), count)
}
func TestQueryAgregation(t *testing.T) {
	var result AggregationResult
	err := DB.Model(&model.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(111012121), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(100000000), result.MaxBalance)
	assert.Equal(t, float64(3.70040403333e+07), result.AvgBalance)
}

func TestQueryAgregationGroupByAndHaving(t *testing.T) {
	// Group("User.Id"): Mengelompokkan hasil berdasarkan kolom Id pada tabel User. Ini berarti hasil yang dikembalikan akan dikelompokkan berdasarkan nilai unik pada kolom Id di tabel User.
	// Having("sum(balance) > ?", 1000000): Berfungsi seperti klausa WHERE dalam SQL, tetapi bekerja setelah GROUP BY. Dalam hal ini, hanya grup-grup yang memiliki jumlah saldo lebih besar dari 1.000.000 yang akan dimasukkan dalam hasil.
	var result []AggregationResult
	err := DB.Model(&model.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").Group("User.Id").Having("sum(balance) > ?", 1000000).Find(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	var users []model.User
	err := DB.WithContext(ctx).Model(&model.User{}).Preload(clause.Associations).Find(&users).Error
	assert.Nil(t, err)
	for _, data := range users {
		fmt.Println(data)
		fmt.Println("Address  ", data.Address)
		fmt.Println("Wallet  ", data.Wallet)
	}
}

// kita buat func yang bernilai gorm db dan mengembalikan gorm db untuk di panggil di dalam scopes
func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}
func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScope(t *testing.T) {
	var wallet []model.Wallet
	// kita gunakan scopes untuk memanggil func yang berisi data *gorm.DB untuk di jalankan
	// scopes juga bisa menerima lebih dari satu method DB.Scopes(method1, method2, .....)
	err := DB.Scopes(BrokeWalletBalance).Find(&wallet).Error
	assert.Nil(t, err)

	wallet = []model.Wallet{}
	err = DB.Scopes(SultanWalletBalance).Find(&wallet).Error
	assert.Nil(t, err)
}

func TestMigrator(t *testing.T) {
	// kita menggunakan migratot dan automigrate untuk membuat table pada database secara manual dan ini tidak di sarankan
	// karna sangat patal jika terjadi perubahan pada table secara tiba tiba
	err := DB.Migrator().AutoMigrate(&model.GuestBook{})
	assert.Nil(t, err)
}

// lihat penjelasannya di model user
func TestHook(t *testing.T) {
	user := model.User{
		Password: "ukiki",
		Name: model.Name{
			FirstName: "kukis",
		},
		Information: "hook 2",
	}
	err := DB.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.ID)
	fmt.Println(user.ID)

}
