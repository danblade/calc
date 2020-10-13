package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		calc, err := eval(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(calc)
	}

}
