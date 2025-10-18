package main

import "fmt"

func main() {

	var a, b int
	var hasil int

	fmt.Print("Masukkan a dan  b: ")
	fmt.Scan(&a, &b)

	hasil = 1

	for i := 1; i <= b; i++ {

		hasil = hasil * a
	}

	fmt.Printf("%d dipangkatkan %d adalah: %d\n", a, b, hasil)
}
