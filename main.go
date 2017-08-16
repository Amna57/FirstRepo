package main

import "fmt"

func main() {

	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)
	fmt.Println("You entered", input)

	fmt.Print("Enter another number: ")
	var input2 int
	fmt.Scanln(&input2)
	fmt.Println("You entered", input2)

	for i := 1; i <= 100; i++ {
		if i%input == 0 && i%input2 == 0 {
			fmt.Println(i, ": fizzbuzz")
		} else if i%input == 0 {
			fmt.Println(i, ": fizz")
		} else if i%input2 == 0 {
			fmt.Println(i, ": buzz")
		} else {
			fmt.Println(i)
		}
	}
}
