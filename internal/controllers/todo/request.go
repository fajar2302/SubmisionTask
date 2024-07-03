package todo

type CreateTodoRequest struct {
    Title  string `json:"title"`
}

type UpdateTodoRequest struct {
    Title string `json:"title"`
}
