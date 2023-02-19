package database

import (
	"gorm.io/gorm"
	"fmt"

	
	model "main/internal/models"
)


func migration(DB *gorm.DB) {
	DB.AutoMigrate(&model.Note{})
    fmt.Println("Database Migrated")
}
