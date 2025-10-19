package main

import "fmt"

func main() {
	var A int
	var B int
	var hasil int

	fmt.Print("Masukkan 2 bilangan: ")
	fmt.Scan(&A, &B)

	hasil = 1
	for i := 0; i < B; i++ {
		hasil = hasil * A
	}

	fmt.Println(hasil)
}
