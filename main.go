package main

import "fmt"

var (
	temp = 33
	wind = 2.2
	city = "Marrakech"
)

var (
	customerName    = "Adam"
	customerBalance = 20
	isMarried       = false
	isDating        = true
	isSingle        = false
)

var (
	accountBalace = -100
	accountCredit = -200
)

var (
	hasGas         = true
	hasKeyIgnition = true
)

var ()

func main() {

	// Declaring Variables
	country := "Morocco"
	var lastName string
	lastName = "Ahmdan"
	fmt.Println(lastName)
	fmt.Println(country)
	fmt.Printf("Customer %s has %d$ in his bank account \n", customerName, customerBalance)
	fmt.Printf("today the temperature in %s is %d, and the wind is %0.1f \n", city, temp, wind)

	// If else statements
	if isMarried {
		fmt.Println("You recieve a bonus of 50$")
	} else if isDating {
		fmt.Println("You recieve a bonus of 25$")
	} else {
		fmt.Println("You must be single, you don't recieve a bonus")
	}

	if hasGas && hasKeyIgnition {
		fmt.Println("Can drive car")
	}
}
