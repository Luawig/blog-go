package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string    `gorm:"type:varchar(500);not null;" json:"content"`
	Time    time.Time `gorm:"type:datetime;not null" json:"time"`

	Article   Article `gorm:"foreignKey:ArticleID" json:"article"`
	ArticleID uint    `gorm:"type:int;not null" json:"article_id"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	UserID    uint    `gorm:"type:int;not null" json:"user_id"`
}
