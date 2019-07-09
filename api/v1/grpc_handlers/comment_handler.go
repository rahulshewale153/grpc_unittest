package grpc_handlers

import (
	"context"
	"grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/grpc/comment_grpc"
	"net/http"
	"time"
)

type CommentServer struct {
	CommentService services.CommentInterface
}

func NewCommentHttpHandler(commentService services.CommentInterface) *CommentServer {
	return &CommentServer{CommentService: commentService}

}
func (cs CommentServer) Commentlist(ctx context.Context, void *comment_grpc.Void) (*comment_grpc.CommentResponse, error) {
	var commentlist comment_grpc.CommentList
	var commentresponse comment_grpc.CommentResponse
	commentdb, dataflag := cs.CommentService.ShowAllComment(ctx)
	if !dataflag {
		commentresponse.Status = http.StatusOK
		commentresponse.Success = true
		commentresponse.Message = "data not present"
		return &commentresponse, nil
	}
	for _, commentdb := range commentdb.Comments {
		var comment comment_grpc.Comment
		comment.Id = int64(commentdb.ID)
		comment.Nickname = commentdb.Nickname
		comment.Articleid = int64(commentdb.ArticleID)
		comment.Commentcreationdate = commentdb.CommentCreationDate.String()
		comment.Content = commentdb.Content
		commentlist.Comments = append(commentlist.Comments, &comment)
	}
	commentresponse.Status = http.StatusOK
	commentresponse.Success = true
	commentresponse.Message = "data present"
	commentresponse.Comments = commentlist.Comments
	return &commentresponse, nil

}

func (cs CommentServer) Addcomment(ctx context.Context, commentdata *comment_grpc.AddComment) (*comment_grpc.CommentResponse, error) {
	var commentresponse comment_grpc.CommentResponse
	datedb, err := time.Parse(DATEFORMAT, commentdata.Commentcreationdate)
	if err != nil {
		commentresponse.Status = http.StatusBadRequest
		commentresponse.Success = false
		commentresponse.Message = "data not present"
		return &commentresponse, nil
	}
	var commnetInsert models.Comment
	commnetInsert.ArticleID = uint(commentdata.Articleid)
	commnetInsert.Nickname = commentdata.Nickname
	commnetInsert.CommentCreationDate = datedb
	commnetInsert.Content = commentdata.Content
	flag := cs.CommentService.AddComment(ctx, &commnetInsert)
	if !flag {
		commentresponse.Status = http.StatusOK
		commentresponse.Success = false
		commentresponse.Message = "Something is wrong, record not save"
		return &commentresponse, nil
	}
	commentresponse.Status = http.StatusOK
	commentresponse.Success = true
	commentresponse.Message = "Record save successfully"
	return &commentresponse, nil

}

func (cs CommentServer) Searchcomment(ctx context.Context, commentdata *comment_grpc.SearchComment) (*comment_grpc.CommentResponse, error) {
	var commentlist comment_grpc.CommentList
	var commentresponse comment_grpc.CommentResponse
	commentdb, dataflag := cs.CommentService.ShowComment(ctx, commentdata.Id)
	if !dataflag {
		commentresponse.Status = http.StatusOK
		commentresponse.Success = true
		commentresponse.Message = "data not present"
		return &commentresponse, nil
	}
	var comment comment_grpc.Comment
	comment.Id = int64(commentdb.ID)
	comment.Nickname = commentdb.Nickname
	comment.Articleid = int64(commentdb.ArticleID)
	comment.Commentcreationdate = commentdb.CommentCreationDate.String()
	comment.Content = commentdb.Content
	commentlist.Comments = append(commentlist.Comments, &comment)

	commentresponse.Status = http.StatusOK
	commentresponse.Success = true
	commentresponse.Message = "data present"
	commentresponse.Comments = commentlist.Comments
	return &commentresponse, nil

}
func (cs CommentServer) Searcharticlecomment(ctx context.Context, commentdata *comment_grpc.SearchArticleComment) (*comment_grpc.CommentResponse, error) {
	var commentlist comment_grpc.CommentList
	var commentresponse comment_grpc.CommentResponse
	commentdb, dataflag := cs.CommentService.ShowArticleComment(ctx, commentdata.Id)
	if !dataflag {
		commentresponse.Status = http.StatusOK
		commentresponse.Success = true
		commentresponse.Message = "data not present"
		return &commentresponse, nil
	}
	for _, commentdb := range commentdb.Comments {
		var comment comment_grpc.Comment
		comment.Id = int64(commentdb.ID)
		comment.Nickname = commentdb.Nickname
		comment.Articleid = int64(commentdb.ArticleID)
		comment.Commentcreationdate = commentdb.CommentCreationDate.String()
		comment.Content = commentdb.Content
		commentlist.Comments = append(commentlist.Comments, &comment)
	}
	commentresponse.Status = http.StatusOK
	commentresponse.Success = true
	commentresponse.Message = "data present"
	commentresponse.Comments = commentlist.Comments
	return &commentresponse, nil
}
