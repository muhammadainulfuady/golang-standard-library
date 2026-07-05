package main

import (
	"errors"
	"fmt"
)

var (
	ErrPembagian = errors.New("Tidak bisa di bagi dengan 0")
)

func bagi(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, fmt.Errorf("gagal menghitung: %w", ErrPembagian)
	}
	return angka / pembagi, nil
}

func main() {
	result, err := bagi(100, 0)
	if err != nil {
		if errors.Is(err, ErrPembagian) {
			fmt.Println("eror : ", err)
		} else {
			fmt.Println("Aman : ", result)
		}
	}
}
