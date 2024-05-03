package model

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type SoldRecord struct {
	Model
	BookID uuid.UUID `gorm:"type:uuid;not null" json:"book_id"`
	Amount int       `gorm:"type:integer;not null" json:"amount" validate:"required, min=1"`
	TotalPrice float64 `gorm:"type:float;not null" json:"total_price" validate:"required, min=1"`
}