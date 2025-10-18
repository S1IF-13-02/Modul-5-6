package main

import "fmt"

func main() {
	var bp, bk int
	hasil := 1

	fmt.Print("masukan angka : ")
	fmt.Scan(&bp)
	fmt.Print("masukan pangkat : ")
	fmt.Scan(&bk)

	for i := 1; i <= bk; i++ {
		hasil = hasil * bp
	}
	fmt.Println(hasil)
}
