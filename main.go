package main

import (
	"log"
	// swagger embed files
	// gin-swagger middleware

	"github.com/joho/godotenv"

	"github.com/khilmi-aminudin/todo-app/app"
	"github.com/khilmi-aminudin/todo-app/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "todo-list app"
	docs.SwaggerInfo.Description = "this is a simple todo list api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.RunServer()
}
