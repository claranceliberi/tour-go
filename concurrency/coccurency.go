package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string)
	anotherChan := make(chan string)

	go func() {
		myChan <- "Hello"
	}()

	go func() {
		anotherChan <- "World"
	}()

	select {
		case msg1 := <-myChan:
			fmt.Println(msg1)
		case msg2 := <-anotherChan:
			fmt.Println(msg2)
	}
}
