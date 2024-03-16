package request

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}
