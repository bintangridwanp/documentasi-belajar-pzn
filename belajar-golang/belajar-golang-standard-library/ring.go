package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5) // Membuat ring baru dengan ukuran 5

	// Mengisi ring dengan nilai "Value 0" hingga "Value 4"
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	//data.Value = "Value 1"
	//
	//data = data.Next()
	//data.Value = "Value 2"
	//
	//data = data.Next()
	//data.Value = "Value 3"
	//
	//data = data.Next()
	//data.Value = "Value 4"
	//
	//data = data.Next()
	//data.Value = "Value 5"

	// Mengiterasi dan menampilkan setiap nilai dalam ring
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
