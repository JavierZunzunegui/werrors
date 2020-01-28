## Efficient Error Serialisation 

This presents an alternative implementation of `WrapError.Error` that is more efficient for long error chains.

This is a draft implementation to demostrate the ease with which optimisations to error serialisation can be delivered under this proposal, 
but is not a complete implementation of it.

The results of the benchmark in this implementation:

```
go test ./... -test.bench=Bench -test.benchmem
goos: linux
goarch: amd64
pkg: github.com/JavierZunzunegui/werrors/extension/optimize/werrors

# current implementation
BenchmarkWrapError_LegacyError/depth-1-4         	18647126	      60.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkWrapError_LegacyError/depth-2-4         	 9439263	       124 ns/op	      48 B/op	       2 allocs/op
BenchmarkWrapError_LegacyError/depth-3-4         	 6313530	       191 ns/op	      96 B/op	       3 allocs/op
BenchmarkWrapError_LegacyError/depth-5-4         	 3576243	       331 ns/op	     208 B/op	       5 allocs/op
BenchmarkWrapError_LegacyError/depth-10-4        	 1560608	       761 ns/op	     720 B/op	      10 allocs/op
BenchmarkWrapError_LegacyError/depth-20-4        	  533863	      1932 ns/op	    2592 B/op	      20 allocs/op

# alternative implementation
BenchmarkWrapError_Error/depth-1-4               	10047628	       101 ns/op	      80 B/op	       2 allocs/op
BenchmarkWrapError_Error/depth-2-4               	 9173968	       127 ns/op	      96 B/op	       2 allocs/op
BenchmarkWrapError_Error/depth-3-4               	 7820640	       156 ns/op	     112 B/op	       2 allocs/op
BenchmarkWrapError_Error/depth-5-4               	 5998074	       199 ns/op	     128 B/op	       2 allocs/op
BenchmarkWrapError_Error/depth-10-4              	 2499643	       482 ns/op	     336 B/op	       3 allocs/op
BenchmarkWrapError_Error/depth-20-4              	 1219206	       963 ns/op	     736 B/op	       4 allocs/op
```
