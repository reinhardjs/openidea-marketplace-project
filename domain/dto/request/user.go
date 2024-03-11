package request

type RegisterUserRequest struct {
	Username string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
