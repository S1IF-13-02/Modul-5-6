package main 

import "fmt"

func main(){
	var awal,akhir int
	fmt.Println("Masukkan Awal: ")
	fmt.Scan(&awal)
	fmt.Println("Masukkan Akhir: ")
	fmt.Scan(&akhir)
	fmt.Println()
	for iterasi := awal ; iterasi <= akhir ; iterasi++{
		fmt.Print(iterasi, " ")
	}
}