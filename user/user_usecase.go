package user

import (
	"context"
	"time"

	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
)

type Usecase interface {
	Register(ctx context.Context, request *request.RegisterUserRequest) (response.RegisterUserResponse, error)
}

type Repository interface {
	Register(ctx context.Context, request *domain.User) (domain.User, error)
}

type userUsecase struct {
	Repository     Repository
	ContextTimeout time.Duration
}

func NewUsecase(repository Repository, timeout time.Duration) Usecase {
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
