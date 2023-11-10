Data Types:

1.Numeric Types:

    int8, int16, int32, int64:Signed integers of specific sizes.
    uint8, uint16, uint32, uint64: Unsigned integers of specific sizes.
    float32, float64: Floating-point numbers for representing real numbers.

2.Boolean Type:

            bool: Represents boolean values, which can be either true or false.

3.Character Types:

            byte: An alias for uint8. It represents a single 8-bit byte.
            rune: An alias for int32. It represents a Unicode code point.

4.String Type:

            string: Represents a sequence of characters (Unicode code points) and is often used to store text.

5.Composite Types:

            Array: A fixed-size collection of elements of the same type. The size is part of the array's type.
            Slice: A dynamically sized, flexible view of an array. Slices are more commonly used than arrays.
            Map: An unordered collection of key-value pairs, where keys are unique.
            Struct: A composite data type that groups together variables of different types under a single name.

6.Complex Types:

           complex64, complex128: Complex numbers with 32-bit and 64-bit real and imaginary parts.

Reference: https://go.dev/ref/spec#Types
