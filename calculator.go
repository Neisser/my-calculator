package myCalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) parseInteger(operated string) (int, error) {
	result, err := strconv.Atoi(operated)
	return result, err
}

func (c calc) operate(input string, operator string) (int, error) {
	cleanInput := strings.Split(input, operator)
	first, err := c.parseInteger(cleanInput[0])
	if err != nil {
		return 0, err
	}
	second, err := c.parseInteger(cleanInput[1])
	if err != nil {
		return 0, err
	}
	switch operator {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		return first / second, nil
	default:
		fmt.Println(operator, "Operator is not supported!")
		return 0, nil
	}
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ProcessResult(input string, operator string) {
	c := calc{}
	value, err := c.operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

// Main calculator function
// func main() {
// 	input := readInput()
// 	operator := readInput()
// 	processResult(input, operator)
// }
