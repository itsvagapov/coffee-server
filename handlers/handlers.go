package handlers

import (
	"coffee-server/models"
	"coffee-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDrinks(c *gin.Context) {
	drinks, err := service.GetDrinks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка напитков"})
		return
	}
	c.JSON(http.StatusOK, drinks)
}


func GetDrinksInStock(c *gin.Context) {
	drinks, err := service.GetDrinksInStock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении напитков в наличии"})
		return
	}
	c.JSON(http.StatusOK, drinks)
}


func GetDrinkByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	drink, err := service.GetDrinkByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Напиток не найден"})
		return
	}
	c.JSON(http.StatusOK, drink)
}


func CreateDrink(c *gin.Context) {
	var input models.DrinkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	drink, err := service.CreateDrink(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, drink)
}

func UpdateDrink(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	var input models.DrinkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	drink, err := service.UpdateDrink(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drink)
}


func DeleteDrink(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	if err := service.DeleteDrink(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Напиток успешно удалён"})
}

func parseID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return 0, err
	}
	return id, nil
}