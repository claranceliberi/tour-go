package main

// import (
// 	"fmt"
// 	"math/rand"
// )

func add(x int, y int) int {
	return x + y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

// func main(){
	
	// fmt.Println("My favorite number is", rand.Intn(1000))
	// fmt.Println(add(42, 13))

	// a,b := swap("3", "4")
	// fmt.Println("Swap 3,4 =>",a,b)

	// fmt.Println("Naked Return")
	// x,y := split(17)
	// fmt.Println(x,y)

	// var i, j int = 1, 2 // bariables with initializers
	// fmt.Println("i is", i, "j is ", j, "and c is", c, "and python is", python, "and java is", java)
// }