package user

import (
	"context"
	"strconv"
	"time"

	"github.com/openidea-marketplace/auth"
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
	AuthUsecase    auth.Usecase
}

func NewUsecase(repository Repository, timeout time.Duration, authUsecase auth.Usecase) Usecase {
	return &userUsecase{
		Repository:     repository,
		ContextTimeout: timeout,
		AuthUsecase:    authUsecase,
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
	if err != nil {
		return
	}

	base := 10 // decimal
	token, err := usecase.AuthUsecase.GenerateToken(strconv.FormatInt(user.ID, base), user.Username)
	if err != nil {
		return
	}

	res = response.RegisterUserResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}

	return
}
