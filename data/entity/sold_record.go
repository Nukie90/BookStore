package entity

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type SoldRecord struct {
	Model
	BuyerID uuid.UUID `gorm:"type:uuid;not null" json:"buyer_id"`
	BookList []Book `gorm:"many2many:sold_record_books;" json:"book_list"`
	TotalPrice float64 `gorm:"type:numeric(10,2);not null" json:"total_price"`
}