package userdto

type CreateUser struct {
	Name     string `json:"name" from:"name" validate:"required"`
	Email    string `json:"email" from:"email" validate:"required"`
	Password string `json:"password" from:"id" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" from:"name"`
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"id"`
}

type UserResponse struct {
	Name     string `json:"name" from:"name"`
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"id"`
}