package main

import "fmt"

func main(){
	var angka1, angka2 int
	fmt.Scan(&angka1, &angka2)
	var hasil int = 0
	for i := 0 ; i < angka2; i++{
		hasil = hasil + angka1
	}
	fmt.Println(hasil)
}