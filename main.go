package main

import (
	"grpc_unittest/api"
	"grpc_unittest/api/v1/grpc_handlers"
	"grpc_unittest/api/v1/handlers"
	"grpc_unittest/api/v1/middleware"
	"grpc_unittest/api/v1/repository"
	"grpc_unittest/api/v1/services"
	"grpc_unittest/configs"
	"grpc_unittest/database/connection"
	"grpc_unittest/grpc/article_grpc"
	"grpc_unittest/grpc/comment_grpc"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

//initialize the code
func init() {
	//Set Env variable for config file
	configs.Config.Read("production")
}

const (
	PRODUCTION  int8 = 1
	DEVELOPMENT int8 = 0
)
const (
	GRPC string = "grpc"
	API  string = "api"
)

func main() {
	log.Println("Main Running!")
	if configs.Config.Handler == API {
		router := api.Route{}
		router.Router = mux.NewRouter()
		rm := middleware.RequestMiddleware{}
		router.Router.Use(rm.RequestIdGenerator)
		//Create Database Connection
		connectionService := connection.NewDatabaseConnection()
		//Article
		articleRepository := repository.NewArticleRepository(connectionService)
		articelService := services.NewArticleService(articleRepository)
		handlers.NewArticleHttpHandler(articelService, router)

		//Comemnt
		commentRepository := repository.NewCommentRepository(connectionService)
		commentService := services.NewCommentService(commentRepository)
		handlers.NewCommentHttpHandler(commentService, router)

		log.Fatal(http.ListenAndServe(":"+configs.Config.Port, router.Router))
	}
	if configs.Config.Handler == GRPC {
		//Create Database Connection
		connectionService := connection.NewDatabaseConnection()
		//Article
		articleRepository := repository.NewArticleRepository(connectionService)
		articelService := services.NewArticleService(articleRepository)
		articlesServer := grpc_handlers.NewArticleHttpHandler(articelService)

		//Comemnt
		commentRepository := repository.NewCommentRepository(connectionService)
		commentService := services.NewCommentService(commentRepository)
		commentServer := grpc_handlers.NewCommentHttpHandler(commentService)
		srv := grpc.NewServer()
		article_grpc.RegisterArticlesServer(srv, articlesServer)
		comment_grpc.RegisterCommentsServer(srv, commentServer)

		l, err := net.Listen("tcp", ":8888")
		if err != nil {
			log.Fatalf("Could not Listen to : 8888 %v", err)
		}
		log.Fatal(srv.Serve(l))
	}
}
