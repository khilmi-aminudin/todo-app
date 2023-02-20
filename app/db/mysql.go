package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connectionstring string

func NewGormDB() *gorm.DB {
	connectionstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	// db, err := sql.Open("mysql", connectionstring)
	db, err := gorm.Open(mysql.Open(connectionstring), &gorm.Config{})
	if err != nil {
		panic("error opening database connection: " + err.Error())
	}

	return db
}

func NewDB() *sql.DB {
	connectionstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))
	db, err := sql.Open("mysql", connectionstring)
	if err != nil {
		panic("error opening database connection: " + err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
