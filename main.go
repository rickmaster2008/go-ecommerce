package main

import (
	"log"
	"net/http"
	"newproject/app"
	"newproject/controllers"
	"newproject/database"
	"newproject/middlewares"
	"newproject/models"
	"newproject/response"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No se pudo cargar el archivo .env")
		return
	}
	db := database.Open()
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Customer{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.OrderItem{}, &models.Status{})

	app := app.NewApp()

	app.Static("./static", "/static/")

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users := []models.User{}
		db.Find(&users)
		response.Render(w, users, "templates/home/index.html")
	})
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
	app.Resource("/products", controllers.BaseController{M: models.Product{}}, middlewares.AuthMiddleware)
	app.Resource("/categories", controllers.BaseController{M: models.Category{}})
	app.Resource("/users", controllers.BaseController{M: models.User{}}, middlewares.AuthMiddleware)
	app.Listen(":8080")
}
