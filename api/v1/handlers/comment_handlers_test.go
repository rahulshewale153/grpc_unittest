package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grpc_unittest/api"
	models "grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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

func TestNewCommentHttpHandler(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	NewCommentHttpHandler(&mock_CommentService{}, router)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHttpCommentHandler_CreateComment(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewCommentHttpHandler(&mock_CommentService{}, router)

	type args struct {
		Nickname            string
		ArticleID           int
		CommentCreationDate string
		Content             string
	}
	type argswrong struct {
		Nickname            string
		ArticleID           string
		CommentCreationDate string
		Content             int
	}
	tests := []struct {
		name       string
		args       interface{}
		statuscode int
	}{
		// TODO: Add test cases.
		{"Get comment 1", args{"RAhul Shinde", 1, "2019-02-04", "Good Article"}, 200},
		{"Get comment 2", args{"", 1, "2019-02-04", "Good Article"}, 400},
		{"Get comment 3", argswrong{"80 mindset 20 Skill", "80 mindset 20 Skill", "80 mindset 20 Skill", 80}, 400},
		{"Get comment 4", "", 400},
		{"Get comment 5", args{"RAhul Shinde", 1, "209-02-04", "Good Article"}, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.args)
			req, err := http.NewRequest("POST", "/comment/", bytes.NewBuffer(jsonValue))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/comment/", httphandler.CreateComment)
			router.ServeHTTP(rr, req)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			if status := rr.Code; status != tt.statuscode {
				t.Errorf("handler returned wrong status code: got %v want %v data",
					status, http.StatusOK)
			}
		})
	}
}

func TestHttpCommentHandler_ShowComment(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewCommentHttpHandler(&mock_CommentService{}, router)

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args

		//want string
	}{
		// TODO: Add test cases.
		{"Get Comment ", args{"1"}},
		{"Get Comment ", args{"0"}},
		{"Get Comment ", args{"s"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/comment/%s", tt.args.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/comment/{id}", httphandler.ShowComment)
			router.ServeHTTP(rr, req)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

func TestHttpCommentHandler_ShowAllComment(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewCommentHttpHandler(&mock_CommentService{}, router)
	req, err := http.NewRequest("GET", "/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httphandler.ShowAllComment)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHttpCommentHandler_ShowAllArticleComment(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewCommentHttpHandler(&mock_CommentService{}, router)

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args

		//want string
	}{
		// TODO: Add test cases.
		{"Get Comment ", args{"1"}},
		{"Get Comment ", args{"0"}},
		{"Get Comment ", args{"s"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/article/%s/comments", tt.args.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/article/{id}/comments", httphandler.ShowAllArticleComment)
			router.ServeHTTP(rr, req)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}
