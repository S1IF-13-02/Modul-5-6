package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah kerucut (n): ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var r, t float64

		fmt.Printf("Masukkan jari-jari (r) dan tinggi (t) kerucut ke-%d: ", i+1)
		fmt.Scanln(&r, &t)

		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t

		fmt.Printf("Volume kerucut ke-%d: %.12f\n", i+1, volume)
	}
}
