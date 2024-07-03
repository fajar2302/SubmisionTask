package todo

import (
	"net/http"
	"strconv"

	"apibe23/internal/helper"
	"apibe23/internal/models"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
    todoModel *models.TodoModel
}

func NewTodoController(todoModel *models.TodoModel) *TodoController {
    return &TodoController{
        todoModel: todoModel,
    }
}

func (tc *TodoController) CreateTodo(c echo.Context) error {
    var reqBody CreateTodoRequest
    if err := c.Bind(&reqBody); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }

    todo := models.Todo{
        Title:  reqBody.Title,
        UserID: reqBody.UserID,
    }

    if err := tc.todoModel.InsertTodo(todo); err != nil {
        return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "Todo created successfully"})
}

func (tc *TodoController) GetAllTodo(c echo.Context) error {
    todos, err := tc.todoModel.GetAllTodo()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
    }

    var response []TodoResponse
    for _, todo := range todos {
        response = append(response, NewTodoResponse(todo))
    }

    return c.JSON(http.StatusOK, response)
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {
    userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    var reqBody UpdateTodoRequest
    if err := c.Bind(&reqBody); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }

    newData := models.Todo{
        Title: reqBody.Title,
    }

    if err := tc.todoModel.Update(uint(userID), newData); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Todo updated successfully"})
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
    userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    if err := tc.todoModel.Delete(uint(userID)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete todo"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
}
