## Error Stack Frames 

For an wrapped error produced by a call to `Wrap` at file `caller.go` at line `N`,
stack frames add this information to the `Error()` output, for example as `"err: caller.go(N)"`.

This is a draft implementation to demostrate the ease with which error stack frames can be delivered under this proposal, 
but is not a complete implementation of it.
