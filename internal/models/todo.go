package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) UpdateTodo(todoID uint, updatedData Todo) error {
    // Cari aktivitas berdasarkan ID
    todo := &Todo{}
    if err := tm.db.First(&todo, todoID).Error; err != nil {
        return err
    }

    // Perbarui aktivitas dengan data yang baru
    todo.Activity = updatedData.Activity

    // Simpan perubahan ke dalam database
    if err := tm.db.Save(&todo).Error; err != nil {
        return err
    }

    return nil
}

func (tm *TodoModel) DeleteTodo(todoID uint) error {
    var todo Todo
    if err := tm.db.Where("id = ?", todoID).Delete(&todo).Error; err != nil {
        return err
    }
    return nil
}

func (tm *TodoModel) GetAllTodos() ([]Todo, error) {
    var todos []Todo
    err := tm.db.Find(&todos).Error
    if err != nil {
        return nil, err
    }

    return todos, nil
}

