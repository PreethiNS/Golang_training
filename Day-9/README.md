Maps:

    In Go , a map is a built-in data structure that allows you to store and retrieve key-value pairs. Each key in a map is unique, and you can use it to access its associated value. Maps are similar to dictionaries

1.Creating a Map:

    You can create a map in Go using the make function or a map literal.

2.Adding and Accessing Elements:

    You can add key-value pairs to a map using the syntax map[key] = value, and you can access values using the key

3.Deleting Elements:

    You can delete a key-value pair from a map using the delete function

4.Checking for Existence:

    To check if a key exists in a map, you can use a two-value assignment with the map access operation.

5.Iterating Over a Map:

    You can iterate over the key-value pairs of a map using a for loop.

6.Maps as Reference Types:

    Maps in Go are implemented as reference types. When you assign a map to another variable or pass it as an argument to a function, both variables (the original and the copy) refer to the same map in memory. Changes made to the map through one variable will affect the other.