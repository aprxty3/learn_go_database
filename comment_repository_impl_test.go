package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_database/entity"
	"go_database/repository"
	"testing"
)

func TestInsertComment(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	comments := entity.Comments{
		Email:   "test@gmail.com",
		Comment: "test comment",
	}

	result, err := commentRepository.Insert(ctx, comments)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 22)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
