package main

import "fmt"

func main() {
	fmt.Print("Enter distance in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input * 0.3048)

	fmt.Println("The distance in metres is", output)
}
