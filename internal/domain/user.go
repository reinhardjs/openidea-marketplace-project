package domain

import (
	"context"

	"github.com/openidea-marketplace/internal/domain/dto/request"
	"github.com/openidea-marketplace/internal/domain/dto/response"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUsecase interface {
	Register(ctx context.Context, request *request.RegisterUserRequest) ([]response.RegisterUserResponse, error)
}

type UserRepository interface {
	Register(ctx context.Context, request User) (User, error)
}
