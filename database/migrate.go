package database

import (
	"fmt"

	"gorm.io/gorm"

	model "main/internal/models"
)


func migration(DB *gorm.DB) {
	DB.AutoMigrate(&model.Note{})
    fmt.Println("Database Migrated")
}
