package main

import "fmt"

//mengembalikan nila dari func 1 ke func 2

func sayHello(nama string) string {
	hello := "hallo " + nama
	return hello
}

func main() {
	result := sayHello("bintang")
	fmt.Println(result)
}