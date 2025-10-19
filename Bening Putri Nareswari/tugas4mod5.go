package main

import "fmt"

func main() {
	var n int
	hasil := 1

	fmt.Print("Masukkan bilangan bulat non-negatif (n): ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		hasil *= i
	}

	fmt.Printf("\nHasil faktorial dari %d adalah: %d\n", n, hasil)
}
