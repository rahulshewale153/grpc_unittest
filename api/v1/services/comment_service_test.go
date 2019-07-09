package services

import (
	"context"
	"grpc_unittest/api/models"
	"reflect"
	"testing"
)

type mock_CommentRepository struct {
}

func (a mock_CommentRepository) Save(ctx context.Context, comment *models.Comment) bool {
	if comment.Nickname == "" {
		return false
	}
	return true
}

func (a mock_CommentRepository) Show(ctx context.Context, id int64) (models.Comment, bool) {
	commentdb := models.Comment{}
	if id == 0 {
		return commentdb, false
	}
	commentdb.Nickname = "Rahul Shewale"
	commentdb.ArticleID = 2
	commentdb.ID = 1
	commentdb.CommentCreationDate = GetDummyDate()
	commentdb.Content = "Nice Article"
	return commentdb, true
}

func (a mock_CommentRepository) ShowAll(ctx context.Context) ([]models.Comment, bool) {
	commentdball := []models.Comment{}

	commentdb := models.Comment{}
	commentdb.Nickname = "Rahul Shewale"
	commentdb.ArticleID = 2
	commentdb.ID = 1
	commentdb.CommentCreationDate = GetDummyDate()
	commentdb.Content = "Nice Article"
	commentdball = append(commentdball, commentdb)
	return commentdball, true
}

func (a mock_CommentRepository) ShowArticleComment(ctx context.Context, id int64) ([]models.Comment, bool) {
	commentdball := []models.Comment{}
	if id == 2 {
		commentdb := models.Comment{}
		commentdb.Nickname = "Rahul Shewale"
		commentdb.ArticleID = 2
		commentdb.ID = 1
		commentdb.CommentCreationDate = GetDummyDate()
		commentdb.Content = "Nice Article"
		commentdball = append(commentdball, commentdb)
		return commentdball, true
	}
	return commentdball, false
}

func TestCommentService_AddComment(t *testing.T) {
	commentService := NewCommentService(&mock_CommentRepository{})
	var record1 models.Comment
	record1.Nickname = "Rahul Shewale"
	record1.ArticleID = 1
	record1.CommentCreationDate = GetDummyDate()
	record1.Content = "ABSCD"

	var record2 models.Comment

	type args struct {
		comment *models.Comment
	}
	tests := []struct {
		name   string
		fields *CommentService
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"Add Comment 1", commentService, args{&record1}, true},
		{"Add Comment 2", commentService, args{&record2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				commentRepository: tt.fields.commentRepository,
			}
			if got := a.AddComment(context.Background(), tt.args.comment); got != tt.want {
				t.Errorf("CommentService.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentService_ShowComment(t *testing.T) {
	commentService := NewCommentService(&mock_CommentRepository{})
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *CommentService
		args   args
		want   CommentResponse
		want1  bool
	}{
		// TODO: Add test cases.
		{"Show Comment 1", commentService, args{1}, CommentResponse{1, "Rahul Shewale", 2, "Nice Article", GetDummyDate()}, true},
		{"Show Comment 2", commentService, args{0}, CommentResponse{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				commentRepository: tt.fields.commentRepository,
			}
			got, got1 := a.ShowComment(context.Background(), tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.ShowComment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CommentService.ShowComment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCommentService_ShowAllComment(t *testing.T) {
	commentService := NewCommentService(&mock_CommentRepository{})
	var commentallresponse CommentAllResponse
	commentallresponse.Comments = append(commentallresponse.Comments, CommentResponse{1, "Rahul Shewale", 2, "Nice Article", GetDummyDate()})

	tests := []struct {
		name   string
		fields *CommentService
		want   CommentAllResponse
		want1  bool
	}{
		// TODO: Add test cases.
		{"Show All Comment 1", commentService, commentallresponse, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				commentRepository: tt.fields.commentRepository,
			}
			got, got1 := a.ShowAllComment(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.ShowAllComment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CommentService.ShowAllComment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCommentService_ShowArticleComment(t *testing.T) {
	commentService := NewCommentService(&mock_CommentRepository{})
	var commentallresponse CommentAllResponse
	commentallresponse.Comments = append(commentallresponse.Comments, CommentResponse{1, "Rahul Shewale", 2, "Nice Article", GetDummyDate()})
	var commentallresponse1 CommentAllResponse
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields *CommentService
		args   args
		want   CommentAllResponse
		want1  bool
	}{
		// TODO: Add test cases.
		{"Show Article Comment 1", commentService, args{2}, commentallresponse, true},
		{"Show Article Comment 2", commentService, args{3}, commentallresponse1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				commentRepository: tt.fields.commentRepository,
			}
			got, got1 := a.ShowArticleComment(context.Background(), tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.ShowArticleComment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CommentService.ShowArticleComment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
