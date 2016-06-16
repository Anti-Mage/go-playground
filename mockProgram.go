package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	flag := flag.String("f", "mockProgram.conf", "config file")
	var a, b int
	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}

		fmt.Printf("%d + %d = %d\n", a, b, a+b)
		time.Sleep(time.Second)
	}
}
