package repository

import (
	"context"
	"fmt"
	"go_mysql"
	"go_mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Id:      0,
		Email:   "repository@test.com",
		Comment: "Test repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 35)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
