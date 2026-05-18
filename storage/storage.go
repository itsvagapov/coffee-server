package storage

import (
	"coffee-server/models"
	"errors"
)

var drinks []models.Drink

var nextID = 1

func init() {
	drinks = []models.Drink{
		{
			ID:               nextID,
			Name:             "Эспрессо",
			Price:            120,
			InStock:          true,
			ContainsCaffeine: true,
			Volume:           30,
			Description:      "Крепкий итальянский кофе с насыщенным вкусом и плотной пенкой.",
		},
		{
			ID:               nextID + 1,
			Name:             "Капучино",
			Price:            180,
			InStock:          true,
			ContainsCaffeine: true,
			Volume:           200,
			Description:      "Нежный кофейный напиток с молочной пеной.",
		},
		{
			ID:               nextID + 2,
			Name:             "Латте",
			Price:            200,
			InStock:          false,
			ContainsCaffeine: true,
			Volume:           300,
			Description:      "Мягкий кофе с большим количеством молока.",
		},
		{
			ID:               nextID + 3,
			Name:             "Ромашковый чай",
			Price:            90,
			InStock:          true,
			ContainsCaffeine: false,
			Volume:           250,
			Description:      "Успокаивающий травяной чай без кофеина.",
		},
	}
	nextID += 4
}

func GetAll() []models.Drink {
	return drinks
}

func GetInStock() []models.Drink {
	result := []models.Drink{}
	for _, d := range drinks {
		if d.InStock {
			result = append(result, d)
		}
	}
	return result
}

func GetByID(id int) (models.Drink, error) {
	for _, d := range drinks {
		if d.ID == id {
			return d, nil
		}
	}
	return models.Drink{}, errors.New("напиток не найден")
}

func Create(input models.DrinkInput) (models.Drink, error) {
	if input.Name == nil || input.Price == nil {
		return models.Drink{}, errors.New("поля name и price обязательны")
	}

	drink := models.Drink{
		ID:    nextID,
		Name:  *input.Name,
		Price: *input.Price,
	}
	nextID++

	if input.InStock != nil {
		drink.InStock = *input.InStock
	}
	if input.ContainsCaffeine != nil {
		drink.ContainsCaffeine = *input.ContainsCaffeine
	}
	if input.Volume != nil {
		drink.Volume = *input.Volume
	}
	if input.Description != nil {
		drink.Description = *input.Description
	}

	drinks = append(drinks, drink)
	return drink, nil
}

func Update(id int, input models.DrinkInput) (models.Drink, error) {
	for i, d := range drinks {
		if d.ID == id {
			if input.Name != nil {
				drinks[i].Name = *input.Name
			}
			if input.Price != nil {
				drinks[i].Price = *input.Price
			}
			if input.InStock != nil {
				drinks[i].InStock = *input.InStock
			}
			if input.ContainsCaffeine != nil {
				drinks[i].ContainsCaffeine = *input.ContainsCaffeine
			}
			if input.Volume != nil {
				drinks[i].Volume = *input.Volume
			}
			if input.Description != nil {
				drinks[i].Description = *input.Description
			}
			return drinks[i], nil
		}
	}
	return models.Drink{}, errors.New("напиток не найден")
}

func Delete(id int) error {
	for i, d := range drinks {
		if d.ID == id {
			drinks = append(drinks[:i], drinks[i+1:]...)
			return nil
		}
	}
	return errors.New("напиток не найден")
}