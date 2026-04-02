package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)
	//var choice int
	//var stringconvert string

	for {

		// Read choice
		//fmt.Print("Choose conversion (upper<1>, lower<2>, cap<3>, reverse<4>): ")
		fmt.Print("Choose  conversion(upper<1>, lower<2>, cap<3>,  reverse<4>, snake<5>, title<6>): ")
		choiceText, _ := reader.ReadString('\n')
		choiceText = strings.TrimSpace(choiceText)

		choice, err := strconv.Atoi(choiceText)
		if err != nil {
			fmt.Println("Invalid choice, please enter a number 1-6")
			continue
		}
		fmt.Println("Enter word to convert: ")
		stringconvert, _ := reader.ReadString('\n')
		stringconvert = strings.TrimSpace(stringconvert)
		if stringconvert == "" {
			fmt.Println("Please enter a word: ")
			continue
		}

		//fmt.Println("Enter string to convert(upper<1>, lower<2>, cap<3>,  reverse<4>): ")
		//fmt.Println("Choose choice: ")
		//fmt.Scanln(&choice)
		//fmt.Println("Enter word to convert: ")
		//fmt.Scanln(&stringconvert)

		switch choice {
		case 1:
			fmt.Println(upper(stringconvert))
		case 2:
			fmt.Println(lower(stringconvert))
		case 3:
			fmt.Println(cap(stringconvert))
		//case 4:
		//  fmt.Println(title(stringconvert))
		case 4:
			fmt.Println(reverse(stringconvert))
		case 5:
			fmt.Println(snake(stringconvert))
		case 6:
			fmt.Println(title(stringconvert))
		}
	}
}

// func readLine(prompt string) string {
//  fmt.Print(prompt)
//  text, _ := reader.ReadString('\n')
//  return strings.TrimSpace(text)
// }

func upper(s string) string {
	return strings.ToUpper(s)
}

func lower(s string) string {
	return strings.ToLower(s)
}

func cap(s string) string {
	text := strings.ToLower(s)
	return strings.Title(text)
}

func reverse(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func snake(s string) string {
	text := strings.ToLower(s)
	result := strings.ReplaceAll(text, " ", "_")
	return result
}
func title(s string) string {
	str := strings.Fields(s)
	for i := range str {
		str[i] = strings.ToLower(str[i])
		if i > 1 && (str[i] == "a" || str[i] == "an" || str[i] == "the" || str[i] == "and" || str[i] == "is") {
			str[i] = strings.ToLower(str[i])
		} else if i > 1 && (str[i] == "but" || str[i] == "or" || str[i] == "for" || str[i] == "nor") {
			str[i] = strings.ToLower(str[i])
		} else if i > 1 && (str[i] == "on" || str[i] == "at" || str[i] == "to" || str[i] == "by" || str[i] == "it") {
			str[i] = strings.ToLower(str[i])
		} else if i > 1 && (str[i] == "in" || str[i] == "of" || str[i] == "up" || str[i] == "as") {
			str[i] = strings.ToLower(str[i])
		} else {
			str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(string(str[i][1:])))
		}
		//return strings.Join(str, " ")

	}
	return strings.Join(str, " ")
}
