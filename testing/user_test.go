package testing

import (
	"context"
	"errors"
	"final-project/entity"
	"final-project/service"
	mockRepo "final-project/test/mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New User Service", func(t *testing.T) {
		mockUserRepo := mockRepo.NewMockUserRepositoryInterface(ctrl)
		userService := service.NewUserService(mockUserRepo)
		require.NotNil(t, userService)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty Email", func(t *testing.T) {
		mockUserRepo := mockRepo.NewMockUserRepositoryInterface(ctrl)
		userService := service.NewUserService(mockUserRepo)
		res, err := userService.Register(context.Background(), entity.User{Email: ""})

		require.Error(t, err)
		require.Equal(t, errors.New("error to validate email"), err)
		require.Nil(t, res)
	})

	t.Run("Success Register", func(t *testing.T) {
		mockUserRepo := mockRepo.NewMockUserRepositoryInterface(ctrl)
		userService := service.NewUserService(mockUserRepo)
		user := &entity.User{
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		userRes := &entity.User{
			ID:       1,
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		mockUserRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(userRes, nil)

		res, err := userService.Register(context.Background(), *user)

		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.ID)
	})
}
