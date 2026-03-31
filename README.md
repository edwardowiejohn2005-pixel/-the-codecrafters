Number Base Converter (Go)

A simple command-line tool written in Go that converts numbers between Hexadecimal, Binary, and Decimal formats.

The program allows the user to choose a conversion operation and input a number through the terminal.

---

Features

The application supports the following conversions:

1. Hexadecimal → Decimal
2. Binary → Decimal
3. Decimal → Binary and Hexadecimal

The program runs continuously in a loop until the user types "quit".

---

Technologies Used

- Go Programming Language
- Standard Go Libraries:
  - "fmt" – used for input and output
  - "strconv" – used for converting numbers between different bases
  - "strings" – used for formatting hexadecimal output

---

Project Structure

.
├── main.go
└── README.md

---

Functions Overview

HexToDecimal

Converts a hexadecimal string into a decimal number.

func HexToDecimal(Hexstr string) (int64, error)

Example:

Input:

FF

Output:

255

---

BinToDecimal

Converts a binary string into a decimal number.

func BinToDecimal(Binstr string) (int64, error)

Example:

Input:

1010

Output:

10

---

DecimalToBin

Converts a decimal number into both Binary and Hexadecimal.

func DecimalToBin(DecStri string) (string, string, error)

Example:

Input:

15

Output:

Binary: 1111
Hexadecimal: F

---

How the Program Works

1. The program displays a menu in the terminal:

Select operation
<1> Hexadecimal
<2> Binary
<3> Bin and Hex to Decimal

2. The user selects an operation.
3. The user enters the number to convert.
4. The program prints the result.
5. The program continues running until the user types "quit".

---

Example Usage

Example 1:

Select operation
<1> Hexadecimal
<2> Binary
<3> Bin and Hex to Decimal

1
Enter Number to be converted
FF

255

Example 2:

Select operation
3
Enter Number to be converted
10

Bin: 1010
Hex: A

---

Running the Program

Make sure Go is installed.

Check your Go installation:

go version

Run the program:

go run main.go

---

Error Handling

The program uses Go's "strconv.ParseInt" function which returns an error if the input value is invalid.

Example error:

invalid syntax

---

Possible Improvements

Future improvements could include:

- Better input validation
- More descriptive error messages
- Support for Octal conversions
- Adding unit tests
- Improving the command-line interface

---

Author

Edward John