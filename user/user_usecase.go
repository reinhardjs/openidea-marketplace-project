package user

import (
	"context"
	"strconv"
	"time"

	"github.com/openidea-marketplace/auth"
	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
	"github.com/openidea-marketplace/domain/entities"
	"github.com/openidea-marketplace/pkg/utils/hashing"
)

type Usecase interface {
	Register(ctx context.Context, request *request.RegisterUserRequest) (response.RegisterUserResponse, error)
}

type Repository interface {
	Insert(ctx context.Context, request *entities.User) error
}

type userUsecase struct {
	Repository     Repository
	ContextTimeout time.Duration
	Log            domain.Logger
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

	var user = entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}

	hashSalt, err := usecase.Hashing.GenerateHash([]byte(request.Password))
	if err != nil {
		return
	}
	user.Password = string(hashSalt.Hash)

	err = usecase.Repository.Insert(ctx, &user)
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
