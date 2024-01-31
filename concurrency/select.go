package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string, 3)
	data := []string{"a", "b", "c"}


	for _, v := range data {
		select {
		case myChan <- v:
		}
	}


	close(myChan)

	for v := range myChan {
		fmt.Println(v)
	}

}
