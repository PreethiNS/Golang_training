Go structure:

Module

        *A module in Go refers to a collection of Go packages
        *A module is defined by a go.mod file located at the root of the project directory. The go.mod file specifies the module's name, version, and dependencies.   
         
Package

        *A package in Go is a collection of Go source files that are organized together in the same directory.
        *Packages serve as a unit of encapsulation and code organization. They allow you to group   related functions,types, and variables together.                                             
        *Packages are identified by a package declaration at the top of each Go source file. 
        For example, ‘package main’ or ‘package mypackage’.   
                                           
Import

        *In Go, the import statement is used to include packages from other modules or your own project.
        *It allows you to use functions, types, and variables defined in other packages in your code.
        Import paths specify the location of the package, such as "fmt" or "myproject/mypackage"

Main function

        *The main function is a special function in Go that serves as the entry point for executable programs.
        *The main function is defined as follows: func main() { ... }.

---------------------------------------------------------------------------------------------------------
Go Playground:

The Go Playground is an online environment that allows you to write, run, and share Go code without installing any software. 
It's a web-based platform that enables users to experiment with the Go programming language.

Key features:

1.Coding Environment
2.Execution
3.Error Checking
4.Sharing
5.Learning

---------------------------------------------------------------------------------------------------------
Variables:

In Go, a variable is a container for storing data. Variables allow you to store and manipulate values such as numbers, text, and more within your program.
It is used to store and manage information, making it easier to work with data and perform various tasks.

Three ways of decalring variable:

    1.Using Var
    2.Using Type inference with var
    3.short variable decalration
