package main

import "fmt"

func namaLengkap() (namaAwal, namaTengah, namaAkhir string){
	namaAwal = "Bintang"
	namaTengah = "Pribadi"
	namaAkhir = "Pribadi"

	return namaAwal, namaTengah, namaAkhir

}

func main() {
	a, b, c := namaLengkap()
	fmt.Println(a, b, c)
}
