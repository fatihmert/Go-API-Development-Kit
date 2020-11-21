package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"apidev.fatihmert.dev/controllers"
	"apidev.fatihmert.dev/states"

	//3rd
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

var PORT string
var MYSQL_USER string
var MYSQL_PASSWORD string
var MYSQL_DATABASE string
var TCP_IP string
var MYSQL_PORT string

func prepareDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("API_PORT")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	TCP_IP = os.Getenv("TCP_IP")

	if TCP_IP == "localhost" || TCP_IP == "0.0.0.0" {
		TCP_IP = "127.0.0.1"
	}

	MYSQL_PORT = os.Getenv("MYSQL_PORT")
}

func initDatabase() {
	var err error
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", MYSQL_USER, MYSQL_PASSWORD, TCP_IP, MYSQL_PORT, MYSQL_DATABASE)

	states.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("Failed connect db")
	}

	fmt.Println("Connection Opened to db")
}

func setupRoutes(app *fiber.App) {
	// local cors middleware
	app.Use(cors.New())

	// API
	api := app.Group("/api")

	// Authentication
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)

	// Hello World
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Hello World!",
		})
	})
}

func main() {
	prepareDotEnv()

	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	log.Fatal(app.Listen(":" + PORT))
}
