package main

import (
	"fmt"
	"strings"
)

func sayG(n string) {
	fmt.Println("hello", n)
}
func sayB(n string) {
	fmt.Println("bye", n)
}

func sayH(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func intials(n string) (string, string) {
	names := strings.Split(n, " ")
	first := names[0]
	last := names[1]
	return string(first[0]), string(last[0])
}

func main() {
	// fmt.Println("Hello, World!")
	// var nameOne string = "John"

	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "mario"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameour := "yoshi"
	// fmt.Println(nameour)

	// //integer
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 35
	// fmt.Println(ageOne, ageTwo, ageThree)

	// //bits and memeory
	// var numOne int8 = 25
	// fmt.Println(numOne)

	// fmt.Print("hlelo")
	// fmt.Print("fwlfw\n")
	// fmt.Println("myafge is :", numOne, "Applie", ageTwo)
	// fmt.Printf("my age is %v and my name is %v\n", ageOne, nameOne)
	// fmt.Printf("my age is %q and my name is %q", ageOne, nameOne)
	// fmt.Printf("\nage is of type %T", ageOne)
	// fmt.Printf("\nage is of type %0.2f", 22.3345)

	// var str = fmt.Sprintf("my age is %v and my name is %v\n", ageOne, nameOne)
	// fmt.Println("saved str : ", str)

	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"yoshi", "ma", "pe", "bowser"}

	// names[1] = "luigi"
	// fmt.Println(ages, len(ages), cap(ages))
	// fmt.Println(names, len(names), cap(names))

	// var scores = []int{100, 50, 60}
	// scores[2] = 55
	// fmt.Println(scores, len(scores), cap(scores))
	// scores = append(scores, 90)
	// fmt.Println(scores, len(scores), cap(scores))

	// rangeOne := names[1:3]
	// rangeTwo := names[1:]
	// fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
	// fmt.Println(rangeTwo, len(rangeTwo), cap(rangeTwo))
	sayH([]string{"yoshi", "mario", "peach"}, sayG)
	sayH([]string{"yoshi", "mario", "peach"}, sayB)
	fmt.Println(intials("Devyansh Mehta"))
}

// This is a simple Go program that prints "Hello, World!" to the console.
// It demonstrates the basic structure of a Go program, including the main package and the main function.
