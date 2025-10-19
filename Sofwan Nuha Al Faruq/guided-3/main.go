package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	hasil := 0
	for i := 0; i < b; i++ {
		hasil += a
	}

	fmt.Println("Hasil perkalian:", hasil)
}
