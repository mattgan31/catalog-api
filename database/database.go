package database

import (
	"fmt"
	"log"

	"example.com/m/v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "alam"
	dbport   = 5432
	dbname   = "catalog_api"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.Product{})
	db.Debug().AutoMigrate(models.User{})

}

func GetDB() *gorm.DB {
	return db
}
