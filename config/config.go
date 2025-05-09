package config

import (
	"fmt"
	"os"

	"GeoTagger/models" // ✅ Import models

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN") // "user:pass@tcp(localhost:3306)/geotagger"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("DB connected")

	DB = database
	// ✅ Reference structs from the models package
	DB.AutoMigrate(&models.User{}, &models.Note{})
}

func GetDB() *gorm.DB {
	return DB
}
