package usecase

import (
	"context"
	"time"

	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
)

type userUsecase struct {
	Repository     domain.UserRepository
	ContextTimeout time.Duration
}

func NewUserUsecase(repository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		Repository:     repository,
		ContextTimeout: timeout,
	}
}

func (usecase *userUsecase) Register(c context.Context, request *request.RegisterUserRequest) (res response.RegisterUserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.ContextTimeout)
	defer cancel()

	var user = domain.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}

	user, err = usecase.Repository.Register(ctx, &user)

	res = response.RegisterUserResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: "abcd-efgh-1234",
	}

	return
}
