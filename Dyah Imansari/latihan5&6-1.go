package main
import "fmt"

func main()  {
	var i, n int
	var hasil int
	hasil = 0
	
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil += i
	}
	fmt.Println(hasil)
}