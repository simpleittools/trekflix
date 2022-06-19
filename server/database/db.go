package database

import (
	"fmt"
	"github.com/simpleittools/trekflix/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// Conn is used to connect to the database and manage the connection
func Conn() {
	dsn := os.Getenv("PGSQLDSN")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the DB")
	} else {
		fmt.Println("connected to PGSQL")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
