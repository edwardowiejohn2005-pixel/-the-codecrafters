**What are variables**

Variables are vessels used for storing data values

**Type of Go-Variable**

- int - this stores integers which are whole numbers
- float32 - this stores numbers with decimals
- string - this stores text such as "Hello World". String value are surrounded by double quotes
- bool - this stores values that have two states: true or false

**Declaring  Variables**

there are basically two ways to declare a variable which are...

- With the keyword "var" : Use the var keyword, followed by variable name and type
--> var school string = learn2earn <--

- with the := sign : Use the := sign followed by the variable value
--> school := learn2earn <--

**Variable declaration with initial value**

If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:

**Example**

package main
import ("fmt")

func main() {
  var student1 string = "Edward"
  var student2 = "chioma"
  x := 2 

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}

**Variable declaration without initial value**

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type

**Example**

package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

**Value Assignment After Declaration**

Yes it is possible to assign a value to a variable after it is declared

**Example**

package main
import ("fmt")

func main() {
  var student1 string
  student1 = "Edward"
  fmt.Println(student1)
}

**GO multiple variable declaration**
In Go, it is possible to declare multiple variables on the same line.

**Example**

package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

if the type key word is not specified, you can declare different types of variables on the same line...

**Example**

package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

**Go Variable Declaration in a Block**

Multiple variable declarations can also be grouped together into a block for greater readability


**Example**

package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}


**Go Variable Naming Rules**

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords


**Multi-Word Variable Names**

Variable names with more than one word can be difficult to read.

There are several techniques you can use to make them more readable:

*Camel Case*

Each word, except the first, starts with a capital letter:

myVariableName = "John"

*Pascal Case*

Each word starts with a capital letter:

MyVariableName = "John"

*Snake case*

Each word is separated by an underscore character:

my_variable_name = "John"

