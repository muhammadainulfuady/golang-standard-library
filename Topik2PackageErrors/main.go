package main

import (
	"errors"
	"fmt"
)

var (
	ErrSaldoTidakCukup = errors.New("Saldo anda tidak cukup untuk ditarik")
)

type AtmMini struct {
	NamaBank    string
	Saldo       int
	NamaNasabah string
}

type AtmMiniErr struct {
	Message string
	Saldo   int
	Ditarik int
}

func (a *AtmMiniErr) Error() string {
	return a.Message
}

func narikUang(a *AtmMini, jumlah int) error {
	if a.Saldo < jumlah {
		return &AtmMiniErr{
			Message: "Saldo tidak cukup",
			Ditarik: jumlah,
			Saldo:   a.Saldo,
		}
	}
	a.Saldo -= jumlah
	return nil
}

func main() {
	atm := &AtmMini{
		NamaBank:    "BRI",
		Saldo:       20000,
		NamaNasabah: "Nama Nasabah Terbaik",
	}

	jumlahTarik := 190000
	err := narikUang(atm, jumlahTarik)
	var en *AtmMiniErr
	if err != nil {
		if errors.As(err, &en) {
			fmt.Printf("Anda menarik %d sedangkan saldo %d : %s\n", en.Ditarik, en.Saldo, en.Message)
			return
		}
	}
	fmt.Printf("Anda menarik %d sisa saldo %d :\n", jumlahTarik, atm.Saldo)
}
