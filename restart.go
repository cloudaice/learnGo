package main

import (
	"fmt"
	"time"
)

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			time.Sleep(time.Second)
			go foo()
		}
	}()

	panic("make failed")
}

func bar(channel chan string) {
	defer func() {
		if r := recover(); r != nil {
			channel <- "failed"
		}
	}()
	panic("fail")
}
func main() {
	go foo()
	timeout := time.After(5 * time.Second)
	channel := make(chan string)
	go bar(channel)
	for {
		select {
		case info := <-channel:
			fmt.Println("channel", info)
			time.Sleep(time.Second)
			go bar(channel)
		case <-timeout:
			return
		}
	}
}
