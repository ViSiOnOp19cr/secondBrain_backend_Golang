package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn:= "host=localhost user=postgres password=yourpassword dbname=gormdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err != nil){
		log.Fatal("failed to connect to db", err)
	}
	fmt.Println("Database connected succesfully")
	DB = db
}

