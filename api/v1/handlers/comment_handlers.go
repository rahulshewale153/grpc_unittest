package handlers

import (
	"encoding/json"
	"grpc_unittest/api"
	models "grpc_unittest/api/models"
	. "grpc_unittest/api/v1"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/configs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//CommentRequest Type  recieves the Comment data from frontend in the specified formats
type CommentRequest struct {
	Nickname            string `json:"nickname"`
	ArticleID           uint   `json:"articleid"`
	Content             string `json:"content"`
	CommentCreationDate string `json:"commentcreationdate"`
}

type HttpCommentHandler struct {
	CommentService services.CommentInterface
}

func NewCommentHttpHandler(commentService services.CommentInterface, router api.Route) *HttpCommentHandler {
	handler := &HttpCommentHandler{CommentService: commentService}
	router.Router.HandleFunc("/comment", handler.CreateComment).Methods("POST")
	router.Router.HandleFunc("/", Index).Methods("GET")
	router.Router.HandleFunc("/comment/{id}", handler.ShowComment).Methods("GET")
	router.Router.HandleFunc("/comments", handler.ShowAllComment).Methods("GET")
	router.Router.HandleFunc("/article/{id}/comments", handler.ShowAllArticleComment).Methods("GET")
	return handler
}

//CreateComment Method
func (c HttpCommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "read comment body", err)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid data")
		return
	}
	commentData := CommentRequest{}
	if err2 := json.Unmarshal(body, &commentData); err2 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "unmarshal comment body", err2)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid data")
		return
	}
	defer r.Body.Close()

	datedb, err := time.Parse(DATEFORMAT, commentData.CommentCreationDate)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "convert date format", err)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid date format")
		return
	}
	commentdb := models.Comment{}
	commentdb.Nickname = commentData.Nickname
	commentdb.ArticleID = commentData.ArticleID
	commentdb.Content = commentData.Content
	commentdb.CommentCreationDate = datedb
	flag := c.CommentService.AddComment(r.Context(), &commentdb)
	if !flag {
		WriteErrorResponse(w, http.StatusBadRequest, "unable to create comment ")
		return
	}
	WriteOKResponse(w, "commented on article successfully!")

}

//ShowComment Method shows only specific Comment with referance to article_id
func (c HttpCommentHandler) ShowComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentid := params["id"]
	commentid_number, err := strconv.ParseInt(commentid, 10, 64)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "parse comment id into int", err)
		WriteErrorResponse(w, http.StatusOK, "invaild comment id")
		return
	}
	commentdb, flag := c.CommentService.ShowComment(r.Context(), commentid_number)
	if !flag {
		WriteErrorResponse(w, http.StatusOK, "comment not found")
		return
	}
	WriteOKResponse(w, commentdb)
}

//ShowAllComment Method
func (c HttpCommentHandler) ShowAllComment(w http.ResponseWriter, r *http.Request) {
	commentdb, flag := c.CommentService.ShowAllComment(r.Context())
	if !flag {
		WriteErrorResponse(w, http.StatusOK, "comment not found")
		return
	}
	WriteOKResponse(w, commentdb)
}

//ShowComment Method shows only specific Comment with referance to article_id
func (c HttpCommentHandler) ShowAllArticleComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleid := params["id"]
	articleid_number, err := strconv.ParseInt(articleid, 10, 64)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "parse article id into int", err)
		WriteErrorResponse(w, http.StatusOK, "invaild comment id")
		return
	}
	commentdb, flag := c.CommentService.ShowArticleComment(r.Context(), articleid_number)
	if !flag {
		WriteErrorResponse(w, http.StatusOK, "article comment not found")
		return
	}
	WriteOKResponse(w, commentdb)
}
