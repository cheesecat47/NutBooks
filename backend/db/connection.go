package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("Failed to connect to MySQL: %v", err)
	}
	defer sqlDB.Close()

	// https://medium.com/@SlackBeck/golang-database-sql-패키지-삽질기-3편-커넥션-풀-a8c220f7af3d
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Panicf("Cannot connect to MySQL: %v", err)
	}

	log.Println("Connection opened to MySQL")

	MigrateMysql()
}
