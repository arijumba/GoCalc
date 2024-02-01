package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func RomanToInt(num string) int {
	a := []rune(num)
	if unicode.IsDigit(a[0]) {
		panic("Use a single number system")
	}
	roman := make(map[string]int)
	roman["I"] = 1
	roman["II"] = 2
	roman["III"] = 3
	roman["IV"] = 4
	roman["V"] = 5
	roman["VI"] = 6
	roman["VII"] = 7
	roman["VIII"] = 8
	roman["IX"] = 9
	roman["X"] = 10
	result := roman[num]
	if result == 0 {
		panic("Enter a range of values from I to X")
	} else {
		return result
	}
}

func IntToRoman(num int) string {
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanStr := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNum := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			romanNum += romanStr[i]
		}
	}
	return romanNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the expression:")
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(text)
	operands := strings.Split(text, " ")

	if len(operands) != 3 {
		fmt.Println("Invalid input format. Please provide two operands and an operator separated by spaces.")
		return
	}

	operands[0] = (strings.TrimSpace(operands[0]))
	operands[1] = (strings.TrimSpace(operands[1]))
	operands[2] = (strings.TrimSpace(operands[2]))

	for _, operand := range operands {
		if operand == " " || operand == "" {
			panic("Enter correct value format, entered string is not a mathematical operation")
		}
	}

	values := "XVI"
	values = strings.TrimSpace(values)
	result := 0
	text = strings.ToUpper(text)
	if strings.ContainsAny(text, values) {
		a := RomanToInt(operands[0])
		b := RomanToInt(operands[2])
		switch operands[1] {
		case "+":
			result = a + b
		case "-":
			if a >= b {
				result = a - b
			} else {
				panic("There are no negative numbers in the Roman system")
			}
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			{
				panic("Error: invalid operator")
			}
		}
		if IntToRoman(result) == "" {
			panic("There is no zero in the Roman numeral system")
		} else {
			fmt.Println(IntToRoman(result))
		}
	} else {
		opp1, err0 := strconv.Atoi(operands[0])
		if err0 != nil {
			panic("Only integer value")
		}
		opp2, err := strconv.Atoi(operands[2])
		if err != nil {
			panic("Only integer value")
		}
		if (opp1 <= 10) && (opp2 <= 10) {
			switch strings.TrimSpace(operands[1]) {
			case "+":
				result = opp1 + opp2
			case "-":
				result = opp1 - opp2
			case "*":
				result = opp1 * opp2
			case "/":
				result = opp1 / opp2
			default:
				{
					panic("Invalid operator")
				}
			}
			fmt.Printf("%d\n", result)
		} else {
			panic("Enter a range of values from 1 to 10")
		}
	}
}
