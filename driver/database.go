package driver

import (
	"java-agro/features/email"
	"java-agro/features/faq"
	"java-agro/features/highlights"
	"java-agro/features/owners"
	"java-agro/features/products"
	"java-agro/features/question"
	"java-agro/features/user"
	"java-agro/features/youtube"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "admin:Rahasia12345@tcp(database-1.cll1d4icdkat.us-east-2.rds.amazonaws.com)/javaAgro?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&user.User{}, &email.Email{}, &highlights.Highlights{}, &faq.Faq{}, &owners.Owners{}, &products.Products{}, &question.Question{}, &youtube.Youtube{})
}
