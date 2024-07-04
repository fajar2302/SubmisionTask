package todo

import "apibe23/internal/models"

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Title string `json:"title"`
}

func ToModelTodo(tr CreateTodoRequest, userID uint) models.Todo {
	return models.Todo{
		UserID: userID,
		Title:  tr.Title,
	}
}