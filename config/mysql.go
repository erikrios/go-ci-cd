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
		DB_Username: "b3e103f75ab23c",
		DB_Password: "4f3b2a25",
		DB_Port:     "3306",
		DB_Host:     "us-cdbr-east-05.cleardb.net",
		DB_Name:     "heroku_41e2f92743a9807",
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
