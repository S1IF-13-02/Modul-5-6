package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Masukkan niali a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	hasil := 0
	for i := 1; i <= b; i++ {
		hasil = hasil + a
	}

	fmt.Println("Hasil perkalian: ", hasil)
}
