package config

import (
	"fmt"

	"github.com/erikrios/go-clean-arhictecture/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func NewMySQLDatabase() (*gorm.DB, error) {
	config := Config{
		DB_Username: "root",
		DB_Password: "erikrios",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "library",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	return DB, err
}

func MigrateMySQLDatabase(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{})
}
