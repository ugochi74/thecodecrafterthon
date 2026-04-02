package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	lastResult    float64
	hasLastResult bool
	history       []string
	calcHistory   []string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to C&C Console. Type 'help' for commands.")

	for {
		fmt.Print("\nC&C> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		// pipe check must be BEFORE splitting into parts
		if strings.Contains(input, "|") {
			handlePipe(input)
			continue
		}

		// split OUTSIDE the if block
		parts := strings.Fields(input)
		if input == "Calc" {
			fmt.Println("Choose operation(calc add <a> b>, calc sub <a> <b>, calc div <a> <b>, calc mul <a> <b>, calc power <a> <b>, lastresult, history, haslastresult)")
		}

		switch strings.ToLower(parts[0]) {
		case "calc":
			handleCalc(parts[1:])
		case "base":
			handleBase(parts[1:])
		case "str":
			handleStr(parts[1:])
		//case "help":
		//printHelp()
		case "history":
			printHistory()
		case "haslastresult":
			if hasLastResult {
				fmt.Println(" Yes, there is a last result.")
			} else {
				fmt.Println(" No result yet in this session.")
			}
		case "lastresult":
			if !hasLastResult {
				fmt.Println(" No previous result in this session.")
			} else {
				fmt.Printf(" Last Result: %g\n", lastResult)
			}
		case "exit":
			fmt.Println("Shutting down. Goodbye.")
			return
		default:
			fmt.Printf(" Unknown command: %q. Type 'help'\n", parts[0])
		}
	}
}

// func handleCalc(parts []string) {}
func handleBase(parts []string) {}
func handleStr(parts []string)  {}
func handlePipe(input string)   {}

// func printHelp()                {}
func printHistory() {
	if len(history) == 0 {
		fmt.Println(" No history yet.")
		return
	}
	fmt.Println(" History:")
	start := len(history) - 5
	if start < 0 {
		start = 0
	}
	for i, entry := range history[start:] {
		fmt.Printf("   %d. %s\n", i+1, entry)
	}
}

func setLastResult(n float64) {
	lastResult = n
	hasLastResult = true
}

func addHistory(entry string) {
	history = append(history, entry)
}

// keep math and strconv used
var _ = math.Mod
var _ = strconv.ParseFloat

func handleCalc(parts []string) {
	// step 1 — check if parts is empty
	if len(parts) == 0 {
		//fmt.Println("Error: missing command. Usage: calc <operation> <a> <b>")
		return
	}

	// step 2 — get operation
	operation := strings.ToLower(parts[0])
	//fmt.Println("Choose operation(calc add <a> b>, calc sub <a> <b>, calc div <a> <b>, calc mul <a> <b>, calc power <a> <b>, lastResult, history, hasLastResult)")

	// step 3 — check for "last"
	if operation == "last" || operation == "lastresult" {
		if !hasLastResult {
			fmt.Println(" No previous result in this session.")
			return
		}
		fmt.Printf(" Last Result: %g\n", lastResult)
		return
	}

	// step 4 — check for "history"
	if operation == "history" {
		if len(calcHistory) == 0 {
			fmt.Println("No calc history yet.")
			return
		}
		fmt.Println("Calc History:")
		start := len(calcHistory) - 5
		if start < 0 {
			start = 0
		}
		for i, entry := range calcHistory[start:] {
			fmt.Printf("       %d. %s\n", i+1, entry)
		}
		return
	}

	// step 5 — all other operations need 2 arguments
	if len(parts) < 3 {
		fmt.Println(" Error: missing arguments. Usage: calc <operation> <a> <b>")
		return
	}

	// step 6 — parse a and b
	a, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Printf(" Error: %q is not a valid number\n", parts[1])
		return
	}
	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Printf("Error: %q is not a valid number\n", parts[2])
		return
	}

	// step 7 — switch on operation
	var result float64
	switch operation {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			fmt.Println("Error: cannot divide by zero.")
			return
		}
		result = a / b
	case "mod":
		if b == 0 {
			fmt.Println("Error: cannot mod by zero.")
			return
		}
		result = math.Mod(a, b)
	case "pow":
		result = math.Pow(a, b)

	default:
		fmt.Printf("Error: unknown operation %q. Use: add sub mul div mod pow\n", operation)
		return
	}

	// step 8 — print result and store it
	fmt.Printf(" Result: %g\n", result)
	setLastResult(result)
	entry := fmt.Sprintf("calc %s %g %g → %g", operation, a, b, result)
	calcHistory = append(calcHistory, entry)
	addHistory(entry)
}
