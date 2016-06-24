package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//dont do this, see above edit

type address struct {
	First_address  string `json:"first_address"`
	Second_address string `json:"second_address"`
}

type example struct {
	Nodeid    string    `json:"node_id"`
	Max_procs int       `json:"max_procs"`
	Addrs     []address `json:"address"`
}

func createExample() *example {
	var ret example
	var addr1 address
	addr1.First_address = "1.2.3.4"
	addr1.Second_address = "5.6.7.8"
	var addr2 address
	addr2.First_address = "11.22.33.44"
	addr2.Second_address = "55.66.77.88"
	ret.Nodeid = "ahksdh12-asd1232-kasdjfl"
	ret.Max_procs = 8
	ret.Addrs = append(ret.Addrs, addr1)
	ret.Addrs = append(ret.Addrs, addr2)

	return &ret
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "	")
	return out.Bytes(), err
}

func prettyPrint() {
	ret := createExample()

	jsonInfo, err := json.Marshal(*ret)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	b, err := prettyprint(jsonInfo)
	if err != nil {
		fmt.Println("line 56 file error:", err)
		return
	}

	fmt.Printf("%s", b)
}

func main() {
	f, err := os.Open("json.conf")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	jsonInfo, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	b, err := prettyprint(jsonInfo)
	fmt.Printf("%s", b)

	fmt.Println("\n\n\n")
	prettyPrint()
}
