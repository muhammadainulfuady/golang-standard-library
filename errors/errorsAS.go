package main

import (
	"errors"
	"fmt"
)

type Peserta struct {
	Nama string
	Umur int
}

type CustomErrPeserta struct {
	Umur        int
	UmurMinimal int
}

func (e *CustomErrPeserta) Error() string {
	return "Umur tidak memenuhi syarat daftar"
}

func (p *Peserta) Daftar() error {
	if p.Umur < 17 {
		return &CustomErrPeserta{
			Umur:        p.Umur,
			UmurMinimal: 17,
		}
	}
	return nil
}

func main() {
	peserta := Peserta{
		Nama: "Budi",
		Umur: 2,
	}

	var customErr *CustomErrPeserta
	err := peserta.Daftar()
	if errors.As(err, &customErr) {
		fmt.Println(err)
		fmt.Println("Umur peserta :", customErr.Umur)
		fmt.Println("Minimal umur :", customErr.UmurMinimal)
	}else{
		fmt.Println("Sukses selalu yah anda bisa daftar")
	}
}
