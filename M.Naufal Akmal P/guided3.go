package main

import "fmt"

func main (){
	var x, y, hasil int
	fmt.Print("Masukan angka 1 : ")
	fmt.Scan(&x)
	fmt.Print("Masukan angka 2 : ")
	fmt.Scan(&y)

	for i := 1; i <= y ; i++{
		hasil = hasil + x
	}
	fmt.Println(" ")
	fmt.Print("Hasilnya : ", hasil, " ")
	fmt.Println(" ")
}