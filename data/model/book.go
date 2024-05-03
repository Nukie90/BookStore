package model

import (

	_ "github.com/lib/pq"
)

type Book struct {
	Model
	Title       string    `gorm:"type:varchar(100);not null" json:"title" validate:"required, min=4, max=100"`
	Author      string    `gorm:"type:varchar(100);not null" json:"author" validate:"required, min=4, max=100"`
	Description string    `gorm:"type:text;not null" json:"description" validate:"required, min=4"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price" validate:"required, min=1"`
	Stock       int       `gorm:"type:integer;not null" json:"stock" validate:"required, min=1"`
}