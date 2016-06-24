package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	ret, err := os.Getwd()
	if err != nil {
		fmt.Println("get work directory error:", err)
	}
	cmd := exec.Command(ret + "/workdirtest")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = "./../../"

	err = cmd.Run()
	if err != nil {
		fmt.Println("get work directory error:", err)
	}
}
