package repository

import (
	"context"
	"final-project/entity"
	"final-project/helper"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepositoryInterface interface {
	Register(ctx context.Context, user entity.User) (*entity.User, error)
	Login(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user entity.User) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
}

type UserRepository struct {
	pgpool *pgxpool.Pool
}

func (u UserRepository) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	//TODO implement me
	sql := "update users set email=$1, username=$ where id=$3"
	_, err := u.pgpool.Exec(ctx, sql, user.Email, user.Username, user.ID)

	if err != nil {
		fmt.Println("errUpdateRepo: " + err.Error())
		return nil, err
	}
	return &user, nil

}

func (u UserRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	var user entity.User

	sql := `select id, username, email, password, age from users where id=$1`
	row := u.pgpool.QueryRow(ctx, sql, id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func NewUserRepository(pgpool *pgxpool.Pool) UserRepositoryInterface {
	return &UserRepository{pgpool: pgpool}
}

func (u UserRepository) Login(ctx context.Context, email string) (*entity.User, error) {
	//TODO implement me
	sql := "select id, username, email, password, age from users where email=$1"
	row := u.pgpool.QueryRow(ctx, sql, email)

	var userLogin entity.User

	err := row.Scan(&userLogin.ID, &userLogin.Username, &userLogin.Email, &userLogin.Password, &userLogin.Age)

	if err != nil {
		return nil, err
	}

	return &userLogin, nil
}

func (u UserRepository) Register(ctx context.Context, user entity.User) (*entity.User, error) {
	//TODO implement me
	password, errHash := helper.HashPassword(user.Password)

	if errHash != nil {
		fmt.Println("errHash: " + errHash.Error())
		return nil, errHash
	}

	user.Password = password

	sql := "insert into users(username, password, age, email, created_at) values($1, $2, $3, $4, now())"

	_, err := u.pgpool.Exec(ctx, sql, user.Username, user.Password, user.Age, user.Email)

	if err != nil {
		fmt.Println("errRegisterRepo: " + err.Error())
		return nil, err
	}
	return &user, nil
}
