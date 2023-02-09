package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/khilmi-aminudin/todo-app/app"
)

// const projectDirName = "todo-app" // change to relevant project name

// func loadEnv() {

// 	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
// 	currentWorkDirectory, _ := os.Getwd()
// 	rootPath := projectName.Find([]byte(currentWorkDirectory))
// 	err := godotenv.Load(string(rootPath) + "/.env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

func main() {
	// loadEnv()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.RunServer()
}
