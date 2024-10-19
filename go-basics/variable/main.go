package main

import (
	"fmt"
)

func main() {
	age := []int{
		1,
		2,
		3,
		4,
		5,
	}
	age = append(age, 10)
	for _, value := range age {
		fmt.Printf("this is my value %d \n", value)
	}
}
