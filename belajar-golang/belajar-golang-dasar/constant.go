package main

import "fmt"

func main(){

	// contant adalah variable yang nilainya tdk bisa di rubah lagi setelah pertama kali di beri nilai.
	// tidak wajin mendlekarasikan nama tipe data
	// nilai conts tdk wajib di gunakan, jika const tdk di pakai maka kode tdk akan error (beda dgn variable)

	const namaAwal string = "bintang"
	const namaAkhir = "ridwan"

	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)

	// Deklarsi mutiple Constant

	const (
		namaTeman = "Bambang"
		namaGuru = "Tamtoyo"
	)

	fmt.Println(namaTeman)
	fmt.Println(namaGuru)

}