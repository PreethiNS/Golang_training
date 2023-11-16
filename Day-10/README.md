Structs:

    struct is a composite data type that allows you to group together variables of different types under a single name. 
    It enables you to create custom data structures that represent real-world entities or complex data relationships.
    A struct is defined by specifying the names and types of its fields. It acts as a blueprint for creating instances or objects of that particular structure.

struct tags:

    struct tags are annotations added to the fields of a struct.
    The common use cases for struct tags include JSON encoding/decoding, database mapping, and other forms of serialization and deserialization.

    simple usecase:

    JSON data is mapped to struct fields and then stored in a database. This process often involves multiple steps:


            1.JSON Data Retrieval:JSON data might be received from an external API, a client request, or another source.
            2.Unmarshaling JSON to Structs: The JSON data is unmarshaled (decoded) into Go structs. This process involves mapping the JSON keys to the corresponding struct fields. If needed, struct tags, like json, are used to specify the mapping explicitly.
            3.Data Manipulation: Once the JSON data is transformed into Go structs, it can be manipulated or processed as needed, 
            such as performing validations, applying business logic, or modifying the data structure.
            4.ORM (Object-Relational Mapping) or Database Interaction: If the intention is to store this data in a database, an ORM or direct database interaction is used. ORM libraries like GORM, SQLBoiler, or others in Go provide functionality to map these Go structs to database tables, handle CRUD (Create, Read, Update, Delete) operations, and perform actions like creating or updating database records.
            5.Marshaling Structs to Database Records: Before storing data in the database, the Go structs might be marshaled (encoded) into a format suitable for the specific database backend.
            6.Database Storage: Finally, the data, usually in the form of SQL statements or commands, is executed against the database to insert or update the records.


If you don't use struct tags in Go, you can still perform marshalling and unmarshalling of data, but it might not be as automatic or seamless.you would need to manually create functions/methods to handle this conversion. Struct tags are optional metadata that helps in automatically mapping fields between Go structs and external representations like JSON, database columns, etc.

Embedded structs:

    you can create a new struct by embedding one or more existing structs. This is known as "embedding" or "composition" in Go. When you embed a struct within another struct, the fields and methods of the embedded struct are promoted to the outer struct, making them accessible through the outer struct.

Comments:
        There are two types of comments in Go:

        1.Single-line comments: These start with // and continue until the end of the line

        2.Multi-line comments: These start with /* and end with */. They can span multiple lines

        Comments in programming are incredibly useful for several reasons:

            1.Documentation
            2.Clarity and Readability
            3.Debugging and Troubleshooting

Note: Over-commenting or leaving outdated comments can clutter the code and make it harder to read. Comments should be clear, concise, and focused on explaining the why, not just the what, to maximize their effectiveness in aiding understanding and maintenance of the codebase.