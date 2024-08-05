package main

import "fmt"

func main() {

	// Mendefinisikan variabel untuk nama
	namaAwal := "luffy"
	namaTengah := "D"
	namaAkhir := "monkey"

	// Mencetak string dengan variabel nama menggunakan fmt.Println
	// Menyertakan tanda kutip tunggal di sekitar nama
	fmt.Println("Hello '", namaAwal, namaTengah, namaAkhir, "'")

	// Mencetak string dengan format yang ditentukan menggunakan fmt.Printf
	// Menggunakan placeholder %s untuk mencetak setiap variabel nama
	fmt.Printf("Hello '%s %s %s'\n", namaAwal, namaTengah, namaAkhir)

}
