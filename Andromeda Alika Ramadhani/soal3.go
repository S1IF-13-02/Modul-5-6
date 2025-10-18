package main

import "fmt"

func main() {
	var angka, pangkat, hasil, i int
	fmt.Scan(&angka, &pangkat)

	hasil = 1
	for i = 1; i <= pangkat; i++ {
		hasil = hasil * angka
	}
	fmt.Print(hasil)
}
