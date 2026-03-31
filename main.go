package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDecimal(Hexstr string) (int64, error) {
	num, err := strconv.ParseInt(Hexstr, 16, 64)
	if err != nil {
		return 0, err
	}
	return num, err
}

func BinToDecimal(Binstr string) (int64, error) {
	num2, err := strconv.ParseInt(Binstr, 2, 64)
	if err != nil {
		return 0, err
	}
	return num2, err
}

func DecimalToBin(DecStri string) (string, string, error) {
	DecStr, err := strconv.ParseInt(DecStri, 10, 64)
	num3 := strconv.FormatInt(DecStr, 2)
	num4 := strconv.FormatInt(DecStr, 16)
	return strings.ToUpper(num4), num3, err
}

func main() {
	for {
		var converter string
		var num string

		fmt.Println("Select operation")
		fmt.Println("<1> Hexadecimal")
		fmt.Println("<2> Binary")
		fmt.Println("<3> Bin and Hex to Decimal")
		fmt.Scan(&converter)

		if converter == "quit" {
			fmt.Println("Have a nice day")
			break

		}

		fmt.Println("Enter Number for convertion")
		fmt.Scan(&num)

		if num == "0" {
			fmt.Println("Empty")
			continue
		}
		if converter == "0" {
			fmt.Println("Empty")
			continue
		}
		switch converter {
		case "1":
			fmt.Println(HexToDecimal(num))
		case "2":
			fmt.Println(BinToDecimal(num))
		case "3":
			bin, hex, err := DecimalToBin(num)

			fmt.Printf("Bin: %s\n", bin)

			fmt.Printf("Hex: %s\n", hex)
			fmt.Println(err)
		}

	}
}
