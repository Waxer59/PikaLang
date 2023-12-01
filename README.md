<p align="center">
  <img src="./docs/img/logo.webp" height="250px"/>
</p>

# PikaLang

Pika, or Pikalang, is a programming language designed to be simple, efficient and easy to use. It is written in Go and its syntax is inspired by several languages. Pika uses the `.pk` extension for its files.

- [PikaLang](#pikalang)
  - [CLI](#cli)
  - [Syntax](#syntax)
    - [Variables \& constants declaration](#variables--constants-declaration)
      - [variables](#variables)
      - [constants](#constants)
    - [If statements](#if-statements)
      - [Else Statement](#else-statement)
      - [Else If Statement:](#else-if-statement)
    - [Function Declaration](#function-declaration)
    - [Anonymous functions](#anonymous-functions)
    - [Switch statement](#switch-statement)
      - [Multiple Cases](#multiple-cases)
      - [Logical Cases](#logical-cases)
    - [Comments](#comments)
      - [Single-line Comments](#single-line-comments)
      - [Multi-line Comments](#multi-line-comments)
    - [Loops](#loops)
      - [For Loop](#for-loop)
      - [While Loop](#while-loop)
      - [Break Statement](#break-statement)
      - [Continue Statement](#continue-statement)
    - [Operators](#operators)
      - [Assignment Operators](#assignment-operators)
      - [Increment and Decrement Operators](#increment-and-decrement-operators)
        - [++ (Post-increment)](#-post-increment)
        - [-- (Post-decrement)](#---post-decrement)
        - [++ (Pre-increment)](#-pre-increment)
        - [-- (Pre-decrement)](#---pre-decrement)
      - [Logical operators](#logical-operators)
        - [OR](#or)
        - [AND](#and)
        - [NOT](#not)
      - [Ternary operator](#ternary-operator)
      - [String operators](#string-operators)
        - [Concatenation](#concatenation)
      - [Math operators](#math-operators)
        - [Addition](#addition)
        - [Subtraction](#subtraction)
        - [Multiplication](#multiplication)
        - [Division](#division)
        - [Module](#module)
        - [Exponentiation](#exponentiation)
      - [Comparison operators](#comparison-operators)
        - [Equality operator (==)](#equality-operator-)
        - [Inequality operator (!=)](#inequality-operator-)
        - [Greater than operator (\>)](#greater-than-operator-)
        - [Less than operator (\<)](#less-than-operator-)
        - [Greater than or equal to operator (\>=)](#greater-than-or-equal-to-operator-)
        - [Less than or equal to operator (\<=)](#less-than-or-equal-to-operator-)
    - [Data structures](#data-structures)
      - [arrays](#arrays)
        - [Array Declaration](#array-declaration)
        - [Negative Indexes](#negative-indexes)
        - [Accessing Array Elements](#accessing-array-elements)
        - [Modifying Array Elements](#modifying-array-elements)
        - [Array Length](#array-length)
      - [object](#object)
    - [Primitive data types](#primitive-data-types)
      - [string](#string)
      - [number](#number)
      - [boolean](#boolean)
      - [null](#null)
    - [Native functions](#native-functions)
      - [`print()`](#print)
      - [`len()`](#len)
      - [`includes()`](#includes)
      - [`push()`](#push)
      - [`pop()`](#pop)
      - [`shift()`](#shift)
      - [`indexOf()`](#indexof)
      - [`isNaN()`](#isnan)
      - [`isNull()`](#isnull)
      - [`prompt()`](#prompt)
      - [`randNum()`](#randnum)
      - [`pow()`](#pow)
      - [`string()`](#string-1)
      - [`num()`](#num)
      - [`bool()`](#bool)
      - [`toUpperCase()`](#touppercase)
      - [`toLowerCase()`](#tolowercase)
      - [`capitalize()`](#capitalize)
      - [`startsWith()`](#startswith)
      - [`endsWith()`](#endswith)
      - [`reverseString()`](#reversestring)
      - [`typeof()`](#typeof)
      - [`concat()`](#concat)

## CLI

The Pikalang CLI (Command Line Interface) provides a convenient way to interact with the Pikalang interpreter and run programs written in this language.

```bash
NAME:
   pika - A simple pika compiler

USAGE:
   pika [global options] command [command options] [arguments...]

VERSION:
   <latest-version>

COMMANDS:
   run   Run a file
   help  Show help
   repl  Start the repl

GLOBAL OPTIONS:
   --version, -v  print the version
```

With go installed on your computer, run the following command to install the PikaLang CLI

```bash
go install github.com/Waxer59/PikaLang/cmd/pika@latest
```

## Syntax

Pikalang is a programming language designed to be simple and expressive. This section describes the basic syntax of Pikalang and the fundamental elements that make up a program in this language.

### Variables & constants declaration

#### variables

In Pikalang, variables are declared using the `var` keyword followed by the variable name. For example:

```js
var foo = "bar"
var bar = 42
```

#### constants

In Pikalang, constants are declared using the `const` keyword followed by the constant name, constants as they are not indicated are immutable. For example:

```js
const foo = "bar"
const bar = 42
```
### If statements

The 'if' statement is used to execute a block of code only if a specified condition is true. The syntax for the 'if' statement in our language supports two forms:

* First form: if (condition) { }
This form encloses the code block within curly braces immediately after the condition.

Example:
``` js
if (x > 5) {
    print("x is greater than 5")
}
```

* Second form: if condition { }
This form omits the parentheses around the condition and directly includes the code block.

Example:
```js
if x > 5 {
    print("x is greater than 5")
}
```

#### Else Statement
The 'else' statement follows an 'if' statement and is used to specify a block of code that should be executed if the condition in the preceding 'if' statement evaluates to false. The 'else' statement does not require any conditions. It is optional and can be omitted if not needed.

Example:
```js
if (x > 5) {
    print("x is greater than 5")
} else {
    print("x is not greater than 5")
}
```

#### Else If Statement:
The 'else if' statement allows for the evaluation of multiple conditions in a series of sequential checks. It is used when there are more than two possible outcomes based on different conditions. Multiple 'else if' statements can follow an 'if' statement, and the code block associated with the first true condition is executed. Only one code block will be executed, even if multiple conditions are true. Similar to the 'if' statement, the 'else if' statement supports the two syntax forms:

* First form: else if (condition) { }
This form encloses the code block within curly braces immediately after the condition.

Example:
```js
if (x > 5) {
    print("x is greater than 5")
} else if (x < 5) {
    print("x is less than 5")
}
```

* Second form: else if condition { }
This form omits the parentheses around the condition and directly includes the code block.

Example:
```js
if (x > 5) {
    print("x is greater than 5")
} else if x < 5 {
    print("x is less than 5")
}
```

Please note that nested 'if' statements are supported, allowing the inclusion of further 'if,' 'else if,' or 'else' statements within the code blocks.

### Function Declaration

In Pikalang, functions are defined using the keyword `fn`, followed by the function name, the parameters in parentheses, and the return type. The `return` statement is used to specify the value to be returned by the function.

Example:

```rs
fn add(x, y) {
  return x + y
}
```

In the example above, the `add` function takes two parameters, `x` and `y`, and returns their sum using the `return` statement. The return type of the function is not explicitly specified in the example, but it can be inferred based on the returned value.

The `return` statement is used to exit a function and return a value. Once a `return` statement is encountered, the function terminates, and the value specified after the `return` keyword is returned to the caller.

Example usage:

```rs
fn main() {
  var result = add(3, 4)
  print(result) // Output: 7
}

fn add(x, y) {
  return x + y
}
```

In the example above, the `add` function is called within the `main` function, and the returned value is assigned to the variable `result`. The value of `result` is then printed, resulting in the output `7`.

### Anonymous functions

Anonymous functions are those functions that are defined without a specific name. They are useful for situations where temporary functionality is needed without the requirement of declaring a function with a formal name.

```js
() => {
  // ...
}

const myFn = () => {
  // ...
}

const obj = {
  myObjFunc: () => {
    // ...
  }
}
```

### Switch statement

The switch statement allows you to perform different actions based on the value of a given expression. It provides a concise way to write multiple conditional statements and improve the readability of your code.

The switch statement evaluates the given condition and compares it against different cases. When a match is found, the corresponding block of code is executed. If no match is found, an optional default case can be specified to handle such scenarios.

In PikaLang's switch statement, the 'break' statement is not required. After executing a matching case block, the control automatically exits the switch statement. This means that each case is isolated and does not fall through to the next case by default. If you want to fall through to the next case, you can omit the 'break' statement.

In Pikalang the brackets in the switch statement parameter are optional so there are two types of syntax for the switch statement:

```go
switch condition {
  // Do something
}
```

```go
switch (condition) {
  // Do something
}
```

#### Multiple Cases

To execute the same block of code for multiple cases, you can use the `case` keyword with a comma-separated list of values. This allows you to specify multiple cases that should execute the same block of code. Here's an example to illustrate this concept: 

```go
switch expression {
    case value1, value2, value3:
        // Code to be executed for value1, value2, and value3
    case value4, value5:
        // Code to be executed for value4 and value5
    default:
        // Code to be executed if no matching case is found
}
```

#### Logical Cases

Switch cases can also have boolean expressions that if true will execute the case.

```go
switch expression {
    case value1, value2, value3:
        // Code to be executed for value1, value2, and value3
    case value4, value5 > 10:
        // Code to be executed for value4 and value5
    default:
        // Code to be executed if no matching case is found
}
```

### Comments

In PikaLang, comments are used to add explanatory notes or annotations within the code that are ignored by the compiler or interpreter. They are meant to provide information to developers and are not executed as part of the program.

PikaLang supports both single-line and multi-line comments.

#### Single-line Comments

A single-line comment starts with `//` and extends until the end of the line. It is used to add a comment on a single line.

```js
var greeting = "Hello, world!" // This is a single-line comment
```

> The comment text after `//` is ignored by the compiler or interpreter.

#### Multi-line Comments

Multi-line comments, also known as block comments, allow you to add comments that span multiple lines. They start with `/*` and end with `*/`. Everything between these delimiters is treated as a comment and is ignored by the compiler or interpreter.

```js
/*
This is a multi-line comment.
It can span multiple lines.
*/

var x = 10
```

> Remember that comments are for humans reading the code, so write clear and concise comments that enhance code readability and maintainability.

### Loops

#### For Loop

The `for` loop is a control flow statement in programming that allows you to repeatedly execute a block of code based on a specified condition. It is commonly used to iterate over a collection of items or perform a certain task a specific number of times.

```go
for initialization; condition; post {
    // Code to be executed
}
```

- The `initialization` step is optional and is used to initialize variables or set up the loop.
- The `condition` is evaluated before each iteration, and if it evaluates to `true`, the loop continues. If it evaluates to `false`, the loop terminates.
- The `post` step is executed after each iteration and is typically used to update the loop variables.

#### While Loop

The `while` loop is a control flow statement that executes a block of code repeatedly as long as a specified condition is true. It is used when the number of iterations is unknown and depends on the condition being evaluated.

Example of use:

```js
var i = 0
while i < 5 {
    print(i)
    i++
}
```

In the example above, the code inside the `while` loop will execute repeatedly as long as the condition `i < 5` is true. The variable `i` is incremented on each iteration.

#### Break Statement

The `break` statement is used to exit from a loop prematurely. It is often used when a certain condition is met and there is no need to continue the remaining iterations of the loop.

Example of use:

```js
var i = 0
while i < 5 {
    if i == 2 {
        i++
        break
    }
    print(i)
    i++
}
```

In the example above, the `break` statement is encountered when `i` is equal to 5, causing the loop to exit immediately.

#### Continue Statement

The `continue` statement is used to skip the remaining code in a loop iteration and move to the next iteration. It is often used when a certain condition is met, and you want to skip executing the rest of the loop's code for that specific iteration.

Example of use:

```js
var i = 0
while i < 5 {
    if i == 2 {
        i++
        continue
    }
    print(i)
    i++
}
```

### Operators

Operators are symbols or characters used in programming languages to perform operations on variables, values, or expressions. They are used to manipulate and compare data, control program flow, and perform logical operations.

#### Assignment Operators

Assignment operators are used to assign values to variables and perform operations simultaneously. They combine the assignment (=) operator with another arithmetic or logical operator to perform the operation and assign the result to the variable in a concise way.

Here are the assignment operators:

- `+=`: Adds a value to the variable and assigns the result to the variable.

Example of use:
```js
var x = 5
x += 3 // Equivalent to x = x + 3
// After this operation, the value of x will be 8
```

- `-=`: Subtracts a value from the variable and assigns the result to the variable.

Example of use:
```js
var x = 10
x -= 4 // Equivalent to x = x - 4
// After this operation, the value of x will be 6
```

- `*=`: Multiplies the variable by a value and assigns the result to the variable.

Example of use:
```js
var x = 3
x *= 5 // Equivalent to x = x * 5
// After this operation, the value of x will be 15
```

- `**=`: Raises the variable to a power and assigns the result to the variable.

Example of use:
```js
var x = 2
x **= 3 // Equivalent to x = x ** 3
// After this operation, the value of x will be 8
```

- `/=`: Divides the variable by a value and assigns the result to the variable.

Example of use:
```js
var x = 10
x /= 2 // Equivalent to x = x / 2
// After this operation, the value of x will be 5
```

- `%=`: Calculates the remainder of dividing the variable by a value and assigns the result to the variable.

Example of use:
```js
var x = 10
x %= 3 // Equivalent to x = x % 3
// After this operation, the value of x will be 1
```

- `=`: Assigns a value to the variable.

Example of use:
```js
var x = 5
var y = 10
x = y // Assigns the value of y to x
// After this operation, the value of x will be 10
```

#### Increment and Decrement Operators

The increment and decrement operators are used to modify the value of a variable by incrementing or decrementing it by 1. These operators can be applied both as post-increment/post-decrement operators and pre-increment/pre-decrement operators.

##### ++ (Post-increment)

The `++` operator is used to increment the value of a variable by 1. It can be used both as a post-increment and a pre-increment operator.

As a post-increment operator, the value of the variable is first used in the expression, and then it is incremented.

Example of post-increment use:

```js
var x = 5
var y = x++ // The value of y is assigned the current value of x (5), and then x is incremented to 6
// After this operation, the value of y is 5 and the value of x is 6
```

##### -- (Post-decrement)

The `--` operator is used to decrement the value of a variable by 1. It can be used both as a post-decrement and a pre-decrement operator.

As a post-decrement operator, the value of the variable is first used in the expression, and then it is decremented.

Example of post-decrement use:

```js
var x = 5
var y = x-- // The value of y is assigned the current value of x (5), and then x is decremented to 4
// After this operation, the value of y is 5 and the value of x is 4
```

##### ++ (Pre-increment)

As a pre-increment operator, the value of the variable is first incremented, and then it is used in the expression.

Example of pre-increment use:

```js
var x = 5
var y = ++x // The value of x is first incremented to 6, and then the value of y is assigned 6
// After this operation, the value of y is 6 and the value of x is 6
```

##### -- (Pre-decrement)

As a pre-decrement operator, the value of the variable is first decremented, and then it is used in the expression.

Example of pre-decrement use:

```js
var x = 5
var y = --x // The value of x is first decremented to 4, and then the value of y is assigned 4
// After this operation, the value of y is 4 and the value of x is 4
```

#### Logical operators

##### OR

The OR operator is a logical operator that operates on two or more operands, and it returns true if at least one of the operands is true. It can be represented using the symbol "||". Here's an example of the OR operator in action:

```js
var x = 5
var y = 10
var z = x > 3 || y < 5
print(z)  // Output: true
```

In the example above, the expression ` x > 3 || y < 5` evaluates to `true` because at least one of the conditions is true (in this case, `x > 3` is true).

##### AND

The AND operator is a logical operator that operates on two or more operands, and it returns true only if all the operands are true. It can be represented using the symbol "&&". Here's an example of the AND operator in action:

```js
var x = 5
var y = 10
var z = x > 3 && y < 5
print(z)  // Output: false
```
In the example above, the expression ` x > 3 && y < 5` evaluates to `false` because one of the conditions `y < 5` is false.

##### NOT

The NOT operator is a logical operator that operates on a single operand and returns the opposite boolean value. It can be represented using the symbol "!". Here's an example of the NOT operator in action:

```js
var x = 5
var y = 10
var z = !(x > 3)
print(z)  // Output: false
```

In the example above, the expression `!(x > 3)` evaluates to `false` because the condition `x > 3` is true, but the NOT operator reverses the boolean value.

#### Ternary operator

Ternary operators, also known as conditional operators, are operators that evaluate a boolean expression and return a result based on that evaluation. They have the following general syntax:

```js
condition ? true_expression : false_expression
```

The condition is a boolean expression that is evaluated first. If the condition is true, the value of true_expression is returned. If the condition is false, the value of false_expression is returned.

```js
var age = 18
var message = age >= 18 ? "You are an adult" : "You are a minor"
```

In this example, the condition `age >= 18` is evaluated. If the age is greater than or equal to 18, the result will be "You are an adult". Otherwise, the result will be "You are a minor". The resulting value is assigned to the variable `message`.

#### String operators

##### Concatenation

The concatenation operator is used to join two or more strings together. It takes two string values and returns a new string that is the combination of the operands.

```js
"Hello" + " " + "World" // Hello World
```

#### Math operators

Math operators are symbols or characters used to perform mathematical operations on one or more operands to produce a result. Here are some commonly used math operators

##### Addition

The addition operator is used to add two or more numbers together.

```js
1 + 1
```

##### Subtraction

The subtraction operator is used to subtract one number from another.

```js
1 - 1
```

##### Multiplication

The multiplication operator is used to multiply two or more numbers.

```js
1 * 1
```

##### Division

The division operator is used to divide one number by another.

```js
1 / 1
```

##### Module

The modulo operator is used to find the remainder after division.

```js
1 % 1
```

##### Exponentiation

Exponentiation is a mathematical operation that involves raising a number to a certain power. In programming, the exponentiation operator is denoted by ** (two asterisks).

```js
4 ** 4 // base ** exponent
```

#### Comparison operators

Comparison operators are used to compare two values and evaluate whether a condition is true or false.

##### Equality operator (==)

The equality operator compares two values and returns true if they are equal and false if they are different.

```js
5 == 5 // Returns true
5 == 6 // Returns false
```

##### Inequality operator (!=)

The inequality operator compares two values and returns true if they are different and false if they are equal.

```js
5 != 5  // Returns false
10 != 5 // Returns true
```

##### Greater than operator (>)

The greater than operator compares two values and returns true if the value on the left is greater than the value on the right.

```js
10 > 5  // Returns true
5 > 10  // Returns false
```

##### Less than operator (<)

The less than operator compares two values and returns true if the value on the left is less than the value on the right. 

```js
5 < 10  // Returns true
10 < 5  // Returns false
```

##### Greater than or equal to operator (>=)

The greater than or equal to operator compares two values and returns true if the value on the left is greater than or equal to the value on the right.

```js
10 >= 5  // Returns true
5 >= 10  // Returns false
5 >= 5   // Returns true
```

##### Less than or equal to operator (<=)

The less than or equal to operator compares two values and returns true if the value on the left is less than or equal to the value on the right.

```js
5 <= 10  // Returns true
10 <= 5  // Returns false
5 <= 5   // Returns true
```

### Data structures

Data structures are fundamental tools used in computer science and programming to organize and manipulate data efficiently. They provide a way to store and manage data in a structured format, enabling operations such as insertion, deletion, searching, and sorting. There are various types of data structures, each with its own characteristics and uses.

#### arrays

In programming, an array is a data structure that allows you to store multiple values of the same type in a sequential manner. Each value in an array is called an element, and each element is accessed by its index, which represents its position in the array.

##### Array Declaration

In Pikalang, arrays can be declared using square brackets `[]` and the desired length or by initializing it with elements enclosed in curly braces `{}`.

Example of array declaration:

```rs
var numbers = [1, 2, 3, 4, 5]
var fruits = ["apple", "banana", "orange"]
```

In the example above, we declared an array named `numbers` with five elements and an array named `fruits` with three elements.

##### Negative Indexes

In PikaLang you can use negative indices to access positions in an array. `-1` would represent the last position in the array.

```js
var arr = [ 1, 2, 3]
//         -3  -2 -1
arr[-1] = 4 // [ 1, 2, 4]
print(arr[-2]) // 2
arr[-20] = 123 // ERROR!
```

##### Accessing Array Elements

You can access individual elements in an array by specifying their index within square brackets `[]`. The index starts from 0 for the first element and increments by 1 for each subsequent element.

Example of accessing array elements:

```rs
var numbers = [1, 2, 3, 4, 5]
print(numbers[0]) // Output: 1
print(numbers[2]) // Output: 3

var fruits = ["apple", "banana", "orange"]
print(fruits[1]) // Output: banana
```

In the example above, we accessed the first element of the `numbers` array (`numbers[0]`) and the second element of the `fruits` array (`fruits[1]`).

##### Modifying Array Elements

You can modify the value of an element in an array by assigning a new value to a specific index.

Example of modifying array elements:

```rs
var numbers = [1, 2, 3, 4, 5]
numbers[2] = 10
print(numbers) // Output: [1, 2, 10, 4, 5]
```

In the example above, we modified the value of the third element in the `numbers` array from 3 to 10.

##### Array Length

The length of an array is the number of elements it contains. In Pikalang, you can obtain the length of an array using the `len()` function.

Example of getting the array length:

```rs
var numbers = [1, 2, 3, 4, 5]
var length = len(numbers)
print(length) // Output: 5
```

In the example above, we used the `len()` function to get the length of the `numbers` array, which is 5.

#### object

Objects are a fundamental data structure used to store and manipulate data. They consist of key-value pairs, where each value can be of any type, including other objects

To create an object, you can use the syntax of curly braces {} and define its properties and values within them. Here's a basic example:

```js
const val3 = 33

const obj = {
  val1: "Hi!!",
  val2: 3.14,
  val3,
  val4: {
    hello: "world"
  }
}
```

The properties in an object can be any valid string, and the values can be of any data type, such as numbers, strings, arrays and other objects.

To access the properties of an object, you can use dot notation `object.property` or bracket notation `object['property']`. Here's an example:

```js
print(object.property1) // Access a property using dot notation
print(object['property2']) // Access a property using bracket notation
```

You can also modify or add properties to an object at any time:

```js
object.property1 = newValue // Modify the value of an existing property
object.newProperty = newValue // Add a new property to the object
```

### Primitive data types

Primitive data types refer to basic or fundamental types of data that are built-in within a programming language. These data types are used to represent simple values and are typically not composed of other data types. In this document, we will explore four commonly used primitive data types: string, number, boolean, and null.

#### string

A string is a data type used to represent a sequence of characters. It can include letters, numbers, symbols, and whitespace. In most programming languages, strings are typically enclosed within single quotes ('') or double quotes (""). For example:

```js
"This is a string"
'This is another string'
```

#### number

The number data type is used to represent numeric values. It can include both integers (whole numbers) and floating-point numbers (decimal numbers). Numbers can be used for mathematical calculations, comparisons, and other numerical operations. For example:

```js
1234
1234.5
-1234
```

#### boolean

The boolean data type represents a logical value, which can be either true or false. Booleans are often used in programming to control the flow of code based on conditions. They are fundamental in decision-making processes and control structures such as if statements and loops. For example:

```js
true
false
```

#### null

Null is a special value that represents the absence of any object or value. It is typically used to indicate that a variable does not currently have a value assigned to it. Null is different from an empty string or zero, as it signifies the intentional lack of a value. For example:

```js
var name = null
```

### Native functions

The PikaLang language provides some predefined native functions to perform common tasks. These functions can be used directly without the need to define them beforehand.

#### `print()`

The print function is used to `print` a value to standard output. It takes an argument of any type and displays its representation in text form.

Example of use:

```py
print("Hi, Pika!!")
```

#### `len()`

The `len` function is used to obtain the length of a string.

Example of use:

```go
len("Hi, Pika!!") // This will return the number 10
```

#### `includes()`

The `includes` function is used to check if an array contains a specific element.

Example of use:

```go
var arr = [1, 2, 3, 4, 5]
includes(arr, 3) // This will return true
includes(arr, 6) // This will return false
```

#### `push()`

The `push` function is used to add elements to the end of an array.

Example of use:

```go
var arr = [1, 2, 3]
push(arr, 4, 5) // This will modify the array to [1, 2, 3, 4, 5]
```

#### `pop()`

The `pop` function is used to remove the last element from an array.

Example of use:

```go
var arr = [1, 2, 3, 4, 5]
pop(arr) // This will modify the array to [1, 2, 3, 4]
```

#### `shift()`

The `shift` function is used to obtain the first element from an array.

Example of use:

```go
var arr = [1, 2, 3, 4, 5]
shift(arr) // This will return [1]
```

#### `indexOf()`

The `indexOf` function is used to find the index of an element in an array.

Example of use:

```go
var arr = [1, 2, 3, 4, 5]
indexOf(arr, 3) // This will return 2
indexOf(arr, 6) // This will return -1
```

#### `isNaN()`

The `isNaN` function is used to check if a value is NaN (Not-a-Number).

Example of use:

```go
isNaN(10) // This will return false
isNaN(NaN) // This will return true
```

#### `isNull()`

The `isNull` function is used to check if a value is null.

Example of use:

```go
isNull(null) // This will return true
isNull(10) // This will return false
```

#### `prompt()`

The `prompt` function is used to display a message to the user and wait for input from the console.

Example of use:

```go
var name = prompt("Enter your name: ") // This will display "Enter your name: " and wait for user input.
```

#### `randNum()`

The `randNum` function is used to generate a random number within a specified range.

Example of use:

```go
randNum(1, 10) // This will return a random number between 1 and 10 (inclusive).
```

#### `pow()`

The `pow` function is used to calculate the power of a number.

Example of use:

```go
pow(2, 3) // This will return 8, as 2 raised to the power of 3 is 8.
```

#### `string()`

The `string` function is used to convert a value into a string representation.

Example of use:

```go
string(10) // This will return the string "10"
string(null) // This will return the string "null"
string(true) // This will return the string "true"
```

#### `num()`

The `num` function is used to convert a value into a numeric representation.

Example of use:

```go
num("10") // This will return the number 10
num("3.14") // This will return the number 3.14
```

#### `bool()`

The `bool` function is used to convert a value into a boolean representation.

Example of use:

```go
bool(0) // This will return false
bool("hello") // This will return true
bool(null) // This will return false
```

#### `toUpperCase()`

The `toUpperCase` function is used to convert a string to uppercase.

Example of use:

```go
toUpperCase("hello") // This will return "HELLO"
toUpperCase("WORLD") // This will return "WORLD"
```

#### `toLowerCase()`

The `toLowerCase` function is used to convert a string to lowercase.

Example of use:

```go
toLowerCase("Hello") // This will return "hello"
toLowerCase("WORLD") // This will return "world"
```

#### `capitalize()`

The `capitalize` function is used to capitalize the first letters of a string.

Example of use:

```go
capitalize("hello") // This will return "Hello"
capitalize("world") // This will return "World"
```

#### `startsWith()`

The `startsWith` function is used to check if a string starts with a specific prefix.

Example of use:

```go
startsWith("hello world", "hello") // This will return true
startsWith("hello world", "world") // This will return false
```

#### `endsWith()`

The `endsWith` function is used to check if a string ends with a specific suffix.

Example of use:

```go
endsWith("hello world", "world") // This will return true
endsWith("hello world", "hello") // This will return false
```

#### `reverseString()`

The `reverseString` function is used to reverse the characters in a string.

Example of use:

```go
reverseString("hello") // This will return "olleh"
reverseString("world") // This will return "dlrow"
```

#### `typeof()`

The `typeof` function is used to determine the type of a value.

Example of use:

```go
typeof(10) // This will return "number"
typeof("hello") // This will return "string"
typeof(true) // This will return "boolean"
```

#### `concat()`

The `concat` function is used to concatenate varius strings.

Example of use:

```js
concat("Hello", " ", "World!") // This will return "Hello World!"
```
