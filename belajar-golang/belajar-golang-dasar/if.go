package main

import "fmt"

func main() {

	//if adalah salah satu kode yang digunakan utk percabangan
	//percabangan artinya, kita bisa mengeksekusi kode program tertentu ketika suatu kondisi sudah terpenuhi 	

	nama := "Bintang"

	if nama == "Bintang" {
		fmt.Println("Hello Ganteng")
	} else if nama == "Deni" {
		fmt.Println("Hallo Deni")
	} else { 
		fmt.Println("HI, Bukan Deni or Bintang")
	}

	// kode program if short statement

	if lenght := len(nama); lenght < 5 {
		fmt.Println("nama sudah benar")
	} else {
		fmt.Println("nama terlalu panjang")
	}
}