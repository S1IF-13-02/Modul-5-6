package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukan kerucut : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var ja, tk float64

		fmt.Print("masukan jari-jari alas: ")
		fmt.Scan(&ja)
		fmt.Print("masukan tinggi : ")
		fmt.Scan(&tk)

		volume := (1.0 / 3.0) * math.Pi * ja * ja * tk

		fmt.Printf("Volume kerucut adalah: %.14f\n", volume)

	}
}
