package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Model struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	Model
	Username string `gorm:"type:varchar(100);not null;unique" json:"username" validate:"required, min=4, max=100"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required, min=8, max=100"`
	UserType string `gorm:"type:varchar(100);not null" json:"user_type" validate:"required, oneof=Owner Customer"`
	ShoppingCart []ShoppingCart `gorm:"foreignKey:UserID" json:"shopping_cart"`
}

type ShoppingCart struct {
	Model
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	BookID uuid.UUID `gorm:"type:uuid;not null" json:"book_id" validate:"required"`
	Title string `gorm:"type:varchar(100);not null" json:"title" validate:"required"`
	Quantity int `gorm:"type:int;not null" json:"quantity" validate:"required"`
	Cost float64 `gorm:"type:float;not null" json:"cost" validate:"required"`
}