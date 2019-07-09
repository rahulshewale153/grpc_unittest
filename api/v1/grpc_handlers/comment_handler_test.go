package grpc_handlers

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/grpc/comment_grpc"
	"testing"
)

type mock_CommentService struct {
}

func (a mock_CommentService) AddComment(ctx context.Context, comment *models.Comment) bool {
	if comment.Nickname == "" {
		return false
	}

	return true
}
func (a mock_CommentService) ShowComment(ctx context.Context, id int64) (services.CommentResponse, bool) {
	comment := services.CommentResponse{}
	if id == 0 {
		return comment, false
	}
	comment.ID = 1
	comment.Nickname = "Rahul Shewale"
	comment.ArticleID = 2
	comment.CommentCreationDate = GetDummyDate()
	comment.Content = "ABSCD"
	return comment, true
}
func (a mock_CommentService) ShowAllComment(ctx context.Context) (services.CommentAllResponse, bool) {
	var commentAllResponse services.CommentAllResponse
	comment := services.CommentResponse{}
	comment.ID = 1
	comment.Nickname = "Rahul Shewale"
	comment.ArticleID = 2
	comment.CommentCreationDate = GetDummyDate()
	comment.Content = "ABSCD"
	commentAllResponse.Comments = append(commentAllResponse.Comments, comment)

	return commentAllResponse, true
}

func (a mock_CommentService) ShowArticleComment(ctx context.Context, id int64) (services.CommentAllResponse, bool) {
	var commentAllResponse services.CommentAllResponse
	if id == 0 {
		return commentAllResponse, false
	}
	comment := services.CommentResponse{}
	comment.ID = 1
	comment.Nickname = "Rahul Shewale"
	comment.ArticleID = 2
	comment.CommentCreationDate = GetDummyDate()
	comment.Content = "ABSCD"
	commentAllResponse.Comments = append(commentAllResponse.Comments, comment)

	return commentAllResponse, true
}
func TestCommentServer_Commentlist(t *testing.T) {
	httphandler := NewCommentHttpHandler(&mock_CommentService{})

	type args struct {
		ctx  context.Context
		void *comment_grpc.Void
	}
	tests := []struct {
		name   string
		fields *CommentServer
		args   args
		//want    *comment_grpc.CommentResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Get Comments", httphandler, args{context.Background(), &comment_grpc.Void{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CommentServer{
				CommentService: tt.fields.CommentService,
			}
			_, err := cs.Commentlist(tt.args.ctx, tt.args.void)
			if err != tt.wantErr {
				t.Errorf("CommentServer.Commentlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentServer.Commentlist() = %v, want %v", got, tt.want)
			} */
		})
	}
}

func TestCommentServer_Addcomment(t *testing.T) {
	httphandler := NewCommentHttpHandler(&mock_CommentService{})

	type args struct {
		ctx         context.Context
		commentdata *comment_grpc.AddComment
	}
	tests := []struct {
		name   string
		fields *CommentServer
		args   args
		//	want    *comment_grpc.CommentResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Add Comment", httphandler, args{context.Background(), &comment_grpc.AddComment{Nickname: "RAhul Shinde", Articleid: 1, Commentcreationdate: "2019-02-04", Content: "Good Article"}}, nil},
		{"Add Comment", httphandler, args{context.Background(), &comment_grpc.AddComment{Articleid: 1, Commentcreationdate: "2019-02-04", Content: "Good Article"}}, nil},
		{"Add Comment", httphandler, args{context.Background(), &comment_grpc.AddComment{Nickname: "RAhul Shinde", Articleid: 1, Commentcreationdate: "Article", Content: "Good Article"}}, nil},
		{"Add Comment", httphandler, args{context.Background(), &comment_grpc.AddComment{Nickname: "RAhul Shinde", Articleid: 1, Commentcreationdate: "201-01-12", Content: "Good Article"}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CommentServer{
				CommentService: tt.fields.CommentService,
			}
			_, err := cs.Addcomment(tt.args.ctx, tt.args.commentdata)
			if err != tt.wantErr {
				t.Errorf("CommentServer.Addcomment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentServer.Addcomment() = %v, want %v", got, tt.want)
			} */
		})
	}
}

func TestCommentServer_Searchcomment(t *testing.T) {
	httphandler := NewCommentHttpHandler(&mock_CommentService{})
	type args struct {
		ctx         context.Context
		commentdata *comment_grpc.SearchComment
	}
	tests := []struct {
		name   string
		fields *CommentServer
		args   args
		//want    *comment_grpc.CommentResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Get Comment", httphandler, args{context.Background(), &comment_grpc.SearchComment{Id: 1}}, nil},
		{"Get Comment", httphandler, args{context.Background(), &comment_grpc.SearchComment{Id: 0}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CommentServer{
				CommentService: tt.fields.CommentService,
			}
			_, err := cs.Searchcomment(tt.args.ctx, tt.args.commentdata)
			if err != tt.wantErr {
				t.Errorf("CommentServer.Searchcomment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentServer.Searchcomment() = %v, want %v", got, tt.want)
			} */
		})
	}
}

func TestCommentServer_Searcharticlecomment(t *testing.T) {
	httphandler := NewCommentHttpHandler(&mock_CommentService{})

	type args struct {
		ctx         context.Context
		commentdata *comment_grpc.SearchArticleComment
	}
	tests := []struct {
		name   string
		fields *CommentServer
		args   args
		//	want    *comment_grpc.CommentResponse
		wantErr interface{}
	}{
		// TODO: Add test cases.
		{"Get Comment using article id", httphandler, args{context.Background(), &comment_grpc.SearchArticleComment{Id: 1}}, nil},
		{"Get Comment using article id", httphandler, args{context.Background(), &comment_grpc.SearchArticleComment{Id: 0}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CommentServer{
				CommentService: tt.fields.CommentService,
			}
			_, err := cs.Searcharticlecomment(tt.args.ctx, tt.args.commentdata)
			if err != tt.wantErr {
				t.Errorf("CommentServer.Searcharticlecomment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/* if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentServer.Searcharticlecomment() = %v, want %v", got, tt.want)
			} */
		})
	}
}
