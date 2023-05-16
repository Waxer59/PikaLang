<p align="center">
  <img src="./docs/img/logo.webp" height="250px"/>
</p>

# PikaLang

Pika, or Pikalang, is a programming language designed to be simple, efficient and easy to use. It is written in Go and its syntax is inspired by several languages. Pika uses the `.pk` extension for its files.

- [PikaLang](#pikalang)
  - [CLI](#cli)
  - [Syntax](#syntax)
    - [Variable declaration](#variable-declaration)
    - [Function declaration](#function-declaration)
    - [Comments](#comments)
      - [Single-line Comments](#single-line-comments)
      - [Multi-line Comments](#multi-line-comments)
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

### Variable declaration

In Pikalang, variables are declared using the `var` keyword followed by the variable name and its type. For example:

```js
var foo = "bar"
var bar = 42
```

In addition, you can also declare constants using the `const` keyword. For example:

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