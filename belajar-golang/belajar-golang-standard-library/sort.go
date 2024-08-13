package main

import (
	"fmt"
	"sort"
)

// Struct User dengan atribut Name dan Age
type User struct {
	Name string
	Age  int
}

// UserSlice adalah slice dari User
type UserSlice []User

// Len mengembalikan panjang dari UserSlice
func (s UserSlice) Len() int {
	return len(s)
}

// Less membandingkan umur dua elemen User dan mengembalikan true jika umur elemen pertama lebih kecil
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Swap menukar posisi dua elemen dalam UserSlice
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"luffy", 30},
		{"zoro", 35},
		{"sanji", 25},
		{"nami", 20},
	}

	// Mengurutkan slice users berdasarkan umur (Age) secara ascending
	sort.Sort(UserSlice(users))

	fmt.Println(users) // Menampilkan hasil urutan
}
