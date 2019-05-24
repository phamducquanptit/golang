package user

import "rest-api/models"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type DataResponse struct {
	User models.UserModel `json:"user"`
}

type LoginResponse struct {
	User  models.UserModel `json:"user"`
	Token string           `json:"token"`
}

type SwaggerLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
