package main

import (
	"fmt"
	"os"
	//"path"
)

func main() {
	ret, err := os.Getwd()
	if err != nil {
		fmt.Println("get work directory error:", err)
	}

	fmt.Println("work directory:", ret)
}
