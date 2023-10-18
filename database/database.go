package database

import (
	"fmt"
	"log"
	"os"

	"github.com/abdulmanafc2001/First-Project-Ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	// // dsn := os.Getenv("DB_URL")
	// // fmt.Println(dsn)

	dns := os.Getenv("DB_URL")
	fmt.Println(dns)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Catagory{})
	DB.AutoMigrate(&models.Brand{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderItem{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.RazorPay{})
	DB.AutoMigrate(&models.Image{})
	DB.AutoMigrate(&models.Coupon{})
	DB.AutoMigrate(&models.Catagory_Offer{})
	DB.AutoMigrate(&models.Wishlist{})
}
