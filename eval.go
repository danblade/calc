package main

import (
	"fmt"
	"strconv"
	"strings"
)

func eval(s string) (int, error) {
	s = strings.TrimSpace(s)
	separated := strings.Fields(s)
	if len(separated) != 3 {
		return 0, fmt.Errorf("Please follow the format of <number> <function> <number>?")
	}
	firstNum, err := strconv.Atoi(separated[0])
	if err != nil {
		return 0, fmt.Errorf("Please follow the format of <number> <function> <number>?")
	}
	secondNum, err := strconv.Atoi(separated[2])
	if err != nil {
		return 0, fmt.Errorf("Please follow the format of <number> <function> <number>?")
	}
	switch separated[1] {
	case "+":
		return (firstNum + secondNum), nil
	case "-":
		return (firstNum - secondNum), nil
	case "/":
		return (firstNum / secondNum), nil
	case "*":
		return (firstNum * secondNum), nil
	default:
		return 0, fmt.Errorf("Invalid operator")
	}
}
