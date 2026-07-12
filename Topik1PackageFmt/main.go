package main

import "fmt"

func main() {
	var nama string
	var umur int
	fmt.Println("Masukkan nama anda : ")
	fmt.Scanln(&nama)

	fmt.Println("Halo", nama)
	fmt.Println(umur)
}
