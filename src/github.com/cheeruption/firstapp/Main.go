package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var i, j int = 1, 5
var (
	y int = 44
	u int = 33
)

var h int = 45

const h51 = 666
const (
	Big   = 1 << 100
	Small = Big >> 99
)
const arr1 = "story"

func main() {
	k, str := 68, "test_string"
	fmt.Println("hi there", rand.Intn(99))
	fmt.Println("This is pi", math.Pi)
	fmt.Println(add(5, 6), add2(5, 6))
	fmt.Println(swap("this", "that"))
	fmt.Println(swap2("this", "that"))
	fmt.Println("Объявление переменных - ", i, j, "\nShort declaration - ", k, str, "\nFactored declaration - ", y, u)
	g := float64(h) + 0.35
	fmt.Println("Конвертация типов переменных -", g, "\nAnd that is a const -", h51)

	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big)) слишком большой int - программа крашится
	fmt.Println(needFloat(Small))
	//fmt.Println(needFloat(Big)) а здесь всё будет ок
	for i := 0; i < len(arr1); i++ {
		if i == 0 {
			fmt.Println("Тут смотрим как записывается IF и цикл FOR:")
		}
		fmt.Printf("     "+"%c", arr1[i])

	}

	fmt.Println("\nВместо while используем сокращенный for")
	for y98 := 0; y98 < 5; {
		fmt.Println("check №", (y98 + 1))
		y98++
	}
	fmt.Println("Испытываем Switch и смотрим текущую Ось = ", whatOS())
	fmt.Printf("Today is - %s", time.Now().Weekday())
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
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func whatOS() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		os = "OS X"
	case "linux":
		os = "Linux"
	default:
		os = "Windows"
	}
	return os
}
