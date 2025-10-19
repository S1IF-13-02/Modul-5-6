package main

import (
	"fmt"
	"math"
)

func main() {
	var i, n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		var r, t float64
		fmt.Printf("Masukkan jari-jari dan tinggi kerucut ke-%d: ", i)
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t
		fmt.Println("Volume kerucut ke-", i, ":", volume)
	}
}
