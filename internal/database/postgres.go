package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"zadaie/internal/model"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	var db *gorm.DB
	var errDB error

	for i := 1; i <= 10; i++ {
		db, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if errDB == nil {
			log.Println("DB connected successfully")
			break
		}

		log.Printf("DB not ready (attempt %d/10): %v\n", i, errDB)
		time.Sleep(2 * time.Second)
	}

	if errDB != nil {
		panic("failed to connect to database after retries: " + errDB.Error())
	}

	// 👇 ВАЖНО: используем гарантированно подключённый db
	err = db.AutoMigrate(&model.Subscription{})
	if err != nil {
		panic(err)
	}

	DB = db
}
