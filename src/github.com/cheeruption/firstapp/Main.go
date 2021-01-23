package main

import (
	"fmt"
	"math"
	"math/rand"
)

var i, j int = 1, 5
var (
	y int = 44
	u int = 33
)

func main() {
	k, str := 68, "test_string"
	fmt.Println("hi there", rand.Intn(99))
	fmt.Println("This is pi", math.Pi)
	fmt.Println(add(5, 6), add2(5, 6))
	fmt.Println(swap("this", "that"))
	fmt.Println(swap2("this", "that"))
	fmt.Println("Объявление переменных - ", i, j, "\nShort declaration - ", k, str, "\nFactored declaration - ", y, u)
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
