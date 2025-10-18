package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua bilangan bulat positif (basis dan pangkat): ")
	fmt.Scan(&a, &b)

	hasil := 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}

	fmt.Printf("Hasil dari %d pangkat %d adalah %d", a, b, hasil)
}
