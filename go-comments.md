**What is a Go-comment**

A comment is the part of the code that won't be executed when the is being ran

Comments can also be used to prevent code execution when testing an alternative code.

Go-language supports single and multi-line comments

**How to know a single line comment**

a single line comment begins with two double forward slashes "//"

Any text between // and the end of the line is ignored by the compiler (will not be executed).


**Example**

// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}

**How to know a multi line comment**

Multi-line comments start with /* and ends with */.

Any text between /* and */ will be ignored by the compiler

**Example**

package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}

