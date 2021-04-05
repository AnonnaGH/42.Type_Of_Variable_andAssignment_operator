package main

import "fmt"

func main() {

	//declare variable

	var str string
	var n, m int
	var mn float32

	// assign value

	str = "Hellow Anonna!"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("value of str=", str)
	fmt.Println("value of n=", n)
	fmt.Println("value of m=", m)
	fmt.Println("value of mn=", mn)

	//declare and assign value to variables

	var city string = "Dhaka"
	var x int = 200

	fmt.Println("value of city=", city)
	fmt.Println("value of x=", x)

	//declare variable with defining its type

	country := "Faridpur"
	val := 20
	fmt.Println("value of country=", country)
	fmt.Println("value of val=", val)

	//define multiple variables

	var (
		name  string
		email string
		age   int
	)

	name = "Anonna"
	email = "john@email.com"
	age = 24
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
