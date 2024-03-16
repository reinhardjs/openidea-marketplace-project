package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
	"github.com/openidea-marketplace/user/usecases"
)

type UserHandler struct {
	Usecase usecases.Usecase
}

func NewUserUsecase(usecase usecases.Usecase) *UserHandler {
	handler := &UserHandler{
		Usecase: usecase,
	}

	return handler
}

func (handler *UserHandler) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	var registerUserRequest request.RegisterUserRequest

	err := c.BodyParser(&registerUserRequest)
	if err != nil {
		return c.Status(getStatusCode(err)).JSON(response.ResponseError{
			Message: err.Error(),
		})
	}

	user, err := handler.Usecase.Register(ctx, &registerUserRequest)
	if err != nil {
		return c.Status(getStatusCode(err)).JSON(response.ResponseError{
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(response.ResponseSuccess{
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
