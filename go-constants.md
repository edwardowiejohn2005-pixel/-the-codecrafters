**What are go-constants**

If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

**Example**
const CONSTNAME type = value

**Declaring a Constant**

Example of declaring a constant in GO

package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}

**Constant Rules**

- Constant names follow the same naming rules as variables
- Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
- Constants can be declared both inside and outside of a function

**Constant Types**

There are two types of constants:

- Typed constants
- Untyped constants

**Typed Constants**
Typed constants are declared with a defined type:

**Example**

package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}

**Untyped Constants**

Untyped constants are declared without a type:

**Example**

package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}

**Multiple Constants Declaration**

Multiple constants can be grouped together into a block for readability:

**Example**

package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}