package main
import "fmt"
 
func main() {
	var n int
	fmt.Print("Masukan bilangan bulat non negatif: ")
	fmt.Scan(&n)

	hasil := 1
	for i := 1; i <= n; i++ {
		hasil=hasil * i
		
	}
	fmt.Println(hasil)
}