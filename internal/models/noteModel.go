package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
    gorm.Model           // Adds some metadata fields to the table
    ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
    Title      string
    SubTitle   string
    Text       string
}


type UpdateNote struct {
    Title    string `json:"title"`
    SubTitle string `json:"sub_title"`
    Text     string `json:"Text"`
}