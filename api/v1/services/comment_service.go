package services

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/repository"
	"time"
)

//CommentResponse Type sends the only required Comment data to frontend in the specified formats
type CommentResponse struct {
	ID                  uint      `json:"id"`
	Nickname            string    `json:"nickname"`
	ArticleID           uint      `json:"articleid"`
	Content             string    `json:"content"`
	CommentCreationDate time.Time `json:"commentcreationdate"`
}

//CommentResponse Type sends the only required Comment data to frontend in the specified formats
type CommentAllResponse struct {
	Comments []CommentResponse `json:"comments"`
}

//UserRepository to manage users persistence
type CommentService struct {
	commentRepository repository.CommentRepositoryInterface
}

func NewCommentService(commentRepos repository.CommentRepositoryInterface) *CommentService {
	return &CommentService{commentRepository: commentRepos}
}

func (a CommentService) AddComment(ctx context.Context, comment *models.Comment) bool {
	commentCreate := a.commentRepository.Save(ctx, comment)
	if !commentCreate {
		return false
	}
	return true
}
func (a CommentService) ShowComment(ctx context.Context, id int64) (CommentResponse, bool) {
	comment := CommentResponse{}
	commentdb, dataFlag := a.commentRepository.Show(ctx, id)
	if !dataFlag {
		return comment, false
	}
	comment.ID = commentdb.ID
	comment.Nickname = commentdb.Nickname
	comment.ArticleID = commentdb.ArticleID
	comment.Content = commentdb.Content
	comment.CommentCreationDate = commentdb.CommentCreationDate
	return comment, true
}
func (a CommentService) ShowAllComment(ctx context.Context) (CommentAllResponse, bool) {
	var commentAllResponse CommentAllResponse
	commentdb, dataFlag := a.commentRepository.ShowAll(ctx)
	if !dataFlag {
		return commentAllResponse, false
	}
	comment := CommentResponse{}
	for _, value := range commentdb {
		comment.ID = value.ID
		comment.Nickname = value.Nickname
		comment.ArticleID = value.ArticleID
		comment.Content = value.Content
		comment.CommentCreationDate = value.CommentCreationDate
		commentAllResponse.Comments = append(commentAllResponse.Comments, comment)
	}
	return commentAllResponse, true
}

func (a CommentService) ShowArticleComment(ctx context.Context, id int64) (CommentAllResponse, bool) {
	var commentAllResponse CommentAllResponse
	commentdb, dataFlag := a.commentRepository.ShowArticleComment(ctx, id)
	if !dataFlag {
		return commentAllResponse, false
	}
	comment := CommentResponse{}
	for _, value := range commentdb {
		comment.ID = value.ID
		comment.Nickname = value.Nickname
		comment.ArticleID = value.ArticleID
		comment.Content = value.Content
		comment.CommentCreationDate = value.CommentCreationDate
		commentAllResponse.Comments = append(commentAllResponse.Comments, comment)
	}
	return commentAllResponse, true
}
