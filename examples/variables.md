# Variables
variables in cscr can be either values or references.
Once declared, their type cannot be changed.

## creating a variable
variables are declared using the <code>:=</code> (walrus) operator

Examples:

    // no initial value
    myVar :=;
    
    // myVar is a number
    myVar := 5;
        
    // myVar is a function
    myVar := func() {
        /* function code */
    }

    // myVar is an object
    myVar := {
        foo: "bar"
    }
    
## changing a variable's value
Changing a variable's value is done using
the <code>=</code> operator.

The new value must be of the same type as the value that
was initially given to the variable

Examples:

    // declaring a variable of type 'number'
    myVar := 5;
    
    // allowed, since 123 is a number
    myVar = 123;
    
    // not allowed, since "hello" is a string
    myVar = "hello";