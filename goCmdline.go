package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func test1() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Run(os.Args)
}

func main() {
	test1()
}
