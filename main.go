package main

import (
	"apibe23/configs"
	"apibe23/internal/controllers/todo"
	"apibe23/internal/controllers/users"
	"apibe23/internal/models"

	"github.com/labstack/echo/v4"
)

// type User struct {
// 	ID       int    `json:"id"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func main() {
	e := echo.New()

	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&models.User{},&models.Todo{})
	um := models.NewUserModel(db)
	uc := users.NewUserController(um)

	tm := models.NewTodoModel(db)
	tc := todo.NewTodoController(tm)
	// Register
	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login)
	e.POST("/addtodo", tc.CreateTodo)
	e.POST("/gettodo", tc.GetAllTodo)
	e.POST("/updatetodo", tc.UpdateTodo)
	e.POST("/deletetodo", tc.DeleteTodo)

	// Login
	// Tampilkan semua data
	// e.GET("/users", GetAllUsers)
	e.Start(":5000")
}
