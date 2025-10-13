package main
import "fmt"

func main() {
	var a, b int
	
	fmt.Print("masukkan dua bilangan bulat: ")
	fmt.Scan(&a)
	
	fmt.Print("masukkan dua bilangan bulat: ")
	fmt.Scan(&b)
	
	hasil := 0
	for i := 0; i < b; i++ {
		hasil += a
	} 
	 fmt.Println(hasil)
}