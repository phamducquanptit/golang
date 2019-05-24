package todo

import (
	"rest-api/models"
)

type CreateRequest struct {
	Title string `json:"title"`
}

type CreateResponse struct {
	Title string `json:"title"`
}

type UpdateRequest struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	Total uint64             `json:"total"`
	Todos []*models.TodoInfo `json:"todos"`
}

type DataResponse struct {
	Todos models.TodoModel `json:"todos"`
}

type UpdateResponse struct {
	Todo UpdateRequest `json:"todo"`
}

type SwaggerListResponse struct {
	Total uint64            `json:"total"`
	Todos []models.TodoInfo `json:"todos"`
}
