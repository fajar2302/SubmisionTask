package todo

import (
	"net/http"
	"strconv"

	"apibe23/internal/helper"
	"apibe23/internal/models"
	"apibe23/internal/utils"

	"github.com/golang-jwt/jwt/v5"
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
	// Mendapatkan token dari context
    user := c.Get("user").(*jwt.Token)

    // Decode token untuk mendapatkan userID
    userID, err := utils.DecodeToken(user)
    if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid request body", nil))
	}

	var reqBody CreateTodoRequest
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid request body", nil))
	}

	todo := models.Todo{
		Title:  reqBody.Title,
		UserID: uint(userID), // Assign userID to the Todo struct
	}

	if err := tc.todoModel.InsertTodo(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed to create todo", nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Todo created successfully", nil))
}

func (tc *TodoController) GetAllTodo(c echo.Context) error {
	// Mendapatkan token dari context
    user := c.Get("user").(*jwt.Token)

    // Decode token untuk mendapatkan userID
    userID, err := utils.DecodeToken(user)

	todos, err := tc.todoModel.GetAllTodoByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed to fetch todos", nil))
	}

	var response []TodoResponse
	for _, todo := range todos {
		response = append(response, NewTodoResponse(todo))
	}

	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success", response))
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {
	
	// Mendapatkan token dari context
    user := c.Get("user").(*jwt.Token)

    // Decode token untuk mendapatkan userID
    userID, err := utils.DecodeToken(user)

	todoID, err := strconv.ParseUint(c.Param("todoID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid todo ID", nil))
	}

	var reqBody UpdateTodoRequest
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid request body", nil))
	}

	newData := models.Todo{
		Title: reqBody.Title,
	}

	if err := tc.todoModel.Update(uint(todoID), uint(userID), newData); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed to update todo", nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Todo updated successfully", nil))
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	// Mendapatkan token dari context
    user := c.Get("user").(*jwt.Token)

    // Decode token untuk mendapatkan userID
    userID, err := utils.DecodeToken(user)
	todoID, err := strconv.ParseUint(c.Param("todoID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid todo ID", nil))
	}

	if err := tc.todoModel.Delete(uint(todoID), uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed to delete todo", nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Todo deleted successfully", nil))
}
