package main

import (
	"fmt"
	"os"
)

func main() {
	dir1 := "./test/tmp/"
	dir2 := "./test2/tmp2/"
	//err := os.MkdirAll(dir1, 0777)
	/*if err != nil {
		fmt.Println(err)
	}*/

	err := os.Rename(dir1, dir2)
	if err != nil {
		fmt.Println("1", err)
	}

	/*err = os.MkdirAll(dir2, 0777)
	if err != nil {
		fmt.Println(err)
	}*/

	err = os.Rename(dir1, dir2)
	if err != nil {
		fmt.Println("2", err)
	}

}
