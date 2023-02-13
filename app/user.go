package app

type RegisterUserResponse struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type RegisterUserRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

type LoginUserRequest struct {
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

type UpdateUserResponse struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required,email"`
}
