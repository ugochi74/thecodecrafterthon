package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func main() {

func printHelp() {
	fmt.Println("\nThis are the commands:")
	fmt.Println("input1+input2  = Isaddition")
	fmt.Println("input1-input2 = Issubtraction")
	fmt.Println("input1*input2 = Ismultiplication")
	fmt.Println("input1/input2  = Isdivision")
	fmt.Println("input1%input2  = Ismodulus")
	fmt.Println("help = show this message")
	fmt.Println("continue = restart")
	fmt.Println("quit = exit")
}
func main() {
	// ops := map[string]bool{"add": true, "sub": true, "mul": true, "div": true}

	//fmt.Println("CLI Calculator — type 'help' for usage, 'quit' to exit, continue")

	// var greeting string
	var input1, input2 string
	var operator string
	var result float64
	for {
		fmt.Println("Choose operation(continue, exit, help, quit): ")

		var line string
		fmt.Scanln(&line)

		cmd := strings.ToLower(strings.TrimSpace(line))

		if cmd == "quit" {
			fmt.Println("Goodbye!")
			return
		}

		if cmd == "help" {
			printHelp()
			continue
		}
		//fmt.Println("Welcome to CLI calculator. Please choose instruction below(calculator,help): ")
		// fmt.Println(&greeting)
		fmt.Println("Enter num1: ")
		fmt.Scanln(&input1)

		fmt.Println("Choose operator(+, *, /, %, -): ")
		fmt.Scanln(&operator)

		fmt.Println("Enter num2: ")
		fmt.Scanln(&input2)

		num1, err := strconv.ParseFloat(input1, 64)
		num2, err1 := strconv.ParseFloat(input2, 64)
		if err != nil || err1 != nil {
			fmt.Println("Invalid number: ")
			continue

		}

		switch operator {
		case "*":
			result = num1 * num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "/":
			if num2 == 0 {
				fmt.Println("No division by 0: ")
				continue
			}
			result = num1 / num2

		case "%":
			if num2 == 0 {
				fmt.Println("No mod by 0: ")
				continue
			}
			result = float64(int(num1) % int(num2))
		default:
			fmt.Println("unknown command. Type 'help'. ")
			continue
		}
		fmt.Println(result)
		// var options string
		// fmt.Println("Choose options(exit, continue): ")
		// fmt.Scanln(&options)
		// if options == "exit" {
		// 	break
		// }
	}
}
