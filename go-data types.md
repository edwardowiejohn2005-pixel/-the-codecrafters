**Go Data Types**


Data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

Go has three basic data types:

- bool: represents a boolean value and is either true or false
- Numeric: represents integer types, floating point values, and complex types
- string: represents a string value


**Example**

This example shows some of the different data types in Go:

package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}


**Boolean Data Type**

A boolean data type is declared with the bool keyword and can only take the values true or false.

The default value of a boolean data type is false.

**Example**

package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}

**Go Integer Data Types**

Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

The integer data type has two categories:

- Signed integers - can store both positive and negative values
- Unsigned integers - can only store non-negative values

**Signed Integers**

Signed integers, declared with one of the int keywords, can store both positive and negative values:

**Example**

package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}


**Unsigned Intergers**

Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

Example
package main
import ("fmt")

func main() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

**Go Float Data Types**


The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

The float data type has two keywords:

Type|           Size|           Range

float32|        32 bits|    	-3.4e+38 to 3.4e+38.

float64|        64 bits|    	-1.7e+308 to +1.7e+308.


**Go String Data Type**

String Data Type
The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

Example


package main
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
Result:

Type: string, value: Hello!
Type: string, value:
Type: string, value: World 1
