package main

import "fmt"
func main(){
	var n, x int
	fmt.Print("Masukan Awal  : ")
	fmt.Scan(&x)
	fmt.Print("Masukan Batas : ")
	fmt.Scan(&n)
	for iterasi := x ; iterasi <= n ; iterasi++{
	fmt.Print(iterasi, " ")
	}
}