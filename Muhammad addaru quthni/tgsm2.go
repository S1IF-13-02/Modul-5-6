package main

import (
	"fmt"
)

func main() {
	var n, j int
	fmt.Println("input: ")
	fmt.Scan(&n, &j)
	hasil := 1
	for i := 0; i < j; i++ {
		hasil = hasil * n
	}
	fmt.Print(hasil)
}
