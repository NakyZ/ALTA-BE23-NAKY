package main

import (
	"fmt"
	"try/internal/controllers"
	"try/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type genre struct {
	ID int
	Name string
}

func connectDB() ( *gorm.DB, error ) {
	// dsn := dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	var connStr = "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.jratzfdofxgfgondhztd password=xbA5JfgxsajXDdPf port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func GetALLGenre(db *gorm.DB) ([]genre, error) {
	var result []genre
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result,nil
}

func main() {
	db,err := connectDB()
	if err != nil {
		fmt.Println(" Warning !!!", err.Error())
		return
	}
	// Code -> database
	hasil, _ := GetALLGenre(db)
	fmt.Println(hasil)

	// setup := configs.ImportSetting()
	// connection, err := configs.ConnectDB(setup)
	// if err != nil {
	// 	fmt.Println("Stop program, masalah pada database", err.Error())
	// 	return
	// }

	// connection.AutoMigrate(&models.User{}, &models.Todo{})

	var inputMenu int
	um := models.NewUserModel(db)
	uc := controllers.NewUserController(um)

	tu := models.NewTodoModel(db)
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
				}
			}

		} else if inputMenu == 2 {
			uc.Register()
		}
	}

	fmt.Println("terima kasih")

}
