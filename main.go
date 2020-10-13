package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() { //Guess game from Head First Go

	fmt.Println("This is a calculator.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		separated := strings.Fields(input)
		if len(separated) != 3 {
			fmt.Println("Please follow the format of <number> <function> <number>?")
			continue
		}
		firstNum, err := strconv.Atoi(separated[0])
		if err != nil {
			fmt.Println("Please follow the format of <number> <function> <number>?")
			continue
		}
		secondNum, err := strconv.Atoi(separated[2])
		if err != nil {
			fmt.Println("Please follow the format of <number> <function> <number>?")
			continue
		}
		switch separated[1] {
		case "+":
			fmt.Println(firstNum + secondNum)
		case "-":
			fmt.Println(firstNum - secondNum)
		case "/":
			fmt.Println(firstNum / secondNum)
		case "*":
			fmt.Println(firstNum * secondNum)
		default:
			fmt.Println("Invalid operator")
		}

	}

}
