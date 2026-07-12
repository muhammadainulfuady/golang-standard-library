package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrTerlaluPendek = errors.New("username minimal 5 karakter")
	ErrKataTerlarang = errors.New("username mengandung kata terlarang")
)

func validasiUsername(input string) (string, error) {
	// TUGAS:
	// 1. Bersihin spasi di awal/akhir input (pakai strings.TrimSpace)
	data := strings.TrimSpace(input)
	// 2. Cek panjang (pakai len()), kalau < 5 karakter, return ErrTerlaluPendek
	if len(data) < 5 {
		return "", ErrTerlaluPendek
	}
	// 3. Cek apakah mengandung "admin" atau "root" — HARUS case-insensitive
	//    (hint: ubah ke lowercase dulu sebelum dicek pakai strings.Contains)
	lowerData := strings.ToLower(data)
	if strings.Contains(lowerData, "admin") || strings.Contains(lowerData, "root") {
		return "", ErrKataTerlarang
	}

	// 4. Kalau lolos semua, return versi lowercase-nya, error nil
	return lowerData, nil

}

func main() {
	testCases := []string{
		"   Budi123   ",
		"ab",
		"SuperAdmin99",
		"johnDoe",
	}

	for _, tc := range testCases {
		hasil, err := validasiUsername(tc)
		if err != nil {
			fmt.Printf("Input: %-20q -> Error: %s\n", tc, err)
		} else {
			fmt.Printf("Input: %-20q -> Valid: %s\n", tc, hasil)
		}
	}
}
