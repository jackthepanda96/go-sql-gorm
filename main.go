package main

import (
	"be21/config"
	"be21/users"
	"errors"
	"fmt"
)

func main() {
	// aku mulai ngerjain
	database := config.InitMysql()
	config.Migrate(database)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		if input == 1 {
			var isRunning bool = true
			for isRunning {
				var hp string
				var password string
				var loggedIn users.Users
				// fmt.Println("Masukkan HP")
				fmt.Scanln(&hp)
				// fmt.Println("Masukkan Password")
				fmt.Scanln(&password)
				fmt.Println("halo bro")
				err := errors.New("herror")
				if err == nil {
					fmt.Println("Selamat Datang,", loggedIn.Nama)
				} else {
					var inputExit string
					fmt.Print("Input 'EXIT' untuk kembali ke menu sebelumnya")
					fmt.Scanln(&inputExit)
					if inputExit == "EXIT" {
						isRunning = false
					}
				}
			}

			// kalo sukses welcome, kalo gagal isi lagi
		} else if input == 2 {
			var newUser users.Users
			fmt.Print("Masukkan nama ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan nomor HP ")
			fmt.Scanln(&newUser.HP)
			fmt.Print("Masukkan password ")
			fmt.Scanln(&newUser.Password)
			fmt.Print("Masukkan alamat ")
			fmt.Scanln(&newUser.Alamat)
			success, err := users.Register(database, newUser)
			if err != nil {
				fmt.Println("terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}

			if success {
				fmt.Println("selamat anda telah terdaftar")
			}
		}
	}
	fmt.Println("Exited! Thank you")

}
