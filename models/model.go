package models

import (
	"log"

	//"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type TinyUrl struct {
    ID uint64 `json:"id" gorm:"primaryKey"`
    RedirectUrl string `json:"redirectUrl" gorm:"not null"`
    ShortUrl string `json:"url" gorm:"not null"`
    SuccessfulRedirects uint64 `json:"successfulRedirects"`
}

func Setup(){
    postgresdb := "host=containers-us-west-16.railway.app user=postgres password=lmcoCUMkVLR6cgOt5mw8 dbname=railway port=5923"
    
    var err error
    db, err = gorm.Open(postgres.Open(postgresdb), &gorm.Config{})
    if err != nil {
        log.Fatal("Could not open database: ", err)
        panic(err)
    }

    err = db.AutoMigrate(&TinyUrl{})
    if err != nil {
        log.Fatal("Error encountered:", err)
    }
}
