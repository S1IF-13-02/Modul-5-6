package main
import "fmt"

func main() {
	var (
		i, r, t, n int
		volume float64
		π = 3.141592653589793238
	)

	fmt.Scan(&n)
	for i = 1; i <= n; i+=1 {
		fmt.Scan(&r, &t)
		volume = 1.0/3.0 * π * float64(r * r * t)
		fmt.Println(volume)
	}
}