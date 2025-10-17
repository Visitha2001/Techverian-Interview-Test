package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/visitha2001/backend-go/handlers"
	"github.com/visitha2001/backend-go/models"
	"github.com/visitha2001/backend-go/routes"
	"github.com/visitha2001/backend-go/storage"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database configs
	dbConfig := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	db, err := storage.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to database")
	} else {
		log.Println("Database connected successfully")
	}

	// migrations
	if err := models.MigrateItems(db); err != nil {
		log.Fatal("Error migrating items")
	}

	// init app and handlers
	app := fiber.New()
	itemHandler := &handlers.ItemHandler{DB: db}

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// log all requests
	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path(), c.Response().StatusCode())
		return c.Next()
	})
	// health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// register routes
	routes.RegisterItemRoutes(app, itemHandler)

	// run app
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8084"
	}
	log.Println("Server running on port", port)
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
