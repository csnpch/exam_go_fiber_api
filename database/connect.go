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
	
	db_port, err := strconv.ParseUint(config.ConfigENV("DB_PORT", "5432"), 10, 32)
	helpers.CheckError(err, "")

	// Connection URL to connect to Postgres Database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		config.ConfigENV("DB_HOST", "127.0.0.1"), 
		db_port, 
		config.ConfigENV("DB_USER", "postgres"), 
		config.ConfigENV("DB_PASSWORD", "csnp332211"), 
		config.ConfigENV("DB_NAME", "db_ituy"),
	)
    
	fmt.Println("connectionString = " + connectionString)
	DB, err = gorm.Open(postgres.Open(connectionString))
	helpers.CheckError(err, "")

	fmt.Println("Connected to Database :)")
}