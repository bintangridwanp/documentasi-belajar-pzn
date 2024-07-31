package main

import "fmt"

//mengembalikan nilai func lebih dari satu

func getFullName() (string, string, string){
	return "Bintang", "Ridwan", "Pribadi"
}

//nilai value yang di kembalikan harus di ambil, jika ada data yang tidak di butuhkan bisa mengguakan garis bawah (_)  

func main()  {
	firtsName, _, lastName := getFullName()
	fmt.Println(firtsName, lastName)
}