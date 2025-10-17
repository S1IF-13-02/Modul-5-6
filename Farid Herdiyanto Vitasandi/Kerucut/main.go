package main

import "fmt"

func main() {
	 
	var jumlahKerucut int

	fmt.Print("Masukkan jumlah kerucut yang akan dihitung: ")
	fmt.Scan(&jumlahKerucut)

	for i := 1; i <= jumlahKerucut; i++{

	var r, t float64	// -----> r = jari-jari; t = tinggi
	const phi float64 = 3.14

	fmt.Print("Masukkan jari-jari dan tinggi kerucut(",i,"): ")
	fmt.Scan(&r, &t)

	volumeKerucut := (1.0 / 3.0) * phi * r * r * t

	fmt.Println("Volume kerucut (",i,"): ", volumeKerucut)

	}
}
