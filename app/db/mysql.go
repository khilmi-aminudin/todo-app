package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const projectDirName = "todo-app" // change to relevant project name

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func NewDB() *sql.DB {

	loadEnv()

	connectionstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME"))

	db, err := sql.Open("mysql", connectionstring)
	if err != nil {
		panic("error opening database connection: " + err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
