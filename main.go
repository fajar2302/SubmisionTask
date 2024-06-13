package main

import (
	"fmt"
	"todo/configs"
	"todo/internal/controllers"
	"todo/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	connection.AutoMigrate(&models.User{}, &models.Todo{})

	var inputMenu int
	um := models.NewUserModel(connection)
	uc := controllers.NewUserController(um)

	tu := models.NewTodoModel(connection)
	tc := controllers.NewTodoController(tu)

	for inputMenu != 9 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var isLogin = true
			var inputMenu2 int
			data, err := uc.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}

			for isLogin {
				fmt.Println("Selamat datang ", data.Name, ",")
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah Kegiatan")
				fmt.Println("2. Update Kegiatan")
				fmt.Println("3. Hapus Kegiatan")
				fmt.Println("4. Tampilkan daftar kegiatan")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				} else if inputMenu2 == 1 {
					_, err := tc.AddTodo(data.ID)
					if err != nil {
						fmt.Println("error ketika menambahkan aktivitas")
						return
					}
					fmt.Println("berhasil menambahkan aktivitas")
				} else if inputMenu2 == 2 {
					var updatedTodo models.Todo
					fmt.Print("Masukkan ID kegiatan yang ingin diperbarui: ")
					fmt.Scanln(&updatedTodo.ID)
					fmt.Print("Masukkan aktivitas baru: ")
					fmt.Scanln(&updatedTodo.Activity)

					success, err := tc.UpdateTodo(updatedTodo.ID, updatedTodo)
					if err != nil {
						fmt.Println("Error ketika memperbarui kegiatan:", err)
					} else {
						if success {
							fmt.Println("Kegiatan berhasil diperbarui.")
						} else {
							fmt.Println("Gagal memperbarui kegiatan.")
						}
					}
				} else if inputMenu2 == 3 {
					var deletedTodoID uint
					fmt.Print("Masukkan ID kegiatan yang ingin dihapus: ")
					fmt.Scanln(&deletedTodoID)
				
					// Panggil fungsi DeleteTodo dari controller TodoController
					success, err := tc.DeleteTodo(deletedTodoID)
					if err != nil {
						fmt.Println("Error ketika menghapus kegiatan:", err)
					} else {
						if success {
							fmt.Println("Kegiatan berhasil dihapus.")
						} else {
							fmt.Println("Gagal menghapus kegiatan.")
						}
					}
				} else if inputMenu2 == 4 {
					todos, err := tc.GetAllTodos()
					if err != nil {
						fmt.Println("error ketika menampilkan aktivitas")
					} else {
						for _, todo := range todos {
							fmt.Println("Activity:", todo.Activity)
							fmt.Println("Owner:", todo.Owner)
							// Tampilkan properti lainnya sesuai kebutuhan
						}
					}

				}
			}

		} else if inputMenu == 2 {
			uc.Register()
		}
	}

	fmt.Println("terima kasih")

}