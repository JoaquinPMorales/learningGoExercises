package main

import "fmt"

func UpdateSlice(arg1 []string, arg2 string) {
	arg1[len(arg1)-1] = arg2
	fmt.Println("From UpdateSlice...: ", arg1)
}

// This does not work because you cannot extend a slice from a functio
// Since the len and cap are copies, not the real value.
func GrowSlice(arg1 []string, arg2 string) {
	arg1 = append(arg1, arg2)
	fmt.Println("From GrowSlice...: ", arg1)
}

func main() {
	stringSlice := []string{"Hello", "Goodbye", "Good night"}
	fmt.Println(stringSlice)
	UpdateSlice(stringSlice, "Hola")
	fmt.Println(stringSlice)
	fmt.Println(stringSlice)
	GrowSlice(stringSlice, "Hola")
	fmt.Println(stringSlice)
}
