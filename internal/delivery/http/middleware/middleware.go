package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/openidea-marketplace/auth"
	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/dto/request"
)

func NewAuthMiddleware(authUsecase auth.Usecase, log domain.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &request.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		log.Debug(fmt.Sprintf("Authorization : %s", request.Token))

		issuerUsername, err := authUsecase.VerifyToken(request.Token)
		if err != nil {
			log.Warn(fmt.Sprintf("Failed find user by token : %+v", err))
			return fiber.ErrUnauthorized
		}

		log.Debug(fmt.Sprintf("User : %+v", issuerUsername))
		ctx.Locals("auth", issuerUsername)
		return ctx.Next()
	}
}
