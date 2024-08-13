package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New() // Membuat list baru (doubly linked list)

	// Menambahkan elemen ke dalam list
	data.PushBack("luffy")
	data.PushBack("D")
	data.PushBack("monkey")

	// Mengambil elemen pertama dari list
	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Menampilkan elemen pertama ("luffy")

	// Mengambil dan menampilkan elemen selanjutnya ("D")
	next := head.Next()
	fmt.Println(next.Value)

	// Mengambil dan menampilkan elemen selanjutnya ("monkey")
	next = next.Next()
	fmt.Println(next.Value)

	// Iterasi dan menampilkan semua elemen dalam list
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
