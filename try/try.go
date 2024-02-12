package main

import "fmt"

func main() {
	// apakah ada solusi/pendekatan lain untuk membuat suatu program berjalan terus hingga sebuah kondisi terpenuhi?

	var isRunning = true

	for isRunning {
		var input int
		fmt.Scanln(&input)
		if input == 99 {
			isRunning = false
		}
	}

	fmt.Println("End program")
}
