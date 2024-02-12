package main

import (
	"be21/config"
	"be21/users"
	"errors"
	"fmt"
)

func main() {
	// code 0
	database := config.InitMysql()
	config.Migrate(database)
	var input int
	// code 1
	for input != 99 {
		var a = 1
		var b = 5
		fmt.Println(b + a)
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		// code 5
		if input == 1 {
			var isRunning bool = true
			for isRunning {
				fmt.Println("halo bro")
				fmt.Println("hahaa")
				var loggedIn users.Users
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
