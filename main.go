package main

import (
	"apibe23/configs"
	"apibe23/internal/controllers/todo"
	"apibe23/internal/controllers/users"
	"apibe23/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	// t := e.Group("/todos")
	// t.Use(echojwt.WithConfig(
	// 	echojwt.Config{
	// 		SigningKey:    []byte("passkeyJWT"),
	// 		SigningMethod: jwt.SigningMethodHS256.Name,
	// 	},
	// ))

	// Register
	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login)

	e.POST("/addTodo", tc.CreateTodo, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	e.GET("/showTodo", tc.GetAllTodo, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	e.POST("/updatetodo/:todoID", tc.UpdateTodo, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	e.POST("/deletetodo/:todoID", tc.DeleteTodo, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	// t := e.Group("/todos")
	// t.Use(echojwt.WithConfig(
	// 	echojwt.Config{
	// 		SigningKey:    []byte("passkeyJWT"),
	// 		SigningMethod: jwt.SigningMethodHS256.Name,
	// 	},
	// ))
	// t.PO("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())


	// Login
	// Tampilkan semua data
	// e.GET("/users", GetAllUsers)
	e.Start(":5000")
}
