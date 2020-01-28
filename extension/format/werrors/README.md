## Error Formatting 

For a given error chain `err1` - ... - `errN` with default `Error()` serialisation `"err1: ... : errN"`,
formatting allows alternative serialisations such as `"err1 - ... - errN"`, `"err1\n...\nerrN"`, as well as more advanced forms such as right-to-left formatting, JSON, protobuf, etc.

This is a draft implementation to demostrate the ease with which error formatting can be delivered under this proposal, 
but is not a complete implementation of it.
