package main

import (
	"fmt"
	"math"
)

func main() {
	var volume, r, t float64
	var n int

	fmt.Print("masukan banyaknya kerucut : ")
	fmt.Scan(&n)

	for j := 1; j <= n; j++ {
		fmt.Print(" masukan jari2 : ")
		fmt.Scan(&r)
		fmt.Print(" masukan tinggi : ")
		fmt.Scan(&t)

		volume = 1.0 / 3.0 * math.Pi * r * r * t

		fmt.Printf("volume kerucut adalah : %.14f", volume)
	}
}
