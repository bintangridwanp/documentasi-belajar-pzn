package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Mengubah string "true" menjadi nilai boolean
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error()) // Menampilkan pesan error jika terjadi masalah
	} else {
		fmt.Println(result) // Menampilkan hasil konversi
	}

	// Mengubah string "1000" menjadi nilai integer
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error()) // Menampilkan pesan error jika terjadi masalah
	} else {
		fmt.Println(resultInt) // Menampilkan hasil konversi
	}

	// Mengubah integer 999 menjadi string biner
	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary) // Menampilkan hasil konversi biner

	// Mengubah integer 999 menjadi string
	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt) // Menampilkan hasil konversi string
}
