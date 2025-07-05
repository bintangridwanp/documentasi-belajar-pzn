package main

import "fmt"

func main() {

	// Mendeklarasikan variabel untuk nama
	namaAwal := "luffy"
	namaTengah := "D"
	namaAkhir := "monkey"

	// Mencetak string dengan format yang ditentukan menggunakan fmt.Println
	fmt.Println("Hello '", namaAwal, namaTengah, namaAkhir, "'")

	// Mencetak string dengan format yang ditentukan menggunakan fmt.Printf
	fmt.Printf("Hello '%s %s %s'\n", namaAwal, namaTengah, namaAkhir)

}
