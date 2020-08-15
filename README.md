# Cscr
#### Compact scripting language

<b>Note:</b> This is a hobby project, and should not be used within 200m of a production server.
It has many bugs and edge cases that break the lexer, parser and/or evaluator.

## Usage:
1) To build the cli, build the <code>/cmd/cscr/main.go</code> file
2) To run some code, run <code>main yourCode.cscr</code> in the folder with the built executable.
(assuming the build process produced an executable named 'main')

## What is it
Cscr is an interpreted, not-even-typed, start of a simple scripting language.
It's written entirely in Go.

Cscr currently has only numbers (decimal and integer), and text as variable names.
You can create variables, calculate expressions, assign them as values, and even assign variables as values! Oh the joy!

Cscr currently supports basic mathematical 
and logical operations with parenthesis, converting true to 1 and false to 0 (negation is buggy tho)

The parser uses the [shunting-yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
to parse expressions, and so should follow the proper mathematics order of operations.

## Features
#### No scope
Cscr does not have the context or scope, so after a variable is declared, it can be used anywhere.
This even works on the same line of code, and so you can declare variables inside a variable declaration. Funky, right?

#### Return values
In cscr, every operation and expression returns either a value or a reference.
When the program is finished, it will output the value of the last line of code to the terminal.

An expression will return the evaluated value of that expression.
Example:

    // this expression will return 2
    (1 + 1);
    // this expression will return a reference to a, which when evaluated is 3
    (a := 1 + 2)
    // you can also declare a variable inside a variable declaration:
    // b's value will be a reference to 'a'
    b := (a := 1 + 1)
    // changing a will affect b
    a += 1;
    // this will return 1 (=true)
    b == 3
    

Declaring a variable will return a reference to the created variable.

#### Comments
Cscr supports basic, one-line comments.
Lines starting with a double slash '//' will be ignored by the parser.
If the line does not start with a comment, it will be parsed as code all the way through.

#### Values
Values in Cscr can be either values or references.
If a variable is assigned to another variable, a reference to that variable will be assigned.
However, if a reference is used in a mathematical or logical operation,
it will be converted to a value by getting the value of the reference

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
    
Example of converting a reference to a value:
    
    //declare 'a'
    a := 1;
    // now we're declaring b to be the value of the expression (a + 0), which is 1
    b := a + 0;
    // mutating 'a' won't affect 'b' anymore
    a += 1;
    // check it out
    b == a
    // == 0 (false)
