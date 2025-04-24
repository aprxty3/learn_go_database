package repository

import (
	"context"
	"go_database/entity"
)

type CommentsRepository interface {
	Insert(ctx context.Context, comments entity.Comments) (entity.Comments, error)
	FindById(ctx context.Context, id int32) (entity.Comments, error)
	FindAll(ctx context.Context) ([]entity.Comments, error)
}
