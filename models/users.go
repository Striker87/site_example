package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound = errors.New("models: resource not found")
)

type UserService struct {

}

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}

// ByID will look up the id provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := db.Where("id = ?").First(&user).Error
}
