package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func NewProcess(args string) (cmd *exec.Cmd, err error) {

	cmd = exec.Command(args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Pdeathsig: syscall.SIGKILL, // only worked on linux
		Setsid: true,
	}

	return
}

func main() {
	binFile := "./go-daemon-test"
	cmd, err := NewProcess(binFile)
	if err != nil {
		fmt.Println("line 21:error:", err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("line 32:error:", err)
	}

	fmt.Println("pid:", cmd.Process.Pid)

	time.Sleep(time.Duration(1000) * time.Second)
}
