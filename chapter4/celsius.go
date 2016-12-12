package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := ((input - 32) * 5 / 9)

	fmt.Println("The temperature in Celsius is", output)
}
