package todo

type CreateTodoRequest struct {
    Title  string `json:"title"`
    UserID uint   `json:"user_id"`
}

type UpdateTodoRequest struct {
    Title string `json:"title"`
}
