package services

import (
	"context"
	"grpc_unittest/api/models"
)

type ArticleServices interface {
	AddArticle(context.Context, *models.Article) bool
	ShowArticle(context.Context, int64) (ArticleReponse, bool)
	ShowAllArticle(context.Context) (ArticleAllReponse, bool)
}

type CommentInterface interface {
	AddComment(context.Context, *models.Comment) bool
	ShowComment(context.Context, int64) (CommentResponse, bool)
	ShowAllComment(context.Context) (CommentAllResponse, bool)
	ShowArticleComment(context.Context, int64) (CommentAllResponse, bool)
}
