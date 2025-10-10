package model

type UsersRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func ToUserResponse(users string, id int) UsersResponseRegister {
	return UsersResponseRegister{
		Id:   id,
		Name: users,
	}
}

type UsersResponseRegister struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UserGetProfileResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
