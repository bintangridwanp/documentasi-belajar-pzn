package main

import (
	"fmt"
	"reflect"
)

// Struct Sample dengan tag pada field Name
type Sample struct {
	Name string `required:"true" max:"10"`
}

// Struct Person dengan tag pada field Name, Address, dan Email
type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

// Fungsi untuk membaca dan menampilkan informasi field dalam struct
func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name()) // Menampilkan nama tipe dari struct
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type) // Menampilkan nama dan tipe dari field
		fmt.Println(structField.Tag.Get("required"))                 // Menampilkan nilai tag "required"
		fmt.Println(structField.Tag.Get("max"))                      // Menampilkan nilai tag "max"
	}
}

// Fungsi untuk memvalidasi apakah field yang ditag "required" memiliki nilai
func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" { // Memeriksa apakah tag "required" bernilai "true"
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != "" // Memastikan field yang ditag "required" tidak kosong
			if result == false {
				return result // Jika ada field yang kosong, hasil validasi adalah false
			}
		}
	}
	return result // Jika semua field valid, hasilnya adalah true
}

func main() {
	readField(Sample{"Eko"})         // Membaca field dari struct Sample
	readField(Person{"Eko", "", ""}) // Membaca field dari struct Person

	person := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
	}
	fmt.Println(IsValid(person)) // Memvalidasi struct Person dan menampilkan hasilnya
}
