package services

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/repository"
	"time"
)

//ArticleReponse Type  sends the only required Article data to frontend in the specified formats
type ArticleReponse struct {
	ID                  uint      `json:"id"`
	Nickname            string    `json:"nickname"`
	Title               string    `json:"title"`
	ArticleCreationDate time.Time `json:"articlecreationdate"`
	Content             string    `json:"content"`
}

//ArticleReponse Type  sends the only required Article data to frontend in the specified formats
type ArticleAllReponse struct {
	Articles []ArticleReponse `json:"articles"`
}

//UserRepository to manage users persistence
type ArticleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository: articleRepo}
}

func (a ArticleService) AddArticle(ctx context.Context, ar *models.Article) bool {
	articleCreate := a.articleRepository.Save(ctx, ar)
	if !articleCreate {
		return false
	}
	return true
}
func (a ArticleService) ShowArticle(ctx context.Context, id int64) (ArticleReponse, bool) {
	article := ArticleReponse{}
	articleData, dataFlag := a.articleRepository.Show(ctx, id)
	if !dataFlag {
		return article, false
	}
	article.ID = articleData.ID
	article.Nickname = articleData.Nickname
	article.Title = articleData.Title
	article.ArticleCreationDate = articleData.ArticleCreationDate
	article.Content = articleData.Content
	return article, true
}
func (a ArticleService) ShowAllArticle(ctx context.Context) (ArticleAllReponse, bool) {
	var articleAllResponse ArticleAllReponse
	articleData, dataFlag := a.articleRepository.ShowAll(ctx)
	if !dataFlag {
		return articleAllResponse, false
	}
	article := ArticleReponse{}
	for _, value := range articleData {
		article.ID = value.ID
		article.Nickname = value.Nickname
		article.Title = value.Title
		article.ArticleCreationDate = value.ArticleCreationDate
		article.Content = value.Content
		articleAllResponse.Articles = append(articleAllResponse.Articles, article)
	}
	return articleAllResponse, true
}
