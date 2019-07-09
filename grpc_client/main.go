package main

import (
	"fmt"
	"grpc_unittest/grpc/article_grpc"
	"grpc_unittest/grpc/comment_grpc"
	"log"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to backend %v", err)
	}

	fmt.Println("----------------------Article LIST RECORD---------------------")
	client := article_grpc.NewArticlesClient(conn)
	err = ArticleList(context.Background(), client)

	fmt.Println("---------------------- Article ADD RECORD---------------------")
	err = AddArticle(context.Background(), client)

	fmt.Println("----------------------Article Search RECORD---------------------")
	err = SearchArticle(context.Background(), client)

	fmt.Println("----------------------Comment LIST RECORD---------------------")
	commentclient := comment_grpc.NewCommentsClient(conn)
	err = Commentlist(context.Background(), commentclient)

	fmt.Println("---------------------- Comment ADD RECORD---------------------")
	err = AddComment(context.Background(), commentclient)

	fmt.Println("----------------------Comment Search Using CommentID---------------------")
	err = SearchComment(context.Background(), commentclient)

	fmt.Println("----------------------Comment Search Using Article ID---------------------")
	err = SearchArticleComment(context.Background(), commentclient)
}

func ArticleList(ctx context.Context, client article_grpc.ArticlesClient) error {
	l, err := client.Articlelist(ctx, &article_grpc.Void{})
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	for _, article := range l.Articles {
		fmt.Println(article)
	}
	return nil
}

func AddArticle(ctx context.Context, client article_grpc.ArticlesClient) error {

	var articledata article_grpc.AddArticle
	articledata.Nickname = "Pooname Malve"
	articledata.Title = "Cartoon"
	articledata.Articlecreationdate = "2019-02-04"
	articledata.Content = "Pockymon"
	response, err := client.Addarticle(ctx, &articledata)
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	fmt.Println(response)
	return nil
}

func SearchArticle(ctx context.Context, client article_grpc.ArticlesClient) error {

	var articledata article_grpc.SearchArticle
	articledata.Id = 5
	response, err := client.Searcharticle(ctx, &articledata)
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	fmt.Println(response)
	for _, article := range response.Articles {
		fmt.Println(article)
	}
	return nil
}

/*                           Comment Action                   */
func Commentlist(ctx context.Context, client comment_grpc.CommentsClient) error {
	response, err := client.Commentlist(ctx, &comment_grpc.Void{})
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	for _, comment := range response.Comments {
		fmt.Println(comment)
	}
	return nil
}

func AddComment(ctx context.Context, client comment_grpc.CommentsClient) error {

	var commentdata comment_grpc.AddComment
	commentdata.Nickname = "Pooname Malve"
	commentdata.Articleid = 5
	commentdata.Commentcreationdate = "2019-02-04"
	commentdata.Content = "I Like Pockymon"
	response, err := client.Addcomment(ctx, &commentdata)
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	fmt.Println(response)
	return nil
}

func SearchComment(ctx context.Context, client comment_grpc.CommentsClient) error {

	var commentdata comment_grpc.SearchComment
	commentdata.Id = 5
	response, err := client.Searchcomment(ctx, &commentdata)
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	fmt.Println(response)
	for _, comment := range response.Comments {
		fmt.Println(comment)
	}
	return nil
}

func SearchArticleComment(ctx context.Context, client comment_grpc.CommentsClient) error {

	var commentdata comment_grpc.SearchArticleComment
	commentdata.Id = 5
	response, err := client.Searcharticlecomment(ctx, &commentdata)
	if err != nil {
		return fmt.Errorf("Could fetch task %v", err)
	}
	fmt.Println(response)
	for _, comment := range response.Comments {
		fmt.Println(comment)
	}
	return nil
}
