package repository

import (
	"context"
	"final-project/entity"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type PhotoRepoInterface interface {
	CreatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error)
	GetPhotos(ctx context.Context) ([]*entity.Photo, error)
	GetPhotoById(ctx context.Context, id string) (*entity.Photo, error)
	UpdatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error)
	DeletePhoto(ctx context.Context, id string) error
}

type PhotoRepo struct {
	pgPool *pgxpool.Pool
}

func NewPhotoRepo(pgPool *pgxpool.Pool) PhotoRepoInterface {
	return &PhotoRepo{pgPool: pgPool}
}

func (p PhotoRepo) CreatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error) {
	//TODO implement me
	sql := "insert into photos(title, caption, photo_url, user_id, created_at) values($1, $2, $3, $4, now())"

	_, err := p.pgPool.Exec(ctx, sql, photo.Title, photo.Caption, photo.PhotoUrl, id)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &photo, nil

}

func (p PhotoRepo) GetPhotos(ctx context.Context) ([]*entity.Photo, error) {
	row, err := p.pgPool.Query(ctx,
		`
			Select pt.id, pt.title, pt.caption,pt.photo_url, pt.user_id, us.email, us.username
			from photos pt join users us 
			on pt.user_id = us.id`)

	if err != nil {
		fmt.Println("queryRowError", err)
		return nil, err
	}

	defer row.Close()

	var photos []*entity.Photo

	for row.Next() {
		var photo entity.Photo
		err := row.Scan(
			&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID, &photo.User.Email, &photo.User.Username,
		)

		if err != nil {
			fmt.Println("errGetAll", err.Error())
			return nil, err
		}

		photos = append(photos, &photo)
	}

	return photos, nil
}

func (p PhotoRepo) GetPhotoById(ctx context.Context, id string) (*entity.Photo, error) {
	//TODO implement me
	var photo entity.Photo

	sql := "select id, title, caption, photo_url, user_id from photos where id=$1"
	row := p.pgPool.QueryRow(ctx, sql, id)

	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID)

	if err != nil {
		fmt.Println("errById", err.Error())
		return nil, err
	}

	return &photo, nil

}

func (p PhotoRepo) UpdatePhoto(ctx context.Context, id string, photo entity.Photo) (*entity.Photo, error) {
	//TODO implement me
	sql := "update photos set title=$1, caption=$2, photo_url=$3, updated_at=now() where id=$4"

	_, err := p.pgPool.Exec(ctx, sql, photo.Title, photo.Caption, photo.PhotoUrl, id)

	if err != nil {
		fmt.Println("queryRowErrUpdatePho", err.Error())
		return nil, err
	}

	return &photo, nil
}

func (p PhotoRepo) DeletePhoto(ctx context.Context, id string) error {
	//TODO implement me
	sql := "delete from photos where id=$1"

	_, err := p.pgPool.Exec(ctx, sql, id)

	if err != nil {
		fmt.Println("queryRowErrDeletePho", err.Error())
		return err
	}

	return nil
}
