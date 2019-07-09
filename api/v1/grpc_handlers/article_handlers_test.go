package grpc_handlers

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/grpc/article_grpc"
	"testing"
	"time"
)

type mock_ArticleService struct {
}

func GetDummyDate() time.Time {
	datedb, err := time.Parse("2006-01-02", "2019-02-02")
	if err != nil {
	}
	return datedb
}
func (a mock_ArticleService) AddArticle(ctx context.Context, ar *models.Article) bool {
	if ar.Nickname == "" {
		return false
	}

	return true
}
func (a mock_ArticleService) ShowArticle(ctx context.Context, id int64) (services.ArticleReponse, bool) {
	record1 := services.ArticleReponse{}
	if id == 0 {
		return record1, false
	}
	record1.ID = 1
	record1.Nickname = "Rahul Shewale"
	record1.Title = "ABCD"
	record1.ArticleCreationDate = GetDummyDate()
	record1.Content = "ABSCD"
	return record1, true

}
func (a mock_ArticleService) ShowAllArticle(ctx context.Context) (services.ArticleAllReponse, bool) {
	var articleAllResponse services.ArticleAllReponse

	article := services.ArticleReponse{}
	article.ID = 1
	article.Nickname = "Rahul Shewale"
	article.Title = "ABCD"
	article.ArticleCreationDate = GetDummyDate()
	article.Content = "ABSCD"
	articleAllResponse.Articles = append(articleAllResponse.Articles, article)

	return articleAllResponse, true
}

func TestArticlesServer_Articlelist(t *testing.T) {
	grpchandler := NewArticleHttpHandler(&mock_ArticleService{})

	type args struct {
		ctx  context.Context
		void *article_grpc.Void
	}
	tests := []struct {
		name   string
		fields *ArticlesServer
		args   args
		//want    *article_grpc.ArticleResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Artcile List", grpchandler, args{context.Background(), &article_grpc.Void{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := ArticlesServer{
				ArticleService: tt.fields.ArticleService,
			}
			_, err := as.Articlelist(tt.args.ctx, tt.args.void)
			if err != tt.wantErr {
				t.Errorf("ArticlesServer.Articlelist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesServer.Articlelist() = %v, want %v", got, tt.want)
			} */
		})
	}
}

func TestArticlesServer_Addarticle(t *testing.T) {
	grpchandler := NewArticleHttpHandler(&mock_ArticleService{})
	type argsconrrect struct {
		Nickname            string
		Title               string
		Articlecreationdate string
		Content             string
	}
	type argswrong struct {
		Nickname            string
		Title               string
		Articlecreationdate string
		Content             int
	}
	type args struct {
		ctx         context.Context
		articleData *article_grpc.AddArticle
	}
	tests := []struct {
		name   string
		fields *ArticlesServer
		args   args
		//want    *article_grpc.ArticleResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Create Article ", grpchandler, args{context.Background(), &article_grpc.AddArticle{Nickname: "RAhul Shinde", Title: "80 mindset 20 Skill", Articlecreationdate: "2019-02-04", Content: "Good Article"}}, nil},
		{"Create Article ", grpchandler, args{context.Background(), &article_grpc.AddArticle{Nickname: "", Title: "80 mindset 20 Skill", Articlecreationdate: "2019-02-04", Content: "Good Article"}}, nil},
		{"Create Article ", grpchandler, args{context.Background(), &article_grpc.AddArticle{Nickname: "", Title: "80 mindset 20 Skill", Articlecreationdate: "Skill", Content: "Good Article"}}, nil},
		{"Create Article ", grpchandler, args{context.Background(), &article_grpc.AddArticle{}}, nil},
		{"Create Article ", grpchandler, args{context.Background(), &article_grpc.AddArticle{Nickname: "RAhul Shinde", Title: "80 mindset 20 Skill", Articlecreationdate: "209-02-04", Content: "Good Article"}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := ArticlesServer{
				ArticleService: tt.fields.ArticleService,
			}
			_, err := as.Addarticle(tt.args.ctx, tt.args.articleData)
			if err != tt.wantErr {
				t.Errorf("ArticlesServer.Addarticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesServer.Addarticle() = %v, want %v", got, tt.want)
			} */
		})
	}
}

func TestArticlesServer_Searcharticle(t *testing.T) {
	grpchandler := NewArticleHttpHandler(&mock_ArticleService{})
	type args struct {
		ctx            context.Context
		searcchArticle *article_grpc.SearchArticle
	}
	tests := []struct {
		name   string
		fields *ArticlesServer
		args   args
		//want    *article_grpc.ArticleResponse
		wantErr interface{}
	}{
		{"Get Article ", grpchandler, args{context.Background(), &article_grpc.SearchArticle{Id: 1}}, nil},
		{"Get Article ", grpchandler, args{context.Background(), &article_grpc.SearchArticle{Id: 0}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := ArticlesServer{
				ArticleService: tt.fields.ArticleService,
			}
			_, err := as.Searcharticle(tt.args.ctx, tt.args.searcchArticle)
			if err != tt.wantErr {
				t.Errorf("ArticlesServer.Searcharticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticlesServer.Searcharticle() = %v, want %v", got, tt.want)
			} */
		})
	}
}
