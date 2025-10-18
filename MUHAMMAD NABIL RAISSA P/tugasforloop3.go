package main

import "fmt"

func main() {
    var basis int
    var eksponen int

    fmt.Print("Masukkan masukan (basis eksponen): ")
    fmt.Scan(&basis, &eksponen)

    hasil := 1

    for i := 0; i < eksponen; i++ {
        hasil = hasil * basis
    }

    fmt.Println(hasil)
}