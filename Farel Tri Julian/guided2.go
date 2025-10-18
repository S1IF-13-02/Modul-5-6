package main

import "fmt"

func main() {
	var n int

	fmt.Print("masukan jumlah segitiga : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var alas, tinggi int

		fmt.Print("masukan alas : ")
		fmt.Scan(&alas)
		fmt.Print("masukan tinggi : ")
		fmt.Scan(&tinggi)

		luas := (alas * tinggi) / 2
		fmt.Println("luas segitiga", luas)
	}
}
