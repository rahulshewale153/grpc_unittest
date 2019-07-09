package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Article Type sends processed Article data to DB in the required formats
type Article struct {
	gorm.Model
	Nickname            string    //`gorm:"nickname"`
	Title               string    //`gorm:"title"`
	ArticleCreationDate time.Time //`gorm:"article_creation_date"`
	Content             string    //`gorm:"content"`
}
