package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDurasiTidakValid = errors.New("durasi booking harus antara 30 menit sampai 4 jam")

type Booking struct {
	NamaRuangan string
	Mulai       time.Time
	Selesai     time.Time
}

func (b Booking) HitungDurasi() time.Duration {
	return b.Selesai.Sub(b.Mulai)
}

func (b Booking) Validasi() error {
	durasi := b.HitungDurasi()
	if durasi < 30*time.Minute || durasi > 4*time.Hour {
		return ErrDurasiTidakValid
	}
	return nil
}

func (b Booking) InfoJadwal() string {
	return fmt.Sprintf("%s: %s - %s (durasi: %s)",
		b.NamaRuangan,
		b.Mulai.Format("15:04"),
		b.Selesai.Format("15:04"),
		b.HitungDurasi())
}

func main() {
	mulai := time.Date(2026, 7, 12, 9, 0, 0, 0, time.Local)

	booking1 := Booking{
		NamaRuangan: "Ruang A",
		Mulai:       mulai,
		Selesai:     mulai.Add(90 * time.Minute),
	}

	booking2 := Booking{
		NamaRuangan: "Ruang B",
		Mulai:       mulai,
		Selesai:     mulai.Add(10 * time.Minute),
	}

	for _, b := range []Booking{booking1, booking2} {
		err := b.Validasi()
		if err != nil {
			fmt.Println(b.NamaRuangan, "- Error:", err)
			continue
		}
		fmt.Println(b.InfoJadwal())
	}
}