package main 

import "fmt"

func main(){
	var a,b int 
	fmt.Scan(&a)
	fmt.Scan(&b)
	hasil := 1
	for i := 1 ; i <= b ; i++{
		hasil = hasil * a 
	}
	fmt.Println(hasil)
}