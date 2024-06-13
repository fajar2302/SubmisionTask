package controllers

import (
	"fmt"
	"todo/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (tc *TodoController) UpdateTodo(todoID uint, updatedData models.Todo) (bool, error) {
    err := tc.model.UpdateTodo(todoID, updatedData)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (tc *TodoController) DeleteTodo(todoID uint) (bool, error) {
    err := tc.model.DeleteTodo(todoID)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (tc *TodoController) GetAllTodos() ([]models.Todo, error) {
	todos, err := tc.model.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

