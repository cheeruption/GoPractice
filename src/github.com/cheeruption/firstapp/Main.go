package main

import (
	"fmt"
	"math"
	"math/rand"
)

var i, j int = 1, 5

func main() {
	k := "test_string"
	fmt.Println("hi there", rand.Intn(99))
	fmt.Println("This is pi", math.Pi)
	fmt.Println(add(5, 6), add2(5, 6))
	fmt.Println(swap("this", "that"))
	fmt.Println(swap2("this", "that"))
	fmt.Println("Объявления функции через var - ", i, j, "\n"+k)
}
func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func swap(x, y string) string {
	return (y + " and " + x)
}
func swap2(x, y string) (y1, x1 string) {
	y1 = "try"
	x1 = "this"
	return
}
