package repository

import (
	"context"
	"database/sql"
	"final-project/entity"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, id string, comment entity.Comment) (*entity.Comment, error)
	GetAllComment(ctx context.Context) ([]*entity.Comment, error)
	GetCommentById(ctx context.Context, id string) (*entity.Comment, error)
	UpdateComment(ctx context.Context, id string, comment entity.Comment) (*entity.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type CommentRepository struct {
	pgPool *pgxpool.Pool
}

func NewCommentRepository(pgPool *pgxpool.Pool) CommentRepositoryInterface {
	return &CommentRepository{
		pgPool: pgPool,
	}
}

func (c CommentRepository) CreateComment(ctx context.Context, id string, comment entity.Comment) (*entity.Comment, error) {
	//TODO implement me
	sqlQuery := "insert into comments(message, photo_id, user_id, created_at) values($1, $2, $3, now())"

	_, err := c.pgPool.Exec(ctx, sqlQuery, comment.Message, comment.PhotoID, id)

	if err != nil {
		fmt.Println("query create error", err)
		return nil, err
	}

	return &comment, nil
}

func (c CommentRepository) GetAllComment(ctx context.Context) ([]*entity.Comment, error) {
	//TODO implement me
	row, errRow := c.pgPool.Query(ctx, `
				select cm.id, cm.user_id, cm.photo_id, cm.message, cm.created_at, cm.updated_at, 
				p.id, p.title, p.caption, p.photo_url, p.user_id, us.id, us.username, us.email
				from comments cm
				left join photos p on cm.photo_id = p.id
				left join users us on cm.user_id = us.id`)

	if errRow != nil {
		fmt.Println("query get all error", errRow.Error())
		return nil, errRow
	}

	defer row.Close()

	var comments []*entity.Comment

	for row.Next() {
		var comment entity.Comment
		var timeAt sql.NullTime

		err := row.Scan(
			&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt,
			&timeAt, &comment.Photo.ID, &comment.Photo.Title, &comment.Photo.Caption,
			&comment.Photo.PhotoUrl, &comment.Photo.UserID, &comment.User.ID, &comment.User.Username,
			&comment.User.Email,
		)

		if err != nil {
			fmt.Println("err scan get all", err.Error())
			return nil, err
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c CommentRepository) GetCommentById(ctx context.Context, id string) (*entity.Comment, error) {
	//TODO implement me
	var comment *entity.Comment

	sqlQuery := `select id, message, photo_id, user_id from photos where id=$1`
	row, errRow := c.pgPool.Query(ctx, sqlQuery, id)

	if errRow != nil {
		fmt.Println("query by id error", errRow)
		return nil, errRow
	}

	defer row.Close()

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c CommentRepository) UpdateComment(ctx context.Context, id string, comment entity.Comment) (*entity.Comment, error) {
	//TODO implement me
	sqlQuery := `update comments set message=$1, updated_at=now() where id=$2`

	_, errRow := c.pgPool.Exec(ctx, sqlQuery, comment.Message, id)

	if errRow != nil {
		fmt.Println("query update error", errRow)
		return nil, errRow
	}

	return &comment, nil
}

func (c CommentRepository) DeleteComment(ctx context.Context, id string) error {
	//TODO implement me
	sqlQuery := `delete from comments where id=$1`
	_, errRow := c.pgPool.Exec(ctx, sqlQuery, id)

	if errRow != nil {
		fmt.Println("query delete error", errRow)
		return errRow
	}

	return nil
}
