package main

import (
	"fmt"
	"math"
)

func HitungVolumeKerucut(r, t int) float64 {
	return 1.0 / 3.0 * math.Pi * math.Pow(float64(r), 2) * float64(t)
}

func main() {
	var n, r, t int

	fmt.Print("Masukkan Nilai (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan (r) dan (t): ")
		fmt.Scan(&r, &t) 

		result := HitungVolumeKerucut(r, t)

		fmt.Println(result) 
	}

}