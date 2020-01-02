package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        	uint64    	`gorm:"primary_key;auto_increment" json:"id"`
	Name		string    	`gorm:"size:255;not null;" json:"name"`
	Email     	string    	`gorm:"size:100;not null;unique" json:"email"`
	Password  	string    	`gorm:"size:100;not null;" json:"password"`
	CreatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

func (u *User) getAll(db *gorm.DB) (*[]User, error) {
	var err error
	var users = []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}