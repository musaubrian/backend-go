package models

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

type TinyUrl struct {
    ID uint64 `json:"id" gorm:"primaryKey"`
    RedirectUrl string `json:"redirectUrl" gorm:"not null"`
    ShortUrl string `json:"url" gorm:"not null"`
}

func Setup()  {
    postgreDb := ""

    db, err := gorm.Open(postgres.Open(postgreDb), &gorm.Config())

    if err != nil {
        log.Fatal("Could not open db:", err)
        panic(err)
    } 

    err = db.AutoMigrate(&TinyUrl{})
}
