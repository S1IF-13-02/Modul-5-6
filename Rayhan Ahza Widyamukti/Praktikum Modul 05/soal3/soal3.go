package main

import "fmt"

func main() {
	var a, b, i int
	fmt.Print("Masukkan dua bilangan (a dan b): ")
	fmt.Scan(&a, &b)

	hasil := 1
	for i = 1; i <= b; i++ {
		hasil *= a
	}

	fmt.Println("Hasil", a, "pangkat", b, "adalah:", hasil)
}
