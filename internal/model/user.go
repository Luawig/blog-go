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

	CreateTime    time.Time `gorm:"type:datetime;not null" json:"create_time"`
	LastLoginTime time.Time `gorm:"type:datetime;not null" json:"last_login"`

	Comments []*Comment `json:"comments"`
}
