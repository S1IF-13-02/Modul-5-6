package main

import "fmt"

func main() {
	var angka, pangkat int
	var result int = 1

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&angka)
	fmt.Print("Masukkan Pangkat: ")
	fmt.Scan(&pangkat)

	for i := 1; i <= pangkat; i++ {
		result = result * angka
	}

	fmt.Println(result)
}