Types of Variables:

1.Local Variables:

            * Local variables are declared within a function and have a limited scope. They are only accessible within the block of code where they are declared.                                * These variables are created when the function is called and destroyed when the function  exits.

2.Package-Level Variables:

            * Package-level variables are declared outside of any function and are accessible throughout the entire package.
            * They have global scope within the package but are not visible to code outside the package unless they are exported (start with an uppercase letter).

3.Exported Variables:

            *Exported variables are package-level variables that start with an uppercase letter. They are accessible outside the package where they are defined.

---------------------------------------------------------------------------------------------------------

Constants:

           Constants are variables whose values cannot be changed once they are assigned. They are often used for values that should not be modified during 
           program execution.
           Constants can be local or package-level and are declared using the const keyword.

---------------------------------------------------------------------------------------------------------

Pointers:

Pointers are a fundamental concept in Go and are used to store memory addresses, enabling indirect access to values stored in memory. Understanding pointers is crucial for various programming tasks, such as passing values by reference and working with memory allocation.

        var ptr *int

    Assigning Pointers:

        You can assign the memory address of a variable to a pointer using the address-of operator &.
        ptr = &x

    Dereferencing Pointers:

          To access the value pointed to by a pointer, you use the dereference operator *.
          value:= *ptr

    Passing Pointers to Functions:

          Pointers are often used to pass values by reference to functions, allowing modifications to the original variable.
