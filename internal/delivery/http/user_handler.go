package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/openidea-marketplace/domain/dto/request"
	"github.com/openidea-marketplace/domain/dto/response"
	"github.com/openidea-marketplace/user"
)

type UserHandler struct {
	Usecase user.Usecase
}

func NewUserHandler(usecase user.Usecase) *UserHandler {
	return &UserHandler{
		Usecase: usecase,
	}
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

func (handler *UserHandler) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	var loginUserRequest request.LoginUserRequest

	err := c.BodyParser(&loginUserRequest)
	if err != nil {
		return c.Status(getStatusCode(err)).JSON(response.ResponseError{
			Message: err.Error(),
		})
	}

	user, err := handler.Usecase.Login(ctx, &loginUserRequest)
	if err != nil {
		return c.Status(getStatusCode(err)).JSON(response.ResponseError{
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(response.ResponseSuccess{
		Message: "User login successfully",
		Data:    user,
	})
}
