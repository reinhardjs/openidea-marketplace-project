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
	Login(ctx context.Context, request *request.LoginUserRequest) (response.LoginUserResponse, error)
}

type Repository interface {
	Insert(ctx context.Context, request *entities.User) error
	FindByUsername(ctx context.Context, username string) (entities.User, error)
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

	hash, err := usecase.Hashing.GenerateHashFromPassword(request.Password)
	if err != nil {
		return
	}
	user.Password = hash

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

func (usecase *userUsecase) Login(c context.Context, request *request.LoginUserRequest) (res response.LoginUserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.ContextTimeout)
	defer cancel()

	user, err := usecase.Repository.FindByUsername(ctx, request.Username)
	if err != nil {
		res = response.LoginUserResponse{}
		return
	}

	isMatch, err := usecase.Hashing.Compare(user.Password, request.Password)
	if err != nil {
		return
	}

	if !isMatch {
		err = domain.ErrUserWrongPassword
		return
	}

	decimalBase := 10
	token, err := usecase.AuthUsecase.GenerateToken(strconv.FormatInt(user.ID, decimalBase), user.Username, user.Password)
	if err != nil {
		return
	}

	res = response.LoginUserResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}

	return
}
