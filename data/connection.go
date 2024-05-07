package data

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	// des := "postgres://postgres:Gamersking0@localhost:5432/bookstore?sslmode=disable"
	//docker postgres
	des := os.Getenv("DATABASE_URL")
	sqlDB, err := sql.Open("pgx", des)
	if err != nil {
		return nil, errors.New("failed to connect database #1")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database #2")
	}
	
	fmt.Println("Database connected")
	return gormDB, nil
}
