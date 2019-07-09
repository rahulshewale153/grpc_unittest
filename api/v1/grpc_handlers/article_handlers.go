package grpc_handlers

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/grpc/article_grpc"
	"net/http"
	"time"
)

const (
	DATEFORMAT string = "2006-01-02"
)

type ArticlesServer struct {
	ArticleService services.ArticleServices
}

func NewArticleHttpHandler(articleService services.ArticleServices) *ArticlesServer {
	return &ArticlesServer{ArticleService: articleService}

}
func (as ArticlesServer) Articlelist(ctx context.Context, void *article_grpc.Void) (*article_grpc.ArticleResponse, error) {
	var articles article_grpc.ArticleList
	var articleresponse article_grpc.ArticleResponse
	articledb, dataflag := as.ArticleService.ShowAllArticle(ctx)
	if !dataflag {
		articleresponse.Status = http.StatusOK
		articleresponse.Success = true
		articleresponse.Message = "Record not present"
		return &articleresponse, nil
	}
	for _, articledb := range articledb.Articles {
		var article article_grpc.Article
		article.Id = uint64(articledb.ID)
		article.Nickname = articledb.Nickname
		article.Title = articledb.Title
		article.Articlecreationdate = articledb.ArticleCreationDate.String()
		article.Content = articledb.Content
		articles.Articles = append(articles.Articles, &article)
	}

	articleresponse.Status = http.StatusOK
	articleresponse.Success = true
	articleresponse.Message = "Record present"
	articleresponse.Articles = articles.Articles

	return &articleresponse, nil
}

func (as ArticlesServer) Addarticle(ctx context.Context, articleData *article_grpc.AddArticle) (*article_grpc.ArticleResponse, error) {
	var articleresponse article_grpc.ArticleResponse
	datedb, err := time.Parse(DATEFORMAT, articleData.Articlecreationdate)
	if err != nil {
		articleresponse.Status = http.StatusOK
		articleresponse.Success = false
		articleresponse.Message = "Invalid date format"
		return &articleresponse, nil
	}
	var articledb models.Article
	articledb.Nickname = articleData.Nickname
	articledb.Title = articleData.Title
	articledb.ArticleCreationDate = datedb
	articledb.Content = articleData.Content
	flag := as.ArticleService.AddArticle(ctx, &articledb)
	if !flag {
		articleresponse.Status = http.StatusBadRequest
		articleresponse.Success = false
		articleresponse.Message = "Something is wrong Record not Save"
		return &articleresponse, nil
	}
	articleresponse.Status = http.StatusOK
	articleresponse.Success = true
	articleresponse.Message = "Record save sucessfully"
	return &articleresponse, nil
}
func (as ArticlesServer) Searcharticle(ctx context.Context, searcchArticle *article_grpc.SearchArticle) (*article_grpc.ArticleResponse, error) {
	var articleresponse article_grpc.ArticleResponse
	articledb, dataflag := as.ArticleService.ShowArticle(ctx, searcchArticle.Id)
	if !dataflag {
		articleresponse.Status = http.StatusOK
		articleresponse.Success = true
		articleresponse.Message = "Record not present"
		return &articleresponse, nil
	}
	var article article_grpc.Article
	article.Id = uint64(articledb.ID)
	article.Nickname = articledb.Nickname
	article.Title = articledb.Title
	article.Articlecreationdate = articledb.ArticleCreationDate.String()
	article.Content = articledb.Content

	articleresponse.Status = http.StatusOK
	articleresponse.Success = true
	articleresponse.Message = "Record present"
	articleresponse.Articles = append(articleresponse.Articles, &article)
	return &articleresponse, nil
}
