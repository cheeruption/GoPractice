package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
	"github.com/cheeruption/firstapp/pkg/Praktik"
)

var i, j int = 1, 5
var (
	y int = 44
	u int = 33
)

var h int = 45

const h51 = 666
const (
	big   = 1 << 100
	small = big >> 99
)
const arr1 = "yrots"

func main() {
	k, str := 68, "test_string"
	defer fmt.Println("\nhere is Deferred function - hi there", rand.Intn(99))
	fmt.Println("This is pi", math.Pi)
	fmt.Println(add(5, 6), add2(5, 6))
	fmt.Println(swap("this", "that"))
	fmt.Println(swap2("this", "that"))
	fmt.Println("Объявление переменных - ", i, j, "\nShort declaration - ", k, str, "\nFactored declaration - ", y, u)
	g := float64(h) + 0.35
	fmt.Println("Конвертация типов переменных -", g, "\nAnd that is a const -", h51)

	fmt.Println(needInt(small))
	//fmt.Println(needInt(big)) слишком большой int - программа крашится
	fmt.Println(needFloat(small))
	//fmt.Println(needFloat(big)) а здесь всё будет ок
	for i := 0; i < len(arr1); i++ {
		if i == 0 {
			defer fmt.Println("\n Выше видим как записывается IF и цикл FOR.")
		}
		defer fmt.Printf("     "+"%c", arr1[i])

	}

	fmt.Println("\nВместо while используем сокращенный for")
	for y98 := 0; y98 < 5; {
		fmt.Println("check №", (y98 + 1))
		y98++
	}
	fmt.Println("Испытываем Switch и смотрим текущую Ось = ", whatOS())
	fmt.Printf("Today is - %s", time.Now().Weekday())

	//if-then-else SWITCH style
	t := time.Now().Hour()
	fmt.Println("\nLet's check the time:")
	switch {
	case t < 12:
		fmt.Println("wow! morning")
	case t < 17:
		fmt.Println("its evening soon")
	case t < 23:
		fmt.Println("Dawn of the new day")
	default:
		fmt.Println("its reaally late")
	}
	point1 := &t
	fmt.Println("Что насчет pointer'ов?", &t, *point1, *&t == *point1) // dereferencing
	fmt.Println("А вот и структуры подъехали", smth{56, "horses"})
	mare := smth{3, "kon"}
	fmt.Println(mare.that)
	arr2 := [3]int{5, 6, 7}
	arr3 := arr2
	arr4 := &arr2
	arr4[0] = 99
	arr3[0] = 77
	sl := arr2[1:2]
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Arrays and slices are coming...", arr2, arr2[0], arr3, arr4, sl, "\n", s)
	//
	// 13/27 глава по спискам
	//
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

type smth struct {
	this int
	that string
}

fmt.Println(praktik())