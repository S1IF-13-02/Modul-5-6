package main

import (
	"fmt"
)

func main() {
	var n int
	var hasil uint64 = 1
	fmt.Print("Input nilai: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		hasil = hasil * uint64(i)
	}
	fmt.Printf("Hasil dari %d: %d\n ",n , hasil)
}