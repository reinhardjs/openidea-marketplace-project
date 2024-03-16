package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/openidea-marketplace/auth"
	"github.com/openidea-marketplace/domain/dto/request"
)

func NewAuthMiddleware(authUsecase auth.Usecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &request.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		// authUsecase.Log.Debugf("Authorization : %s", request.Token)

		issuerUsername, err := authUsecase.VerifyToken(request.Token)
		if err != nil {
			// authUsecase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		// authUsecase.Log.Debugf("User : %+v", issuerUsername)
		ctx.Locals("auth", issuerUsername)
		return ctx.Next()
	}
}
