package database

import (
	"fmt"
	"main/config"
	"main/utils/helpers"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

func ConnectDB() {
	var err error

	db_port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)
	helpers.CheckError(err, "")

	// Connection URL to connect to Postgres Database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		config.Config("DB_HOST"), 
		db_port, 
		config.Config("DB_USER"), 
		config.Config("DB_PASSWORD"), 
		config.Config("DB_NAME"),
	)
    
	fmt.Println("connectionString = " + connectionString)
	DB, err = gorm.Open(postgres.Open(connectionString))
	helpers.CheckError(err, "")

	fmt.Println("Connected to Database :)")
}