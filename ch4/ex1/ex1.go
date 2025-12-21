package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mySlice := make([]int, 0, 100)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, rand.Intn(100))
	}
	fmt.Println(mySlice)
}
