package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title  string `json:"title"`
    UserID uint   `json:"user_id"`
}

type TodoModel struct {
    db *gorm.DB
}

func NewTodoModel(db *gorm.DB) *TodoModel {
    return &TodoModel{db: db}
}

func (tm *TodoModel) InsertTodo(todo Todo) error {
    if err := tm.db.Create(&todo).Error; err != nil {
        return err
    }
    return nil
}

func (tm *TodoModel) GetAllTodoByUserID(userID uint) ([]Todo, error) {
    var todos []Todo
    if err := tm.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

func (tm *TodoModel) Update(todoID, userID uint, newData Todo) error {
    // Cari entri to-do berdasarkan todoID dan userID
    todo := Todo{}
    if err := tm.db.First(&todo, todoID).Error; err != nil {
        return err
    }

    // Periksa apakah userID sesuai dengan userID pada todo
    if todo.UserID != userID {
        return gorm.ErrRecordNotFound
    }

    // Update data to-do
    todo.Title = newData.Title

    // Simpan perubahan ke database
    if err := tm.db.Save(&todo).Error; err != nil {
        return err
    }

    return nil
}

func (tm *TodoModel) Delete(todoID, userID uint) error {
    // Hapus entri to-do berdasarkan todoID dan userID
    if err := tm.db.Where("id = ? AND user_id = ?", todoID, userID).Delete(&Todo{}).Error; err != nil {
        return err
    }
    return nil
}
