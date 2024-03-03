package main 

import (
	"github.com/gofiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB : db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}