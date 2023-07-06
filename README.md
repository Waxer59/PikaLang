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
    - [Function declaration](#function-declaration)
    - [Ternary operator](#ternary-operator)
    - [Switch statement](#switch-statement)
      - [Multiple Cases](#multiple-cases)
    - [Logical Cases](#logical-cases)
    - [Comments](#comments)
      - [Single-line Comments](#single-line-comments)
      - [Multi-line Comments](#multi-line-comments)
    - [Operators](#operators)
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
      - [object](#object)
    - [Primitive data types](#primitive-data-types)
      - [string](#string)
      - [number](#number)
      - [boolean](#boolean)
      - [null](#null)
    - [Native functions](#native-functions)
      - [print()](#print)
      - [len()](#len)

## CLI

The Pikalang CLI (Command Line Interface) provides a convenient way to interact with the Pikalang interpreter and run programs written in this language.

```bash
COMMANDS:
   run   Run a file
   help  Show help
   repl  Start the repl

GLOBAL OPTIONS:
   --version, -v  print the version
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
    print("x is greater than 5");
}
```

* Second form: if condition { }
This form omits the parentheses around the condition and directly includes the code block.

Example:
```js
if x > 5 {
    print("x is greater than 5");
}
```

#### Else Statement
The 'else' statement follows an 'if' statement and is used to specify a block of code that should be executed if the condition in the preceding 'if' statement evaluates to false. The 'else' statement does not require any conditions. It is optional and can be omitted if not needed.

Example:
```js
if (x > 5) {
    print("x is greater than 5");
} else {
    print("x is not greater than 5");
}
```

#### Else If Statement:
The 'else if' statement allows for the evaluation of multiple conditions in a series of sequential checks. It is used when there are more than two possible outcomes based on different conditions. Multiple 'else if' statements can follow an 'if' statement, and the code block associated with the first true condition is executed. Only one code block will be executed, even if multiple conditions are true. Similar to the 'if' statement, the 'else if' statement supports the two syntax forms:

* First form: else if (condition) { }
This form encloses the code block within curly braces immediately after the condition.

Example:
```js
if (x > 5) {
    print("x is greater than 5");
} else if (x < 5) {
    print("x is less than 5");
}
```

* Second form: else if condition { }
This form omits the parentheses around the condition and directly includes the code block.

Example:
```js
if (x > 5) {
    print("x is greater than 5");
} else if x < 5 {
    print("x is less than 5");
}
```

Please note that nested 'if' statements are supported, allowing the inclusion of further 'if,' 'else if,' or 'else' statements within the code blocks.

### Function declaration

In Pikalang, functions are defined using the keyword `fn`, followed by the function name, the parameters in parentheses and the return type. For example:

```rs
fn add(x,y) {
  // Do something...
}
```

### Ternary operator

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

### Logical Cases

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

var x = 10;
```

> Remember that comments are for humans reading the code, so write clear and concise comments that enhance code readability and maintainability.

### Operators

Operators are symbols or characters used in programming languages to perform operations on variables, values, or expressions. They are used to manipulate and compare data, control program flow, and perform logical operations.

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
print(object.property1); // Access a property using dot notation
print(object['property2']); // Access a property using bracket notation
```

You can also modify or add properties to an object at any time:

```js
object.property1 = newValue; // Modify the value of an existing property
object.newProperty = newValue; // Add a new property to the object
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

#### print()

The print function is used to `print` a value to standard output. It takes an argument of any type and displays its representation in text form.

Example of use:

```py
print("Hi, Pika!!")
```

#### len()

The `len` function is used to obtain the length of a string.

Example of use:

```go
len("Hi, Pika!!") // This will return the number 10
```