package services

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/configs"
	"reflect"
	"testing"
	"time"
)

type mock_ArticleRepository struct {
}

func init() {
	configs.Config.Read("testing")
}
func GetDummyDate() time.Time {
	datedb, err := time.Parse("2006-01-02", "2019-02-02")
	if err != nil {
	}
	return datedb
}
func (_m *mock_ArticleRepository) Show(ctx context.Context, id int64) (models.Article, bool) {
	var record1 models.Article
	if id == 1 {
		record1.ID = 1
		record1.Nickname = "Rahul Shewale"
		record1.Title = "ABCD"
		record1.ArticleCreationDate = GetDummyDate()
		record1.Content = "ABSCD"
		return record1, true
	} else {
		return record1, false
	}
}
func (_m *mock_ArticleRepository) Save(ctx context.Context, ar *models.Article) bool {

	if ar.Title != "" {
		return true
	} else {
		return false
	}
}
func (_m *mock_ArticleRepository) ShowAll(ctx context.Context) ([]models.Article, bool) {
	var r0 []models.Article
	var record1 models.Article
	record1.ID = 1
	record1.Nickname = "Rahul Shewale"
	record1.Title = "ABCD"
	record1.ArticleCreationDate = GetDummyDate()
	record1.Content = "ABSCD"
	r0 = append(r0, record1)
	return r0, true

}

func TestArticleService_AddArticle(t *testing.T) {
	articleService := NewArticleService(&mock_ArticleRepository{})
	type args struct {
		ar *models.Article
	}

	var record1 models.Article
	record1.Nickname = "Rahul Shewale"
	record1.Title = "ABCD"
	record1.ArticleCreationDate = GetDummyDate()
	record1.Content = "ABSCD"

	var record2 models.Article

	tests := []struct {
		name   string
		fields *ArticleService
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"Save record", articleService, args{&record1}, true},
		{"Save record", articleService, args{&record2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleService{
				articleRepository: tt.fields.articleRepository,
			}
			if got := a.AddArticle(context.Background(), tt.args.ar); got != tt.want {
				t.Errorf("ArticleService.AddArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleService_ShowArticle(t *testing.T) {
	articleService := NewArticleService(&mock_ArticleRepository{})
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *ArticleService
		args   args
		want   ArticleReponse
		want1  bool
	}{
		// TODO: Add test cases.
		{"Show Article Service", articleService, args{1}, ArticleReponse{1, "Rahul Shewale", "ABCD", GetDummyDate(), "ABSCD"}, true},
		{"Show Article Service", articleService, args{0}, ArticleReponse{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleService{
				articleRepository: tt.fields.articleRepository,
			}
			got, got1 := a.ShowArticle(context.Background(), tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleService.ShowArticle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ArticleService.ShowArticle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArticleService_ShowAllArticle(t *testing.T) {
	articleService := NewArticleService(&mock_ArticleRepository{})
	var articleallresponse ArticleAllReponse
	articleallresponse.Articles = append(articleallresponse.Articles, ArticleReponse{1, "Rahul Shewale", "ABCD", GetDummyDate(), "ABSCD"})
	tests := []struct {
		name   string
		fields *ArticleService
		want   ArticleAllReponse
		want1  bool
	}{
		// TODO: Add test cases.
		{"ShowAll Article Service", articleService, articleallresponse, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleService{
				articleRepository: tt.fields.articleRepository,
			}
			got, got1 := a.ShowAllArticle(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleService.ShowAllArticle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ArticleService.ShowAllArticle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
