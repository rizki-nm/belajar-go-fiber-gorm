package web

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
}

type GetUserResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ErrorValidateResponse struct {
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}
