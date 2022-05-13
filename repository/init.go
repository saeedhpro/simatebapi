package repository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/helpers/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type db struct {
	MySQL *gorm.DB
}

var (
	DB db
)

func Init() {
	username := env.GetEnv("USERNAME")
	fmt.Println(username)
	password := env.GetEnv("PASSWORD")
	host := env.GetEnv("HOST")
	port, _ := strconv.Atoi(env.GetEnv("DBPORT"))
	schema := env.GetEnv("SCHEMA")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected to db")
	DB.MySQL = db
}
