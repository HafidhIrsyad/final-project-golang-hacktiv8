package service

import (
	"context"
	"errors"
	"final-project/entity"
	"final-project/helper"
	"final-project/repository"
	"fmt"
	"time"
)

type UserServiceInterface interface {
	Register(ctx context.Context, user entity.User) (*entity.User, error)
	Login(ctx context.Context, input entity.UserLogin) (*entity.User, error)
	Update(ctx context.Context, id string, user entity.UserLogin) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{userRepo: userRepo}
}

func (u UserService) Update(ctx context.Context, id string, user entity.UserLogin) (*entity.User, error) {
	//TODO implement me
	userId, errGetId := u.userRepo.GetById(ctx, id)

	if errGetId != nil {
		return nil, errGetId
	}

	userId.Email = user.Email
	userId.Username = user.Username
	userId.UpdatedAt = time.Now()

	updateUser, err := u.userRepo.Update(ctx, *userId)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (u UserService) GetById(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	userId, err := u.userRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (u UserService) Login(ctx context.Context, input entity.UserLogin) (*entity.User, error) {
	//TODO implement me
	email := input.Email
	password := input.Password

	user, err := u.userRepo.Login(ctx, email)

	if err != nil {
		return nil, err
	}

	if !helper.CheckPassHash(password, user.Password) {
		return nil, errors.New("password incorrect")
	}

	return user, nil
}

func (u UserService) Register(ctx context.Context, user entity.User) (*entity.User, error) {
	//TODO implement me
	if _, ok := helper.ValidMailAddress(user.Email); !ok {
		return nil, errors.New("error to validate email")
	}
	if user.Email == "" {
		return nil, errors.New("email must be filled")
	}
	if user.Username == "" {
		return nil, errors.New("username must be filled")
	}
	if user.Password == "" {
		return nil, errors.New("password must be filled ")
	}

	newUser, err := u.userRepo.Register(ctx, user)

	if err != nil {
		fmt.Println("errRegisterSer", err.Error())
		return nil, err
	}

	return newUser, nil
}
