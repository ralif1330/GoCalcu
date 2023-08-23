package main

import (
	"fmt"

	gocalculatorfunction "github.com/abiyyuaqzal/go-calculator-function"
)

func main() {

	var a, b float64
	var operation string

	fmt.Println("Golang Calculator")

	fmt.Println("Enter your first number : ")
	fmt.Scanln(&a)

	fmt.Println("Enter ur second number : ")
	fmt.Scanln(&b)

	fmt.Println("Enter ur operation : ")
	fmt.Scanln(&operation)

	result, err := gocalculatorfunction.Calcualate(a, b, operation)

	if err == nil {
		fmt.Printf("%.3f %s %.3f = %.3f", a, operation, b, result)
	} else {
		fmt.Println(err.Error())
	}

}
