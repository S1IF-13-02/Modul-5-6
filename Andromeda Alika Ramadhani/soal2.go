package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i int
	var jari2, tinggi float64
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&jari2, &tinggi)

		volume := (1.0 / 3.0) * math.Pi * (jari2 * jari2) * tinggi

		fmt.Print(volume)
	}
}
