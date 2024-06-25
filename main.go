package main

import "fmt"

func main() {
	var lastName string

	// var (
	// 	firstName = "Reda"
	// 	age       = 20
	// )
	country := "Morocco"
	lastName = "Bouzad"

	var (
		customerName = "Adam"
		customerSold = 20
	)

	fmt.Println("Hello")
	fmt.Println("Reda")
	fmt.Println(lastName)
	fmt.Println(country)
	// fmt.Println("Customer: " + customerName + " has " + customerSold. + "$ in his bank account")
	fmt.Printf("Customer has %s has %d$ in his bank account \n", customerName, customerSold)
}
