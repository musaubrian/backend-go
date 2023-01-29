// package models defines the model structure of how the data
// will be stored in the database
package models

import (
	"log"
	"os"

	// "os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db is an instance of a gorm DB
// https://gorm.io/gorm
var db *gorm.DB

// TinyUrl defines the structure of how the data will be stored
// in the database
type TinyUrl struct {
	ID                  uint64 `json:"id" gorm:"primaryKey"`
	RedirectUrl         string `json:"redirectUrl" gorm:"not null"`
	ShortUrl            string `json:"url" gorm:"not null"`
	SuccessfulRedirects uint64 `json:"successfulRedirects"`
}

// Setup() sets up the connection to the database
// and initializes the structure (table name and columns)
func Setup() {

	var err error
    if err := godotenv.Load(); err != nil {
        log.Fatal("Could not load .env")
    }
    dsn := os.Getenv("DSN") 
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open database: ", err)
		panic(err)
	}

	err = db.AutoMigrate(&TinyUrl{})
	if err != nil {
		log.Fatal("Could not create tables", err)
	}
}
