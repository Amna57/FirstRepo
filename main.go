package main

import "fmt"

func main() {

	fmt.Println("Commit for beta branch")
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)
	for {
		if input <= 0 {
			fmt.Print("Please enter a number greater than 0: ")
			fmt.Scanln(&input)
		} else {
			break
		}
	}
	fmt.Println("You entered", input)

	fmt.Print("Enter another number: ")
	var input2 int
	fmt.Scanln(&input2)
	for {
		if input2 <= 0 {
			fmt.Print("Please enter a number greater than 0: ")
			fmt.Scanln(&input2)
		} else if input2 == input {
			fmt.Print("Please enter a different number: ")
			fmt.Scanln(&input2)
		} else {
			break
		}
	}
	fmt.Println("You entered", input2)

	fmt.Print("Enter the number of iterations: ")
	var itr int
	fmt.Scanln(&itr)
	for {
		if itr <= 0 {
			fmt.Print("Please enter a number greater than 0: ")
			fmt.Scanln(&itr)
		} else {
			break
		}
	}

	for i := 1; i <= itr; i++ {
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
