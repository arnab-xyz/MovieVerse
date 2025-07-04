package model

type Review struct {
	ID      string `json:"id" gorm:"primaryKey"`
	MovieID string `json:"movieId" gorm:"index"`
	Movie   Movie  `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating" validate:"min=0,max=10"`
}
