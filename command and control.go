   // CodeCrafters — Hackathon 002
         // Squad: [The Structs]
         // Member: [ina isah, iisah]
         // Member: [omafu samuel, somafu]
         // Member: [Otete Benjamin, botete]
         // Member: [Odeh owoicho emmanuel, emmodeh]
         // Member: [Wealth Ibe, wibe]
         // Member: [Odumu John, jodumu]
         // Member: [Edward John, edwjohn]
         // Member: [inaju Daniel, dinaju]


package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var start string

func calc() {
	fmt.Println("Welcome to the Calculator")
	fmt.Println("Enter 'help' for guidelines")

	scanner := bufio.NewScanner(os.Stdin)
	var history []string

	for {
		scanner.Scan()
		text := scanner.Text()
		text = strings.TrimSpace(text)

		if text == "exit" {
			main()
		}

		if text == "help" {
			fmt.Println("calc add <a> <b>      → addition")
			fmt.Println("calc sub <a> <b>      → subtraction")
@@ -77,7 +81,7 @@
		}

		if len(part) < 3 {
			fmt.Println("not enough argument")
			fmt.Println("Error: Not enough argument, e.g add 87 77, sub 8 6.")
			continue
		}

@@ -182,6 +186,11 @@
		if text == "str" {
			mainstring()
			continue

		}
		if text == "exit" {
			fmt.Println("Goodbye..............!!")
			break
		}
	}
}
@@ -190,6 +199,8 @@
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Welcome To Base Converter!")
		fmt.Println("Enter help to continue.")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
@@ -198,131 +209,131 @@
			fmt.Println(" base dec <number>   → convert decimal to binary AND hex")
			fmt.Println(" base hex <number>   → convert hex to decimal")
			fmt.Println(" base bin <number>   → convert binary to decimal")
			fmt.Println("exit")
			continue
		}
		if input == "exit" {
			fmt.Println("goodbye")
			break
			main()
		}
		if input == "" {
			fmt.Println("error")
			continue
		}
		conv := strings.Fields(input)
		value := conv[1]
		num := conv[2]

		if conv[0] != "base" {
			fmt.Println("invalid command")
			continue
		}

		switch num {
		case "hex":
			n, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "bin":
			n, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "dec":
			n, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Printf("Binary: %b\n", n)
			fmt.Printf("Hex: %X\n", n)
		case "exit":
			main()
		default:
			fmt.Println("Invalid base!")
		}
	}
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	return reverse(word[1:]) + word[:1]
}

func snakecase(word string) string {
	return strings.ReplaceAll(word, " ", "_")
}

func mainstring() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a word")

		input, _ := reader.ReadString('\n')
		word := strings.Fields(input)

		var Transformation string

		fmt.Scanln(&word)

		if len(word) == 0 {
			fmt.Println("no text provided")
			continue
		}

		fmt.Println("chose a Transformation:")
		fmt.Println("<1> Toupper")
		fmt.Println("<2> Tolower")
		fmt.Println("<3> Fields")
		fmt.Println("<4> Title")
		fmt.Println("<5> Snake")
		fmt.Println("<6> Reverse")
		fmt.Println("<7> Exit")

		fmt.Scanln(&Transformation)

		switch Transformation {
		case "1":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Result:", strings.ToUpper(input))
		case "2":
			result := word
			input := strings.Join(result, " ")

			fmt.Println("Return:", strings.ToLower(input))
		case "3":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Fields(input))
		case "4":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Title(input))
		case "5":
			result := word
			input := strings.Join(result, " ")
			result1 := strings.ToLower(input)
			fmt.Println(snakecase(result1))

		case "6":
			good := strings.Fields(input)
			for i := 0; i < len(good); i++ {
				good[i] = reverse(good[i])
			}
			fmt.Println(strings.Join(good, " "))
		case "7":
			main()
		default:
			fmt.Println("invalide input")

		}
	}
}