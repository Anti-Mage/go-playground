package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./test.conf"
	/*f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("error in line 12:", err)
	}*/

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error in line 12:", err)
	}

	fmt.Println(string(dat))
}
