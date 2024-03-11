package request

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
