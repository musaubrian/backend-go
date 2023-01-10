//package models defines the model structure of how the data
//will be stored in the database
package models

import (
//	"fmt"
	"log"
//	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//db is an instance of a gorm DB
//
//[Gorm]: https://gorm.io/gorm
var db *gorm.DB

//TinyUrl defines the structure of how the data will be stored
//in the database
type TinyUrl struct {
    ID uint64 `json:"id" gorm:"primaryKey"`
    RedirectUrl string `json:"redirectUrl" gorm:"not null"`
    ShortUrl string `json:"url" gorm:"not null"`
    SuccessfulRedirects uint64 `json:"successfulRedirects"`
}

//Setup() sets up the connection to the database
//and initializes the structure (table name and columns)
func Setup(){

    envErr := godotenv.Load()
    if envErr != nil {
        log.Fatal("Could not load environment file", envErr)
    }

    //PGDATABASE := fmt.Printf("database=",os.Getenv("PGDATABASE"))
   // PGHOST := os.Getenv("PGHOST")
    //PGPASSWORD := os.Getenv("PGPASSWORD")
    //PGPORT := os.Getenv("PGPORT")
    //PGUSER := os.Getenv("PGUSER")

    //postgresSetup := fmt.Printf("host=%s user=%s password=%s dbname=%s port=%d ",PGHOST, PGUSER, PGPASSWORD, PGDATABASE, PGPORT)
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
