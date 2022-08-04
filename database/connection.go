package database

import (
	"fmt"
	"task/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open("mysql", "user:user@123@tcp(mysql:3306)/task_db")

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!!!")
	}
	connection.AutoMigrate(&models.User{})
	DB = connection
}