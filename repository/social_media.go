package repository

import (
	"context"
	"database/sql"
	"final-project/entity"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SocialMediaRepoInterface interface {
	CreateSocialMedia(ctx context.Context, id string, photo entity.SocialMedia) (*entity.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context) ([]*entity.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id string) (*entity.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, id string, photo entity.SocialMedia) (*entity.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id string) error
}

type SocialMediaRepo struct {
	pgPool *pgxpool.Pool
}

func NewSocialMediaRepo(pgPool *pgxpool.Pool) SocialMediaRepoInterface {
	return &SocialMediaRepo{pgPool: pgPool}
}

func (sp SocialMediaRepo) CreateSocialMedia(ctx context.Context, id string, socialMedia entity.SocialMedia) (*entity.SocialMedia, error) {
	//TODO implement me
	sqlQuery := "insert into social_medias(name, social_media_url, user_id, created_at) values($1, $2, $3, now())"

	_, err := sp.pgPool.Exec(ctx, sqlQuery, socialMedia.Name, socialMedia.SocialMediaUrl, id)

	if err != nil {
		fmt.Println("err create sm", err.Error())
		return nil, err
	}

	return &socialMedia, nil

}

func (sp SocialMediaRepo) GetAllSocialMedia(ctx context.Context) ([]*entity.SocialMedia, error) {
	//TODO implement me
	var socialMedias []*entity.SocialMedia

	row, err := sp.pgPool.Query(ctx,
		`
			 select sm.id, sm.name, sm.social_media_url, sm.user_id, sm.created_at, sm.updated_at,
       us.id, us.username
			 from social_medias sm
			 join users us
			 on sm.user_id = us.id;`)

	if err != nil {
		fmt.Println("err get-all sm", err.Error())
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var socialMedia entity.SocialMedia
		var timeAt sql.NullTime

		err := row.Scan(
			&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserID,
			&socialMedia.CreatedAt, &timeAt, &socialMedia.User.ID,
			&socialMedia.User.Username,
		)

		if err != nil {
			fmt.Println("errGetAll", err.Error())
			return nil, err
		}

		socialMedias = append(socialMedias, &socialMedia)
	}

	return socialMedias, nil
}

func (sp SocialMediaRepo) GetSocialMediaById(ctx context.Context, id string) (*entity.SocialMedia, error) {
	//TODO implement me
	sqlQuery := "select id, name, social_media_url, user_id, created_at, updated_at from social_medias where id=$1"
	row := sp.pgPool.QueryRow(ctx, sqlQuery, id)

	var sm entity.SocialMedia

	err := row.Scan(&sm.ID, &sm.Name, &sm.SocialMediaUrl, &sm.UserID, &sm.CreatedAt, &sm.UpdatedAt)

	if err != nil {
		fmt.Println("errByIdSm", err.Error())
		return nil, err
	}

	return &sm, nil
}

func (sp SocialMediaRepo) UpdateSocialMedia(ctx context.Context, id string, sm entity.SocialMedia) (*entity.SocialMedia, error) {
	//TODO implement me
	sqlQuery := "update social_medias set name=$1, social_media_url=$2, updated_at=now() where id=$3"

	_, err := sp.pgPool.Exec(ctx, sqlQuery, sm.Name, sm.SocialMediaUrl, id)

	if err != nil {
		fmt.Println("errUpdateSm", err.Error())
		return nil, err
	}

	return &sm, nil
}

func (sp SocialMediaRepo) DeleteSocialMedia(ctx context.Context, id string) error {
	//TODO implement me
	sqlQuery := "delete from social_medias where id=$1"

	_, err := sp.pgPool.Exec(ctx, sqlQuery, id)

	if err != nil {
		fmt.Println("errDeleteSm", err.Error())
		return err
	}

	return nil
}
