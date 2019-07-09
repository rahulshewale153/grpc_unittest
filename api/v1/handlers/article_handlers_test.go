package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grpc_unittest/api"
	models "grpc_unittest/api/models"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/configs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

type mock_ArticleService struct {
}

func init() {
	configs.Config.Read("testing")
	fmt.Println(configs.Config.Handler)
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

func TestIndex(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	NewArticleHttpHandler(&mock_ArticleService{}, router)
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

func TestHttpArticleHandler_CreateArticle(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewArticleHttpHandler(&mock_ArticleService{}, router)

	type args struct {
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

	tests := []struct {
		name       string
		args       interface{}
		statuscode int
		//want string
	}{
		// TODO: Add test cases.
		{"Get Article ", args{"RAhul Shinde", "80 mindset 20 Skill", "2019-02-04", "Good Article"}, 200},
		{"Get Article ", args{"", "80 mindset 20 Skill", "2019-02-04", "Good Article"}, 400},
		{"Get Article ", argswrong{"80 mindset 20 Skill", "80 mindset 20 Skill", "80 mindset 20 Skill", 80}, 400},
		{"Get Article ", "", 400},
		{"Get Article ", args{"RAhul Shinde", "80 mindset 20 Skill", "209-02-04", "Good Article"}, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.args)
			req, err := http.NewRequest("POST", "/article/", bytes.NewBuffer(jsonValue))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/article/", httphandler.CreateArticle)
			router.ServeHTTP(rr, req)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			if status := rr.Code; status != tt.statuscode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}

func TestHttpArticleHandler_ShowArticle(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewArticleHttpHandler(&mock_ArticleService{}, router)

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args

		//want string
	}{
		// TODO: Add test cases.
		{"Get Article ", args{"1"}},
		{"Get Article ", args{"0"}},
		{"Get Article ", args{"s"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/article/%s", tt.args.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/article/{id}", httphandler.ShowArticle)
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
func TestHttpArticleHandler_ShowAllArticle(t *testing.T) {
	router := api.Route{}
	router.Router = mux.NewRouter()
	httphandler := NewArticleHttpHandler(&mock_ArticleService{}, router)
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httphandler.ShowAllArticle)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
