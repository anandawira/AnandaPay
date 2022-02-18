package repo

import "gorm.io/gorm"

type bookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(Conn *gorm.DB)