package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&n)

	hasil := 0
	for i := 1; i <= n; i ++{
	hasil += i
	}
	fmt.Print("Hasil :", hasil)
	
}