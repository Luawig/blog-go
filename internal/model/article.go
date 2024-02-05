package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string    `gorm:"type:varchar(100);not null" json:"title"`
	Content      string    `gorm:"type:longtext" json:"content"`
	CreatedAt    time.Time `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null" json:"updated_at"`
	CommentCount int       `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int       `gorm:"type:int;not null;default:0" json:"read_count"`

	Comments   []*Comment  `json:"comments"`
	Categories []*Category `gorm:"many2many:article_categories;"`
}
