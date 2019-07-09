package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Comment Type sends processed Comment data to DB in the required formats
type Comment struct {
	gorm.Model
	Nickname            string    //`gorm:"nickname"`
	ArticleID           uint      //`gorm:"article_id"`
	Content             string    //`gorm:"content"`
	CommentCreationDate time.Time //`gorm:"commentcreationdate"`
}
