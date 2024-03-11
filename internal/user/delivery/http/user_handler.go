package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/openidea-marketplace/internal/domain"
	"github.com/openidea-marketplace/internal/domain/dto/request"
)

type UserHandler struct {
	Usecase domain.UserUsecase
}

func NewUserUsecase(echo *echo.Echo, usecase domain.UserUsecase) {
	handler := &UserHandler{
		Usecase: usecase,
	}
	echo.POST("/v1/user/register", handler.Register)
}

func (handler *UserHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var registerUserRequest request.RegisterUserRequest

	err := c.Bind(&registerUserRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), domain.ResponseError{
			Message: err.Error(),
		})
	}

	user, err := handler.Usecase.Register(ctx, &registerUserRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), domain.ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.ResponseSuccess{
		Message: "User registered successfully",
		Data:    user,
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
