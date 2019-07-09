package repository

import (
	"context"
	"fmt"
	"grpc_unittest/api/models"
	"grpc_unittest/configs"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

type Connection struct {
}

func init() {
	configs.Config.Read("testing")
}
func (c Connection) DBConnect() *gorm.DB {
	DB, err := gorm.Open("mysql", "root:root@/article-get-post?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("Connection Sucessfull!")
	if err != nil {
		panic("Failed to connect database!")
	}
	return DB
}
func NewDatabaseConnection() *Connection {
	return &Connection{}
}

func TestArticleCommentRepository_Save(t *testing.T) {
	connectionService := NewDatabaseConnection()
	articlerepository := NewArticleRepository(connectionService)
	type args struct {
		ar *models.Article
	}
	datedb, err := time.Parse("2006-01-02", "2019-02-02")
	if err != nil {
	}
	var record1 models.Article
	record1.Nickname = "Rahul Shewale"
	record1.Title = "ABCD"
	record1.ArticleCreationDate = datedb
	record1.Content = "ABSCD"

	var record2 models.Article
	record2.Nickname = "Rahul Shewale"
	record2.Title = "ABCD"
	record2.Content = "ABSCD"

	tests := []struct {
		name   string
		fields *ArticleCommentRepository
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"Save record", articlerepository, args{&record1}, true},
		{"Save record", articlerepository, args{&record2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleCommentRepository{
				ConnectionService: tt.fields.ConnectionService,
			}
			if got := a.Save(context.Background(), tt.args.ar); got != tt.want {
				t.Errorf("ArticleCommentRepository.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleCommentRepository_Show(t *testing.T) {
	connectionService := NewDatabaseConnection()
	articlerepository := NewArticleRepository(connectionService)
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *ArticleCommentRepository
		args   args
		//want   models.Article
		want1 bool
	}{
		// TODO: Add test cases.
		{"Show record", articlerepository, args{5}, true},
		{"Show record", articlerepository, args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleCommentRepository{
				ConnectionService: tt.fields.ConnectionService,
			}
			_, got1 := a.Show(context.Background(), tt.args.id)
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleCommentRepository.Show() got = %v, want %v", got, tt.want)
			} */
			if got1 != tt.want1 {
				t.Errorf("ArticleCommentRepository.Show() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArticleCommentRepository_ShowAll(t *testing.T) {
	connectionService := NewDatabaseConnection()
	articlerepository := NewArticleRepository(connectionService)
	tests := []struct {
		name   string
		fields *ArticleCommentRepository
		//want   []models.Article
		want1 bool
	}{
		// TODO: Add test cases.
		{"Show All", articlerepository, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ArticleCommentRepository{
				ConnectionService: tt.fields.ConnectionService,
			}

			_, got1 := a.ShowAll(context.Background())
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleCommentRepository.ShowAll() got = %v, want %v", got, tt.want)
			} */
			if got1 != tt.want1 {
				t.Errorf("ArticleCommentRepository.ShowAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
