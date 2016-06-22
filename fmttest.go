package main

import (
	"fmt"
)

func main() {
	a := 123
	fmt.Printf("%-10d\n", a)
	fmt.Printf("%-10s\n", "1234")
	fmt.Printf("%-10s, %s", "1236", "45690")
}
