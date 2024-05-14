package infrastructure

import (
	"fmt"
	"main/data/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type GormConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewGormConfig(v *viper.Viper) *GormConfig {
	return &GormConfig{
		Driver:   v.GetString("postgres.connection.driver"),
		Host:     v.GetString("postgres.connection.host"),
		Port:     v.GetInt("postgres.connection.port"),
		User:     v.GetString("postgres.connection.user"),
		Password: v.GetString("postgres.connection.password"),
		DBName:   v.GetString("postgres.connection.dbname"),
	}
}

func (gc *GormConfig) Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", gc.Host, gc.User, gc.Password, gc.DBName, gc.Port)
	if gc.Driver == "postgres" {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	return nil, fmt.Errorf("unsupported driver: %s", gc.Driver)
}

func (gc *GormConfig) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.ShoppingCart{})
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.SoldRecord{})
}