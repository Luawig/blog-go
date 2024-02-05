package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" validate:"required,min=4,max=12" label:"Username"`
	Password string `gorm:"type:varchar(120);not null" json:"password" validate:"required,min=6,max=100" label:"Password"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email" label:"Email"`

	CreatedAt   time.Time `gorm:"type:datetime;not null" json:"created_at"`
	LastLoginAt time.Time `gorm:"type:datetime" json:"last_login_at"`

	Comments []*Comment `json:"comments"`
}
