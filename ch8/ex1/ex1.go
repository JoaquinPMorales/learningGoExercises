package main

import "fmt"

type IntOrFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T IntOrFloat](t T) T {
	return t * 2
}

func main() {
	fmt.Println("The value is: ", Doubler(2))
	fmt.Println("The value is: ", Doubler(5.0))
}
