package repository

import (
	"context"
	"grpc_unittest/api/models"
)

type ArticleRepository interface {
	Save(context.Context, *models.Article) bool
	Show(context.Context, int64) (models.Article, bool)
	ShowAll(context.Context) ([]models.Article, bool)
}

type CommentRepositoryInterface interface {
	Save(context.Context, *models.Comment) bool
	Show(context.Context, int64) (models.Comment, bool)
	ShowAll(context.Context) ([]models.Comment, bool)
	ShowArticleComment(context.Context, int64) ([]models.Comment, bool)
}
