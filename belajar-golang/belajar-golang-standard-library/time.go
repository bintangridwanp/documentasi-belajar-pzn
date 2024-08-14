package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now() // Mengambil waktu saat ini
	fmt.Println(now)               // Menampilkan waktu saat ini

	// Membuat waktu UTC spesifik: 17 Agustus 2009, 00:00:00
	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)         // Menampilkan waktu dalam UTC
	fmt.Println(utc.Local()) // Mengonversi waktu UTC ke waktu lokal dan menampilkannya

	formatter := "2006-01-02 15:04:05" // Format untuk parsing dan format waktu

	value := "2020-10-10 10:10:10"
	// value := "ASAL" // Contoh value yang salah format
	valueTime, err := time.Parse(formatter, value) // Parsing string ke dalam tipe time.Time sesuai format yang diberikan
	if err != nil {
		fmt.Println("Error", err.Error()) // Menampilkan error jika parsing gagal
	} else {
		fmt.Println(valueTime) // Menampilkan waktu hasil parsing
	}

	// Menampilkan tahun, bulan, hari, dan jam dari waktu yang sudah diparsing
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
