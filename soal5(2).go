package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var r float64
		var t float64
		fmt.Print("Masukkan alas dan tinggi kerucut: ")
		fmt.Scan(&r, &t)

		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t

		fmt.Println("Volume kerucut: ", volume)
	}
}
