package main

import "fmt"

func main(){

	// dalam membuat variable di golang, dalam membuat 1 variable harus menggunaka tipe data yang sama.
	// dalam membuat variable, tidak wajib untuk menyebutkan tipe data.
	// di golang menggunaka kata kunci var itu tdk wajib, bisa di ganti dgn := (nama_variable := isi variable), ini hanya berlaku utk deklarasi pertama.
	// nilai variable harus digunakan, jika variable tdk di gunakan, maka kode akan error.

	var nama string

	//variable nama

	nama = "bintang ridwan"
	fmt.Println(nama)

	nama = "bintang pribadi"
	fmt.Println(nama)

	//variable nama teman

	var namaTeman = "pribadi"
	fmt.Println(namaTeman)

	//variable angka

	 angka := 2023
	fmt.Println(angka)

	//variable nama guru

	namaGuru := "Deni Saepudin"
	fmt.Println(namaGuru)

	//Deklarasi multiple Variable (membuat variable secara sekaligus)

	var (
		namaAwal = "bintang"
		namaTengah = "ridwan"
		namaAkhir = "pribadi"
	)
	
	fmt.Println(namaAwal)
	fmt.Println(namaTengah)
	fmt.Println(namaAkhir)
}