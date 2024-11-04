package config

import (
	"compro-backend/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// ConnectDatabase menginisialisasi koneksi database
func ConnectDatabase() {
	// Ubah sesuai dengan kredensial database Anda
	dsn := "sitb1411_user:Sitetech@2024@tcp(sitetech.my.id:3306)/sitb1411_db_golang?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("Database connected!")

	// Migrasi model
	database.AutoMigrate(
		&models.Categories{},
		&models.Articles{},
	)
}
