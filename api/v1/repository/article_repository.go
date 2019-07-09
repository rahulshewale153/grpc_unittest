package repository

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/configs"
	"grpc_unittest/database/connection"
)

type ArticleCommentRepository struct {
	ConnectionService connection.ConnectionInterface
}

func NewArticleRepository(connectionService connection.ConnectionInterface) *ArticleCommentRepository {
	return &ArticleCommentRepository{connectionService}
}

func (a ArticleCommentRepository) Save(ctx context.Context, ar *models.Article) bool {
	dbconn := a.ConnectionService.DBConnect()
	defer dbconn.Close()
	articleCreate := dbconn.Create(ar)
	if articleCreate.Error != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "article not create: ", articleCreate.Error)
		return false
	}
	return true
}

func (a ArticleCommentRepository) Show(ctx context.Context, id int64) (models.Article, bool) {
	dbconn := a.ConnectionService.DBConnect()
	defer dbconn.Close()
	articledb := models.Article{}
	articledata := dbconn.Where("id=?", id).First(&articledb)
	if articledata.RowsAffected <= 0 {
		return articledb, false
	}
	return articledb, true
}

func (a ArticleCommentRepository) ShowAll(ctx context.Context) ([]models.Article, bool) {
	dbconn := a.ConnectionService.DBConnect()
	defer dbconn.Close()
	articledb := []models.Article{}
	articledata := dbconn.Find(&articledb)
	if articledata.RowsAffected <= 0 {
		return articledb, false
	}
	return articledb, true
}
