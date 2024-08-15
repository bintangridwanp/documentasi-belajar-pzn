package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second     // Durasi 100 detik
	var duration2 time.Duration = 10 * time.Millisecond // Durasi 10 milidetik
	var duration3 time.Duration = duration1 - duration2 // Mengurangi duration1 dengan duration2

	fmt.Println(duration3)                   // Menampilkan durasi hasil pengurangan
	fmt.Printf("duration : %d\n", duration3) // Menampilkan durasi dalam format integer (nanodetik)
}
