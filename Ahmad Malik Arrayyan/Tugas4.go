3.	Tugas 3
Source code
package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Masukkan bilangan dan pangkatnya : ")
    fmt.Scan(&a, &b)

    hasil := 1
    for i := 1; i <= b; i++ {
        hasil *= a
    }

    fmt.Printf("Hasil dari %d pangkat %d adalah %d\n", a, b, hasil)
}