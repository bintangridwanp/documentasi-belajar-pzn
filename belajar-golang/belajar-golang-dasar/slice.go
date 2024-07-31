package main

import "fmt"

func main(){

	// tipe data slice memiliki 3 data yang penting
	// 1. pointer (adalah petunjuk data pertama di array para slice)
	// 2. lenght (panjang dari slice)
	// 3. capacity (kapasitas dari slice, dimana lenght tidak boleh lebih dari capacity)

	bulan := [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
	}

	var bulan1 = bulan[4:7]
	fmt.Println(bulan1)

	//untuk mendapatkan panjang
	fmt.Println(len(bulan1))

	//untuk mendapatkan kapasistas
	fmt.Println(cap(bulan1))



	bulan[4]= "Mei di ubah"
	fmt.Println(bulan1)

	bulan2 := bulan[10:]
	fmt.Println(bulan2)

	//membuat slice baru dengan menambahkan data keposisi terkhir slice, jika kapasiitas slice penuh maka akan membuat array baru,
	
	bulan3 := append(bulan2, "bintang")
	fmt.Println(bulan3)

	//membuat slice baru

	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "Bintang"
	sliceBaru[1] = "Ridwan"

	fmt.Println(sliceBaru)

	//menyalin slice dari source ke destination

	copySlice := make([]string, len(sliceBaru), cap(sliceBaru))
	copy(copySlice, sliceBaru)

	fmt.Println(copySlice)

	//perbedaan slice dan  array
	iniSlice := []int { 1, 2, 3}
	iniArray := [...]int {1, 2, 3}

	fmt.Println(iniSlice)
	fmt.Println(iniArray)

}