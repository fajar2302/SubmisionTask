package todo

import "apibe23/internal/models"

type TodoResponse struct {
    ID    uint   `json:"id"`
    Title string `json:"title"`
}

func NewTodoResponse(todo models.Todo) TodoResponse {
    return TodoResponse{
        ID:    todo.ID,
        Title: todo.Title,
    }
}
