package main

import "fmt"

func main() {
	var hasil, n int

	fmt.Print("Masukan Bilangan bulat positif: ")

	fmt.Scan(&n)

	if n < 0 {

		fmt.Println("Tidak bisa di faktorkan")
	}
	hasil = 1

	for i := 1; i <= n; i++ {

		hasil *= i

	}
	fmt.Println("Hasilnya adalah : ", hasil)
}
