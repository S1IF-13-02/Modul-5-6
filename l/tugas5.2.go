package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t float64

	fmt.Print(" masukan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("masukan r dan t kerucut ke-", i, ":")
		fmt.Scan(&r, &t)

		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println(volume)

	}
}
