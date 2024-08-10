package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  // Pembulatan ke atas dari 1.40 menjadi 2
	fmt.Println(math.Floor(1.60)) // Pembulatan ke bawah dari 1.60 menjadi 1
	fmt.Println(math.Round(1.60)) // Pembulatan ke nilai terdekat dari 1.60 menjadi 2
	fmt.Println(math.Max(10, 11)) // Mengambil nilai maksimum antara 10 dan 11
	fmt.Println(math.Min(10, 11)) // Mengambil nilai minimum antara 10 dan 11
}
