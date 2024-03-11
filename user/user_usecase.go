package user

import (
	"context"
	"strconv"
	"time"

	"github.com/openidea-marketplace/auth"
	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
	"github.com/openidea-marketplace/hashing"
)

type Usecase interface {
	Register(ctx context.Context, request *request.RegisterUserRequest) (response.RegisterUserResponse, error)
}

type Repository interface {
	Register(ctx context.Context, request *domain.User) error
}

type userUsecase struct {
	Repository     Repository
	ContextTimeout time.Duration
	AuthUsecase    auth.Usecase
	Hashing        hashing.Hashing
}

func NewUsecase(repository Repository, timeout time.Duration, authUsecase auth.Usecase, hashing hashing.Hashing) Usecase {
	return &userUsecase{
		Repository:     repository,
		ContextTimeout: timeout,
		AuthUsecase:    authUsecase,
		Hashing:        hashing,
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

	hashSalt, err := usecase.Hashing.GenerateHash([]byte(request.Password))
	if err != nil {
		return
	}
	user.Password = string(hashSalt.Hash)

	err = usecase.Repository.Register(ctx, &user)
	if err != nil {
		return
	}

	decimalBase := 10
	token, err := usecase.AuthUsecase.GenerateToken(strconv.FormatInt(user.ID, decimalBase), user.Username, user.Password)
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
