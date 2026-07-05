package main

import (
	"errors"
	"fmt"
)

// sentinel error - buat kasus yang cukup "cek identitas" doang
var ErrKartuDiblokir = errors.New("kartu diblokir")

// custom error struct - buat kasus yang butuh data tambahan
type SaldoKurangError struct {
	SaldoTersedia int
	JumlahDiminta int
}

func (e *SaldoKurangError) Error() string {
	return fmt.Sprintf("saldo kurang: tersedia %d, diminta %d", e.SaldoTersedia, e.JumlahDiminta)
}

func Tarik(saldo int, jumlah int, kartuDiblokir bool) error {
	if kartuDiblokir {
		return ErrKartuDiblokir
	}
	if jumlah > saldo {
		return &SaldoKurangError{SaldoTersedia: saldo, JumlahDiminta: jumlah}
	}
	return nil
}
func main() {
	testTarik(100000, 50000, false)    // harusnya: Penarikan berhasil
	testTarik(100000, 50000, true)     // harusnya: Kartu ATM anda di blokir
	testTarik(5000000, 9000000, false) // harusnya: Saldo kurang, dst
}

func testTarik(saldo, jumlah int, kartuDiblokir bool) {
	var target *SaldoKurangError
	err := Tarik(saldo, jumlah, kartuDiblokir)
	if err != nil {
		if errors.Is(err, ErrKartuDiblokir) {
			fmt.Println("Kartu ATM anda di blokir")
			return
		}
		if errors.As(err, &target) {
			fmt.Println("Saldo kurang, tersedia:", target.SaldoTersedia, "diminta:", target.JumlahDiminta)
			return
		}
	}
	fmt.Println("Penarikan berhasil")
	fmt.Println("---")
}
