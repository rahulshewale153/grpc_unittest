package repository

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/configs"
	"grpc_unittest/database/connection"
)

//UserRepository to manage users persistence
type CommentRepository struct {
	connectionInterface connection.ConnectionInterface
}

func NewCommentRepository(connectionInterface connection.ConnectionInterface) *CommentRepository {
	return &CommentRepository{connectionInterface}
}

func (a CommentRepository) Save(ctx context.Context, comment *models.Comment) bool {
	dbconn := a.connectionInterface.DBConnect()
	defer dbconn.Close()
	CommentCreate := dbconn.Create(comment)
	if CommentCreate.Error != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "comment not create: ", CommentCreate.Error)
		return false
	}
	return true
}

func (a CommentRepository) Show(ctx context.Context, id int64) (models.Comment, bool) {
	dbconn := a.connectionInterface.DBConnect()
	defer dbconn.Close()
	commentdb := models.Comment{}
	commentdata := dbconn.Where("id=?", id).First(&commentdb)
	if commentdata.RowsAffected <= 0 {
		return commentdb, false
	}
	return commentdb, true
}

func (a CommentRepository) ShowAll(ctx context.Context) ([]models.Comment, bool) {
	dbconn := a.connectionInterface.DBConnect()
	defer dbconn.Close()
	commentdb := []models.Comment{}
	commentdata := dbconn.Find(&commentdb)
	if commentdata.RowsAffected <= 0 {
		return commentdb, false
	}
	return commentdb, true
}

func (a CommentRepository) ShowArticleComment(ctx context.Context, id int64) ([]models.Comment, bool) {
	dbconn := a.connectionInterface.DBConnect()
	defer dbconn.Close()
	commentdb := []models.Comment{}
	commentdata := dbconn.Where("article_id=?", id).Find(&commentdb)
	if commentdata.RowsAffected <= 0 {
		return commentdb, false
	}
	return commentdb, true
}
