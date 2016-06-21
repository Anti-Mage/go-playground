package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world:", i)
	}
}
