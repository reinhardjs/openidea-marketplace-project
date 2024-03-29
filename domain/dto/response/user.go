package response

type RegisterUserResponse struct {
	Username    string `json:"id"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginUserResponse struct {
	Username    string `json:"id"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
