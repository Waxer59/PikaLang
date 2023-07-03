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
    - [Function declaration](#function-declaration)
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

### Function declaration

In Pikalang, functions are defined using the keyword `fn`, followed by the function name, the parameters in parentheses and the return type. For example:

```rs
fn add(x,y) {
    x + y
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