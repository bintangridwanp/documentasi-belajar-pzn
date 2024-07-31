package main

import "fmt"

func main(){

	//operasi boolean int

	var nilaiUas = 90
	var absensi = 80

	var  lulusUas bool = nilaiUas > 80
	var  lulusAbsensi bool = absensi > 75

	var lulus bool = lulusAbsensi && lulusUas

	fmt.Println(lulus)

	//cara singkatnya

	fmt.Println(nilaiUas > 80 && absensi > 75)
}