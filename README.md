# Cscr
#### Compact scripting language

<b>Note:</b> This is a hobby project, and should not be used within 200m of a production server.
It has many bugs and edge cases that break the lexer, parser and/or evaluator.

## Usage:
1) To build, run
<code>go build</code> in the project's root directory.
or build the <code>/cmd/cscr/main.go</code> file
2) To run code, run <code>main <path_to_code></code> in the folder with the built executable.
(assuming <code>go build</code> produced an executable named 'main')

## What is it
Cscr is an interpreted, not-even-typed, start of a simple scripting language.
It's written entirely in Go.

Cscr currently has only numbers (decimal and integer), and text as variable names.
You can create variables, assign them values, and reassign those values. Oh the joy!

Cscr currently supports basic logical operations (excluding negation (!)) by converting true to 1 and false to 0

The parser uses the [shunting-yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
to parse expressions, and so should follow the proper mathematics order of operations.

## Features
#### Return values
In cscr, every operation and expression returns a value or a reference,
and the last operation of an expression will always be returned.
The whole program is one big expression, so the value of the
last line of the program will be printed after execution is done.

There can be multiple expressions or operations in an expression: 
<code>(a := 1; a + 1)</code> would evaluate to 2 outside the parenthesis, since a + 1 == 2.
Declaring a variable will return a reference to the created variable.

#### Comments
Cscr supports basic, one-line comments.
Lines starting with a double slash '//' will be ignored by the parser.
If the line does not start with a comment, it will be parsed as code all the way through.

#### Values
Values in Cscr can be either constants or references.
If a variable is assigned to another variable, a reference to that variable will be assigned.
However, if a reference is used in a mathematical or logical operation, it will be converted to a constant value

Example of a reference:

    //declare 'a'
    a := 1;
    // declare 'b' with a reference to 'a'
    b := a;
    // mutate 'a'
    a += 1;
    // 'b' is also affected
    b == a
    // == 1 (true)
    
Example of converting a reference:
    
    //declare 'a'
    a := 1;
    // now we're declaring b to be the value of the expression (a + 0), which is 1
    b := a + 0;
    // mutating 'a' won't affect 'b' anymore
    a += 1;
    // check it out
    b == a
    // == 0 (false)