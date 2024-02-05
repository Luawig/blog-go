package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string    `gorm:"type:varchar(500);not null;" json:"content"`
	CreatedAt time.Time `gorm:"type:datetime;not null" json:"created_at"`

	Article   *Article `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE" json:"article"`
	ArticleID uint     `gorm:"type:int;not null" json:"article_id"`
	User      *User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	UserID    uint     `gorm:"type:int;not null" json:"user_id"`
}
