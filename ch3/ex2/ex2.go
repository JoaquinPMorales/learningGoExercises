package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	// Converts the string to an slice of runes
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}
