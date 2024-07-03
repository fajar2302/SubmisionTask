package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title  string
    UserID uint
}

type TodoModel struct {
    db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
    return &TodoModel{
        db: connection,
    }
}

func (tm *TodoModel) InsertTodo(newData Todo) error {
    qry := tm.db.Create(&newData)
    if err := qry.Error; err != nil {
        return err
    }
    return nil
}

func (tm *TodoModel) GetAllTodo() ([]Todo, error) {
    var todos []Todo
    if err := tm.db.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

func (tm *TodoModel) Update(userID uint, newData Todo) error {
    // Cari entri to-do berdasarkan userID
    todo := Todo{}
    if err := tm.db.First(&todo, userID).Error; err != nil {
        return err
    }

    // Update data to-do
    todo.Title = newData.Title

    // Simpan perubahan ke database
    if err := tm.db.Save(&todo).Error; err != nil {
        return err
    }

    return nil
}

func (tm *TodoModel) Delete(userID uint) error {
    // Hapus entri to-do berdasarkan userID
    if err := tm.db.Delete(&Todo{}, userID).Error; err != nil {
        return err
    }
    return nil
}
