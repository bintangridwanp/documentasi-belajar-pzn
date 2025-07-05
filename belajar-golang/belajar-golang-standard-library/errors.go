package main

import (
	"errors"
	"fmt"
)

// Mendeklarasikan dua error khusus untuk digunakan dalam fungsi DapatkanId
var (
	MenemukanError      = errors.New("validation error")
	TidakMenemukanError = errors.New("tidak menemukan error")
)

// DapatkanId adalah fungsi yang menerima id sebagai parameter dan mengembalikan error
func DapatkanId(id string) error {
	if id == "" {
		return MenemukanError
	}
	if id == "OnePiece" {
		return TidakMenemukanError
	}

	return nil
}

// main adalah fungsi utama yang memanggil DapatkanId dan menangani error yang dikembalikan
func main() {
	rere := DapatkanId("OnePiece")

	if errors.Is(rere, MenemukanError) {
		fmt.Println("menemukan error")
	} else if errors.Is(rere, TidakMenemukanError) {
		fmt.Println("tidak menemukan error")
	} else {
		fmt.Println("apa itu error")
	}
}
