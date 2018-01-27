package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInteractiveMenu() {
	fmt.Println("Welcome to our Go Calculator")
	fmt.Println("")
	fmt.Println("[1] Add two numbers")
	fmt.Println("[2] Substract two numbers")
	fmt.Println("[3] Multiply two numbers")
	fmt.Println("[4] Divide two numbers")
	fmt.Println("Please select your operation:")
}

func readNumber() int {
	fmt.Println("Please input your number")

	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
	}

	fmt.Println("You have input", i)
	return i

}

func addition(x, y int) int {
	fmt.Println("You are going to sum ", x, " and ", y)
	return x + y
}

func subtract(x, y int) int {
	fmt.Println("You are going to subtract ", x, " and ", y)
	return x - y
}

func multiply(x, y int) int {
	fmt.Println("You are going to multiply ", x, " and ", y)
	return x * y
}

func division(x, y int) int {
	fmt.Println("You are going to multiply ", x, " and ", y)
	if y == 0 {
		fmt.Println("Please don't try to break the rules of nature, please.")
		fmt.Println("One ring to rule them all")
		return 1 // this is cheating
	}

	return x / y
}

func readInput() rune {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("An error has ocurred")
	}

	return char
}

func main() {
	printInteractiveMenu()
	input := readInput()

	switch input {
	case '1':
		result := addition(readNumber(), readNumber())
		fmt.Println("The result of your chosen addition operation is as follows behold: ", result)
	case '2':
		result := subtract(readNumber(), readNumber())
		fmt.Println("The result of your chosen subtraction operation is as follows behold: ", result)
	case '3':
		result := multiply(readNumber(), readNumber())
		fmt.Println("The result of your chosen multiplication operation is as follows behold: ", result)
	case '4':
		result := division(readNumber(), readNumber())
		fmt.Println("The result of your chosen division operation is as follows behold: ", result)
	}

}
