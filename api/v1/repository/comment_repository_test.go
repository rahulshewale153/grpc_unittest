package repository

import (
	"context"
	"grpc_unittest/api/models"
	"testing"
	"time"
)

func TestCommentRepository_Save(t *testing.T) {
	connectionService := NewDatabaseConnection()
	commentrepository := NewCommentRepository(connectionService)
	datedb, err := time.Parse("2006-01-02", "2019-02-02")
	if err != nil {
	}
	type args struct {
		comment *models.Comment
	}
	var record1 models.Comment
	record1.Nickname = "Rahul Shewale"
	record1.ArticleID = 1
	record1.CommentCreationDate = datedb
	record1.Content = "ABSCD"

	var record2 models.Comment
	record2.Nickname = "Rahul Shewale"
	record1.ArticleID = 1
	record2.Content = "ABSCD"

	tests := []struct {
		name   string
		fields *CommentRepository
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"Save record", commentrepository, args{&record1}, true},
		{"Save record", commentrepository, args{&record2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentRepository{
				connectionInterface: tt.fields.connectionInterface,
			}
			if got := a.Save(context.Background(), tt.args.comment); got != tt.want {
				t.Errorf("CommentRepository.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentRepository_Show(t *testing.T) {
	connectionService := NewDatabaseConnection()
	commentrepository := NewCommentRepository(connectionService)

	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *CommentRepository
		args   args
		//	want   models.Comment
		want1 bool
	}{
		// TODO: Add test cases.
		{"Show record", commentrepository, args{5}, true},
		{"Show record", commentrepository, args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentRepository{
				connectionInterface: tt.fields.connectionInterface,
			}
			_, got1 := a.Show(context.Background(), tt.args.id)
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentRepository.Show() got = %v, want %v", got, tt.want)
			} */
			if got1 != tt.want1 {
				t.Errorf("CommentRepository.Show() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCommentRepository_ShowAll(t *testing.T) {
	connectionService := NewDatabaseConnection()
	commentrepository := NewCommentRepository(connectionService)

	tests := []struct {
		name   string
		fields *CommentRepository
		//want   []models.Comment
		want1 bool
	}{
		// TODO: Add test cases.
		{"Show All", commentrepository, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentRepository{
				connectionInterface: tt.fields.connectionInterface,
			}
			_, got1 := a.ShowAll(context.Background())
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentRepository.ShowAll() got = %v, want %v", got, tt.want)
			} */
			if got1 != tt.want1 {
				t.Errorf("CommentRepository.ShowAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCommentRepository_ShowArticleComment(t *testing.T) {
	connectionService := NewDatabaseConnection()
	commentrepository := NewCommentRepository(connectionService)

	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *CommentRepository
		args   args
		//want   []models.Comment
		want1 bool
	}{
		// TODO: Add test cases.
		{"Show record", commentrepository, args{5}, true},
		{"Show record", commentrepository, args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentRepository{
				connectionInterface: tt.fields.connectionInterface,
			}
			_, got1 := a.ShowArticleComment(context.Background(), tt.args.id)
			/* 	if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentRepository.ShowArticleComment() got = %v, want %v", got, tt.want)
			} */
			if got1 != tt.want1 {
				t.Errorf("CommentRepository.ShowArticleComment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
