package main

import "fmt"

func main() {
	var greetings []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	// First two values
	var firstSlice = greetings[:2]
	// Second, third and fourth values
	var secondSlice = greetings[1:4]
	// Last two ones, fourth and fifth values
	var thirdSlice = greetings[3:]
	fmt.Println(greetings)
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
}
