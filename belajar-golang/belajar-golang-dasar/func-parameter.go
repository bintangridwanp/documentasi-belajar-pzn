package main

import "fmt"

func sayHelloTo(firtsName string, middleName string, latsName string){
	fmt.Println("hello", firtsName, middleName, latsName)
}

func main() {
	sayHelloTo("Bintang", "Ridwan", "Pribadi")
}
