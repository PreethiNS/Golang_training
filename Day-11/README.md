Format specifiers:
    Format specifiers are used in functions like Printf or Sprintf from the fmt package to control how values are formatted when printing or helps in convertion.

    A format specifier starts with a percent sign % followed by a letter that represents the type of value you want to format. 

        For example:
            %d is used for formatting integers.
            %f is used for formatting floating-point numbers.
            %s is used for formatting strings.

        Integer formatting:

            %d: Formats an integer as a decimal number.
            %b: Formats an integer in binary (base 2).
            %o: Formats an integer in octal (base 8).

        String formatting:

            %s: Formats a string.
            %q: Formats a string in double quotes.
            %v: Formats the value in its default format.

            In Go, %v and %+v are format specifiers used with the Printf function from the fmt package to format values into strings.

                %v: It's used to format a value in its default format. This specifier adapts its output based on the value's type. For structs, it prints the default format of each field.

                %+v: This format specifier is specifically used with structs. It prints the struct's fields along with their names. It shows the field names and values.

            In Go, Printf and Sprintf are functions provided by the fmt package for formatting strings.

            Printf: This function is used to print formatted text to the standard output (usually the console).
            Sprintf: This function is similar to Printf, but instead of printing the formatted string to the standard output, it returns the formatted string. It takes the same arguments as Printf but returns the resulting string instead of printing it.


Functions: Part-1

        Functions are a fundamental building block in Go (Golang) programming. They allow you to encapsulate a block of code that can be called and reused throughout your program.

        Functions are essential in programming for several reasons:

                Modularity
                Code Reusability
                Code Organization
                Debugging and Troubleshooting


        Function Declaration:

                In Go, you declare a function using the func keyword, followed by the function's name, a parameter list, the return type (if any), and the function body enclosed in curly braces.

                functionName: The name of the function.
                returnType: The data type of the value that the function returns.
                result: The value that the function returns (if the return type is specified).
                parameters:  Parameters are placeholders or variables listed in the function definition. They define the types and order of the values that a function expects to receive when it's called. 
                Arguments: Arguments are the actual values or expressions passed to a function when it is called. They are the real values that are assigned to the parameters defined in the function.

        Anonymous function:
        
                Anonymous functions are functions that are defined without a name. They are also known as "function literals" or "closures." Unlike regular functions, which are declared with a name and can be called later by that name, anonymous functions are defined inline and can be executed immediately or assigned to variables.
                Anonymous functions are useful in scenarios where you need a small, one-time-use function or when you want to define functions dynamically within your code, such as passing a function as an argument to another function or defining a function inside a loop.

