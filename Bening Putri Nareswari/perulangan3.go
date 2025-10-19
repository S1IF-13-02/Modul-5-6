package main

import (
	"fmt"
)

func main() {
	var A, B int
	hasil := 0

	fmt.Print("Masukkan bilangan pertama (A): ")
	fmt.Scanln(&A)

	fmt.Print("Masukkan bilangan kedua (B): ")
	fmt.Scanln(&B)

	fmt.Printf("\nMenghitung: %d x %d\n", A, B)
	fmt.Print("Proses: ")
	for i := 0; i < B; i++ {
		hasil += A
		fmt.Print(A)
		if i < B-1 {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", hasil)
}
