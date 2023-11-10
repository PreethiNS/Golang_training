Type conversion:

    In Go, conversion refers to the process of changing the type of a value to another compatible type. It's a way to transform a value from one data type to another, enabling operations between different types or ensuring compatibility when types are mismatched.

    Explicit Conversion:

        Explicit conversion, also known as type casting, requires you to explicitly specify the conversion you want to perform.
        It's necessary when the data types involved are not compatible or when you want to be explicit about the conversion.

    Note: The Implicit conversion is not supported in Golang because of Type safety


Control Flow:


    Control flow in Go refers to how your program's execution is managed and directed based on conditions and logic. Go provides several control flow constructs, including conditionals and loops.
    It refers to the order in which instructions or statements are executed in a program. It determines how the program runs by allowing different paths or decisions to be taken based on certain conditions.

    1. if statement:

        if condition {
              // Code to execute if the condition is true
        }

        The if statement checks a condition. If the condition is true, the code block inside the curly braces is executed; otherwise, it is skipped.

    2. if-else statement:

        if condition {
            // Code to execute if the condition is true
        } else {
            // Code to execute if the condition is false
        }

        With if-else, if the condition is true, the code inside the first block is executed; otherwise, the code inside the else block is executed.

    3. if-else if statement:

        if condition1 {
            // Code to execute if condition1 is true
        } else if condition2 {
            // Code to execute if condition2 is true
        } else if condition3 {
            // Code to execute if condition3 is true
        } else {
           // Code to execute if none of the conditions are true
        }

        In an if-else if statement, it checks conditions one by one. If the first condition is true, it executes the code inside the first block; otherwise, it checks the next condition. If any condition is true, the corresponding block is executed. If none of the conditions are true, the code inside the else block is executed.

        Note:you can initialize a variable and specify the condition in the same line when using an if statement. This is a concise way to create a scoped variable that is only visible within the if block.

    4. Switch Statement

        In Go, the switch statement provides a way to express multi-way conditionals more cleanly than using a series of if-else statements. It evaluates expressions and compares them against possible values in different case blocks. 