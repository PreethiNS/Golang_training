Functions: Part-2

    Variadic function:

        In Go, a variadic function is a function that can accept a variable number of arguments of the same type. This flexibility allows you to pass zero or more arguments of a specific type when calling the function.

        To declare a variadic function, use an ellipsis ... before the type of the last parameter in the function signature. Within the function, 
        the variadic parameter behaves like a slice.

    Underscore (_):

        In Go, the underscore _ is used as a placeholder for variables that you don't intend to use. It's a way to tell the compiler that a value is being ignored deliberately.

    Defer:

        defer is a Go keyword used to schedule a function call to be executed at the end of the current function's scope, typically before the function returns. It's often used for cleanup or to ensure that certain operations are performed regardless of whether the function encounters an error or executes successfully.

    Panic:

        panic is used to stop normal execution of a function in the case of an unrecoverable error. It immediately stops the current function execution and starts unwinding the stack, executing any deferred functions along the way.
        Once a panic occurs, the program stops executing the remaining code in the current function and moves towards termination.

    Recover:

        To execute the remaining code after a panic occurs, you need to recover from the panic within the same function or block where the panic happens.


Comma,ok:

        In Go, the "comma, ok" idiom is commonly used for error checking and type assertion when working with maps, channels, and other data structures. It allows you to check if an operation was successful and retrieve a value at the same time.

        value: The result or value you want to retrieve.
        ok: A boolean variable that indicates whether the operation was successful (true) or not (false).

        Usage:

            1.Map Lookups:

                In maps, you can use the "comma, ok" idiom to check if a key exists and retrieve its associated value.

