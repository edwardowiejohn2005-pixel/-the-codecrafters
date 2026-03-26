**GO_SYNTAX**

Go file consists of the following

- Package declaration
- Import packages
- Functions
- Statements and expressions

Now using this code

package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}

- Every program belong to a package so it is defined by using package as the key word and this particular code belongs to the main package
- The keyword import allow us to import packages eg. fmt, buffio, string etc
- Having a blank line make your code more readable
- func main() is a function and any code that is inside the curly bracket {} will be executed when commanded
- fmt.Println() basically print out the output of your code 

**What is a Go statement**

fmt.Println("Hello World!") is a go statment reason because hitting the enter key adds ";" to the end of the line implicitly
the left curly bracket { cannot come at the start of a line

**WHAT IS A COMPACT CODE**

Simply a compact code is mostly writing in one line 

*EXAMPLE*

package main; import ("fmt"); func main() { fmt.Println("Hello World!");}

