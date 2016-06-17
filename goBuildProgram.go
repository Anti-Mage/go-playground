package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var sourceStr string

func init() {
	sourceStr = `
package main

import (
	"flag"
	"fmt"
	"time"
)
	
	func main() {
		f := flag.String("f", "mockProgram.conf", "config file")
		var a, b int
		fmt.Println("flag:", *f)
		for {
			_, err := fmt.Scanf("%d %d", &a, &b)
			if err != nil {
				break
			}
	
			fmt.Printf("%d + %d = %d\n", a, b, a+b)
			time.Sleep(time.Second)
		}
	}

	`
}

func main() {
	sourceFile := "mockProgramTmp.go"
	f, err := os.OpenFile(sourceFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	_, err = f.Write([]byte(sourceStr))
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	f.Close()

	cmd := exec.Command("go", "fmt", sourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		fmt.Println("start cmd fmt error:", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("start cmd fmt error:", err)
	}

	time.Sleep(time.Second)

	cmd = exec.Command("go", "build", sourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		fmt.Println("line 75, start cmd build error:", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("line 80, start cmd build error:", err)
	}

}
