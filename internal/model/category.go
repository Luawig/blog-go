package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null" json:"name"`

	Articles []*Article `gorm:"many2many:article_categories;"`
}
