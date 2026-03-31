package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		var input1 float64
		var input2 float64
		var operator string

		fmt.Println("<1> Add")
		fmt.Println("<2> Subtract")
		fmt.Println("<3> Multiply")
		fmt.Println("<4> Divide")
		fmt.Println("<5> Quit")
		fmt.Print("Select Operator: ")
		fmt.Scan(&operator)

		if operator == "5" {
			fmt.Println("Goodbye! Have a nice day!")
			os.Exit(0)
		}

		if operator != "1" && operator != "2" && operator != "3" && operator != "4" && operator != "5" {
			fmt.Println("Help: Please select a valid operator (1-5) to perform a calculation.")
			continue
		}

		fmt.Print("Enter first number: ")
		fmt.Scan(&input1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&input2)

		switch operator {
		case "1":
			fmt.Println("Result =", input1+input2)
		case "2":
			fmt.Println("Result =", input1-input2)
		case "3":
			fmt.Println("Result =", input1*input2)
		case "4":
			if input2 == 0 {
				fmt.Println("Error: Cannot divide by zero.")
			} else {
				fmt.Println("Result =", input1/input2)
			}
		default:
			fmt.Println("Error: Invalid operator selected.")
		}

		fmt.Println()
	}
}
