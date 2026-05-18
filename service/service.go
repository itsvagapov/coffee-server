package service

import (
	"coffee-server/models"
	"coffee-server/storage"
)

func GetDrinks() ([]models.DrinkShort, error) {
	all := storage.GetAll()
	result := make([]models.DrinkShort, 0, len(all))
	for _, d := range all {
		result = append(result, models.DrinkShort{
			ID:    d.ID,
			Name:  d.Name,
			Price: d.Price,
		})
	}
	return result, nil
}

func GetDrinksInStock() ([]models.DrinkShort, error) {
	inStock := storage.GetInStock()
	result := make([]models.DrinkShort, 0, len(inStock))
	for _, d := range inStock {
		result = append(result, models.DrinkShort{
			ID:    d.ID,
			Name:  d.Name,
			Price: d.Price,
		})
	}
	return result, nil
}

func GetDrinkByID(id int) (models.Drink, error) {
	return storage.GetByID(id)
}

func CreateDrink(input models.DrinkInput) (models.Drink, error) {
	return storage.Create(input)
}

func UpdateDrink(id int, input models.DrinkInput) (models.Drink, error) {
	return storage.Update(id, input)
}

func DeleteDrink(id int) error {
	return storage.Delete(id)
}