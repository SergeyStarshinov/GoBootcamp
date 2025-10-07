package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Input left operand: ")
	leftOperand := inputOperand()
	fmt.Print("Input operation: ")
	operation := inputOperation()
	fmt.Print("Input right operand: ")
	rightOperand := inputOperand()

	var result float64
	switch operation {
	case "+":
		result = leftOperand + rightOperand
	case "-":
		result = leftOperand - rightOperand
	case "*":
		result = leftOperand * rightOperand
	default:
		result = leftOperand / rightOperand
	}
	fmt.Printf("Result of calculations: %.3f\n", result)
}

func inputOperand() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			fmt.Print("Invalid input, try again: ")
		} else {
			input := scanner.Text()
			result, errParse := strconv.ParseFloat(input, 64)
			if errParse != nil {
				fmt.Print("Invalid input, try again (any number): ")
			} else {
				return result
			}
		}
	}
}

func inputOperation() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			fmt.Println("Invalid input, try again")
		} else {
			input := scanner.Text()
			if input != "+" && input != "-" && input != "*" && input != "/" {
				fmt.Print("Invalid input, try again (+, -, * or /): ")
			} else {
				return input
			}
		}
	}
}
