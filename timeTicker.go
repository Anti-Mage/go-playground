package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("hello world")
			}
		}
	}()
	time.Sleep(time.Millisecond * 16000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
