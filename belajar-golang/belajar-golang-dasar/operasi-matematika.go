package main

import "fmt"

func main(){

	//operasi matematika

	var result = 10 + 10
	fmt.Println(result)

	a := 5
	b := 5
	c := a * b

	fmt.Println(c)

	//augmented assiggnment

	var i = 10
	i += 10

	fmt.Println(i)

	var x = 50
	x *= 2

	fmt.Println(x)

	//unary operator

	i++ 
	fmt.Println(i)

	x++
	fmt.Println(x)
	
	var negative = -10
	var positive = +10

	fmt.Println(negative)
	fmt.Println(positive)

}