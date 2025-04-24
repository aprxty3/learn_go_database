package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentsRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comments entity.Comments) (entity.Comments, error) {
	script := "insert into comment (email, comment) values (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comments.Email, comments.Comment)
	if err != nil {
		return comments, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comments, err
	}
	comments.Id = int32(id)
	return comments, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	script := "select id, email, comment from comment where id = ? limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	comments := entity.Comments{}
	if err != nil {
		return comments, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&comments.Id, &comments.Email, &comments.Comment)
		if err != nil {
			return comments, err
		}

		return comments, nil
	} else {
		return comments, errors.New("id" + strconv.Itoa(int(id)) + " not found")
	}

}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	script := "select id, email, comment from comment"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var comments []entity.Comments

	for rows.Next() {
		comment := entity.Comments{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
