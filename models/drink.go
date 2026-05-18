package models

type Drink struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	InStock          bool    `json:"inStock"`
	ContainsCaffeine bool    `json:"containsCaffeine"`
	Volume           int     `json:"volume"`
	Description      string  `json:"description"`
}

type DrinkShort struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type DrinkInput struct {
	Name             *string  `json:"name"`
	Price            *float64 `json:"price"`
	InStock          *bool    `json:"inStock"`
	ContainsCaffeine *bool    `json:"containsCaffeine"`
	Volume           *int     `json:"volume"`
	Description      *string  `json:"description"`
}