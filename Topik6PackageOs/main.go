package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("data.js")

	if err != nil {
		log.Fatal("Gagal membuat file", err)
	}
	defer file.Close()

	file.WriteString("function gokilsih")

}
