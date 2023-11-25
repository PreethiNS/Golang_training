Methods:

        In Go, methods are functions associated with a particular type. They allow you to attach behavior to a type by defining functions that can be called on instances of that type.

        Explanation:

            Receiver : A method is declared with a receiver, which is a parameter specifying the type that the method operates on. It allows you to call the method on instances of that type.
            Syntax : func (receiver Type) methodName() { ... }
            Invocation : Methods are invoked using the dot notation on an instance of the associated type.

        Built-in Types:

            However, methods can't be directly attached to built-in types (int, float, etc.). For example, you can't add a method to the built-in int type in your code.

        Wrapping Built-in Types:

            create new types by wrapping built-in types using aliasing or creating new types, and then define methods on these new types.

        Receiver Types:

            Value Receiver: Methods with a value receiver operate on a copy of the value, meaning modifications don't affect the original instance.

            Pointer Receiver: Methods with a pointer receiver operate on the actual value itself, allowing modifications to affect the original instance.

Error Handling:

            Error handling in Go is based on returning error values from functions and using a specific error-checking pattern. 
            This pattern allows you to gracefully handle errors, making your code more robust.

        Returning Errors:

            Functions that may encounter errors return them as an additional return value. By convention, errors are often
            returned as the last value from a function.

            You call a function that returns a result and an error.
            You check if the error is not nil. If it's not nil, an error occurred.
            If an error occurred, you handle it (e.g., print an error message, log it, or take corrective action).
            If there's no error, you use the result.