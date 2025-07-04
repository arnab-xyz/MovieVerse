package model

type Movie struct {
	ID          string   `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Reviews     []Review `json:"reviews" gorm:"foreignKey:MovieID"`
}
