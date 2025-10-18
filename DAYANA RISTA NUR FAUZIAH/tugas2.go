package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var jariJari, tinggi float64
	var volume float64

	fmt.Print("Masukkan jumlah kerucut (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Scan(&jariJari, &tinggi)

		volume = (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi

		fmt.Println(volume)
	}
}
