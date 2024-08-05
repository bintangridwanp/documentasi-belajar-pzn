package main

import (
	"errors"
	"fmt"
)

var (
	// Mendefinisikan dua variabel error dengan pesan kesalahan yang berbeda
	MenemukanError      = errors.New("validation error")
	TidakMenemukanError = errors.New("tidak menemukan error")
)

// Fungsi DapatkanId memeriksa nilai id dan mengembalikan error berdasarkan kondisi tertentu
func DapatkanId(id string) error {
	// Jika id kosong, kembalikan MenemukanError
	if id == "" {
		return MenemukanError
	}

	// Jika id adalah "OnePiece", kembalikan TidakMenemukanError
	if id == "OnePiece" {
		return TidakMenemukanError
	}

	// Jika id valid dan tidak memenuhi kondisi di atas, kembalikan nil (tidak ada error)
	return nil
}

func main() {
	// Memanggil fungsi DapatkanId dengan id kosong dan menyimpan hasilnya di variabel rere
	rere := DapatkanId("")

	// Memeriksa jenis error yang dikembalikan oleh DapatkanId dan mencetak pesan yang sesuai
	if errors.Is(rere, MenemukanError) {
		fmt.Println("menemukan error")
	} else if errors.Is(rere, TidakMenemukanError) {
		fmt.Println("tidak menemukan error")
	} else {
		fmt.Println("apa itu error")
	}
}
