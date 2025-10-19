package main

import "fmt"

func main() {
	var dasar, pangkat int
	hasil := 1

	fmt.Print("Masukkan dasar (bilangan pertama): ")
	fmt.Scanln(&dasar)

	fmt.Print("Masukkan pangkat (bilangan kedua): ")
	fmt.Scanln(&pangkat)

	for i := 0; i < pangkat; i++ {
		hasil *= dasar
	}

	fmt.Printf("\nHasil dari %d dipangkatkan %d adalah: %d\n", dasar, pangkat, hasil)
}
