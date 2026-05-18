package main

import (
	"coffee-server/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	drinks := r.Group("/drinks")
	{
		drinks.GET("", handlers.GetDrinks)

		drinks.GET("/in-stock", handlers.GetDrinksInStock)

		drinks.GET("/:id", handlers.GetDrinkByID)

		drinks.POST("", handlers.CreateDrink)

		drinks.PATCH("/:id", handlers.UpdateDrink)

		drinks.DELETE("/:id", handlers.DeleteDrink)
	}

	log.Println("Сервер запущен на http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}