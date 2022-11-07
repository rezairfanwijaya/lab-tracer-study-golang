package main

import (
	"fmt"
	"log"

	"github.com/rezairfanwijaya/lab-tracer-study-golang/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// get env value
	env, err := helper.GetENV()
	if err != nil {
		log.Println(err)
	}
	username := env["DATABASE_USERNAME"]
	host := env["DATABASE_HOST"]
	port := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	// open onnection
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("connect to database", db)
}
