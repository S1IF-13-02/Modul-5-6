package main

import "fmt"

func main (){

	var jml int
	var luas, alas, tinggi float32

	fmt.Print("Masukan jumlah segitiga : ")
	fmt.Scan(&jml)

	for segitiga := 1 ; segitiga <= jml; segitiga++{
		fmt.Print("Masukan alas (", segitiga,") : ")
		fmt.Scan(&alas)
		fmt.Print("Masukan tinggi (", segitiga,") : ")
		fmt.Scan(&tinggi)

		luas = 0.5 * alas * tinggi
		fmt.Println("luas segitiga : ", luas, " ")
	}
}
