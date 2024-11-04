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
	dsn := "admin:admin@2004@tcp(127.0.0.1:3306)/db_latihan?charset=utf8mb4&parseTime=True&loc=Local"
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
