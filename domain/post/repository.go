package post

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository {
	return repository{
		db,
	}
}

type Repository interface {
}
