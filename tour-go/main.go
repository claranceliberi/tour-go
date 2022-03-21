package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Println(add(42, 13))

	a,b := swap("3", "4")
	fmt.Println("Swap 3,4 =>",a,b)

	fmt.Println("Naked Return")
	x,y := split(17)
	fmt.Println(x,y)

	var i, j int = 1, 2 // bariables with initializers
	fmt.Println("i is", i, "j is ", j, "and c is", c, "and python is", python, "and java is", java)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

