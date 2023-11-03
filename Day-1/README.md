What is Programming?

    programming is the process of crafting a detailed set of instructions in a specific programming language to enable a computer to carry out tasks or solve problems in a systematic and controlled manner.


Typical steps in programming world:

    Define the Problem
    Plan Your Solution
    Select a Programming Language
    Write the Code
    Test Your Code
    Debugging
    Optimize and Refine
    Document Your Code
    Version Control
    Deployment
    Maintenance
    Documentation and User Manuals


Your code to machine code:
    The process of translating English statements or any high-level programming language into machine-understandable code, which ultimately consists of 0's and 1's, involves several steps:
    
            1.High-Level Language Code
            2.Compilation/Interpretation
            3.Intermediate Code
            4.Machine Code
            5.Execution


Compilation process:

    You use a compiler that reads the entire source code and translates it into machine code (binary).
    The compiler checks for syntax errors and translates the entire code into a binary executable.
    This binary file contains the instructions to perform specific task based on your code.
    You run the compiled program, and it executes the instructions directly, without needing to translate each line at runtime.


Interpretation process:

    You use a interpreter that reads and executes your code line by line at runtime.
    The interpreter reads the first line, interprets it, and performs the operation.
    It then proceeds to the next line, interpreting and executing it in sequence.
    Each line is translated to machine code dynamically as the program runs.


Go Introduction:

    Go, often referred to as Golang, is a modern, statically typed programming language known for its simplicity, efficiency, and strong support for concurrency. It was developed by Google engineers and released in 2009.


Features of Go:

    Simplicity: Go has a clean and concise syntax that makes it easy to read and write code. It avoids unnecessary complexity and provides a straightforward approach to programming.

    Efficiency: Go is designed to be fast and efficient, both in terms of compilation and execution. It's an excellent choice for building high-performance applications.

    Concurrency: Go has built-in support for concurrent programming. Goroutines, lightweight threads, allow developers to write concurrent code easily. Channels provide a safe way for goroutines to communicate and synchronize.

    Strong Typing: Go is statically typed, which means that variable types are defined at compile-time. This helps catch errors early in the development process.

    Garbage Collection: Go includes automatic memory management (garbage collection), which simplifies memory handling and reduces the risk of memory-related bugs.

    Standard Library: Go comes with a rich standard library that covers a wide range of tasks, from networking and web development to file handling and more. This eliminates the need to rely heavily on third-party libraries for common tasks.

    Cross-Platform: Go is designed to be portable, making it easy to write code that works on various operating systems and architectures.

    Open Source: Go is open source, and its development is driven by the Go community. This encourages collaboration and innovation.


Why new language?

    Go was born out of frustration with existing languages and environments for the work was doing at Google. Programming had become too difficult and the choice of languages was partly to blame. 
    One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. 
    Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.


Solution: Go

    Go addressed these issues by attempting to combine the ease of programming , dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer.
