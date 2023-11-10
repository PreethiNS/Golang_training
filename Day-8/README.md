break Statement:

        The break statement is used to exit the loop immediately when encountered. It is often used to prematurely terminate a loop based on a certain condition.

Continue Statement:

        The continue statement is used to skip the current iteration and move to the next iteration of the loop without executing the remaining code within the loop for that specific iteration.


Composite data types:

        They are used to group together multiple values into a single unit or structure. These composite types allow for the creation of more complex and organized data structures, aiding in the management and manipulation of data in a program.

        Arrays and Slices

            Arrays and slices are data structures used to store collections of elements. They have some differences in terms of size and flexibility.

            Arrays:

                *An array is a fixed-size collection of elements of the same type.
                *The size of an array is determined at the time of declaration and cannot be changed.
                *Arrays have a zero-based index, meaning the first element is at index 0, the second at 1, and so on.
                *Arrays in Go are not naturally represented as pointers.
                *If you pass an array as an argument to a function, a copy of the entire array is made (pass-by-value).
                *Modifying the elements of the copied array inside the function does not affect the original array.

            Slices:

                *A slice is a dynamically sized, flexible view of an underlying array.
                *Slices have no fixed size and can grow or shrink as needed.
                *They are more commonly used than arrays in Go due to their flexibility.
                *Slices are created using the make function or by slicing an existing array, slice, or string.
                *Slices in Go are naturally represented as pointers to the underlying array.
                *When you pass a slice as an argument to a function, you are passing a reference to the same underlying array (pass-by-reference).
                *Modifying the elements of the slice inside the function affects the original slice.