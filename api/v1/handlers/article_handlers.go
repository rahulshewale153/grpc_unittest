package handlers

import (
	"encoding/json"
	"fmt"
	"grpc_unittest/api"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/configs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	models "grpc_unittest/api/models"
	. "grpc_unittest/api/v1"

	"github.com/gorilla/mux"
)

//ArticleRequest Type recieves the Article data from frontend in the specified formats
type ArticleRequest struct {
	Nickname            string `json:"nickname"`
	Title               string `json:"title"`
	ArticleCreationDate string `json:"articlecreationdate"`
	Content             string `json:"content"`
}

type HttpArticleHandler struct {
	ArticleService services.ArticleServices
}

const (
	DATEFORMAT string = "2006-01-02"
)

func NewArticleHttpHandler(articleService services.ArticleServices, router api.Route) *HttpArticleHandler {
	handler := &HttpArticleHandler{ArticleService: articleService}
	router.Router.HandleFunc("/article", handler.CreateArticle).Methods("POST")
	router.Router.HandleFunc("/article/{id}", handler.ShowArticle).Methods("GET")
	router.Router.HandleFunc("/articles", handler.ShowAllArticle).Methods("GET")
	router.Router.HandleFunc("/", Index).Methods("GET")
	return handler
}

//Index Func
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "server running!\n")
}

//CreateArticle Method
func (httpArticle HttpArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "read body", err)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid data format")
		return
	}

	articleData := ArticleRequest{}
	if err1 := json.Unmarshal(body, &articleData); err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "unmarshal json", err1)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid data format")
		return
	}
	defer r.Body.Close()

	datedb, err := time.Parse(DATEFORMAT, articleData.ArticleCreationDate)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "chnage  date format ", err)
		WriteErrorResponse(w, http.StatusBadRequest, "invalid date format")
		return
	}
	var articledb models.Article
	articledb.Nickname = articleData.Nickname
	articledb.Title = articleData.Title
	articledb.ArticleCreationDate = datedb
	articledb.Content = articleData.Content
	flag := httpArticle.ArticleService.AddArticle(r.Context(), &articledb)
	if !flag {
		WriteErrorResponse(w, http.StatusBadRequest, "unable to create article")
		return
	}
	WriteOKResponse(w, "article created successfully!")

}

//ShowArticle Method shows only specific Article with referance to article_id
func (httpArticle HttpArticleHandler) ShowArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleid := params["id"]
	articleid_number, err := strconv.ParseInt(articleid, 10, 64)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "parse article id", err)
		WriteErrorResponse(w, http.StatusOK, "invaild Article Id"+articleid)
		return
	}
	configs.Ld.Logger(r.Context(), configs.ERROR, "parse article id", err)
	articledb, flag := httpArticle.ArticleService.ShowArticle(r.Context(), articleid_number)
	if !flag {
		WriteErrorResponse(w, http.StatusOK, "article not found")
		return
	}
	WriteOKResponse(w, articledb)
}

//ShowAllArticle Method
func (httpArticle HttpArticleHandler) ShowAllArticle(w http.ResponseWriter, r *http.Request) {
	articledb, flag := httpArticle.ArticleService.ShowAllArticle(r.Context())
	if !flag {
		WriteErrorResponse(w, http.StatusOK, "article not found")
		return
	}
	WriteOKResponse(w, articledb)
}
