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
    userID, _ := strconv.ParseUint(c.Param("userID"), 10, 64)


    var reqBody CreateTodoRequest
    if err := c.Bind(&reqBody); err != nil {
        return c.JSON(500, helper.ResponseFormat(500, "input error", nil))
    }

    todo := models.Todo{
        Title:  reqBody.Title,
        UserID: uint(userID), // Assign userID to the Todo struct
    }

    if err := tc.todoModel.InsertTodo(todo); err != nil {
        return c.JSON(500, helper.ResponseFormat(500, "insert error", nil))
    }

    return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (tc *TodoController) GetAllTodo(c echo.Context) error {
    userID, _ := strconv.ParseUint(c.Param("userID"), 10, 64)

    todos, err := tc.todoModel.GetAllTodoByUserID(uint(userID))
    if err != nil {
        return c.JSON(400, helper.ResponseFormat(400, "input error", nil))

    }

    var response []TodoResponse
    for _, todo := range todos {
        response = append(response, NewTodoResponse(todo))
    }

    return c.JSON(http.StatusOK, response)
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {
    userID, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
    todoID, _ := strconv.ParseUint(c.Param("todoID"), 10, 64)

    var reqBody UpdateTodoRequest
    if err := c.Bind(&reqBody); err != nil {
        return c.JSON(400, helper.ResponseFormat(400, "input error", nil))

    }

    newData := models.Todo{
        Title: reqBody.Title,
    }

    if err := tc.todoModel.Update(uint(todoID), uint(userID), newData); err != nil {
        return c.JSON(500, helper.ResponseFormat(500, "update error", nil))

    }

    return c.JSON(201, helper.ResponseFormat(201, "todo updated successfully", nil))
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
    userID, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
    

    todoID, _ := strconv.ParseUint(c.Param("todoID"), 10, 64)
   

    if err := tc.todoModel.Delete(uint(todoID), uint(userID)); err != nil {
        return c.JSON(400, helper.ResponseFormat(400, "failed deleted todo", nil))
    }

    return c.JSON(201, helper.ResponseFormat(201, "success deleted Todo", nil))
}
