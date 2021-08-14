I'm exploring which interfaces boxing operations allocate.

Results:

```
go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/semistrict/go-experiments/interfacealloc
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIfaceAllocUInt64-8     	10196326	       116.2 ns/op	      71 B/op	       8 allocs/op
BenchmarkIfaceAllocInt32-8      	11765120	        99.98 ns/op	      35 B/op	       8 allocs/op
BenchmarkIfaceAllocByte-8       	65970087	        17.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceAllocInt16-8      	13021972	        91.34 ns/op	      17 B/op	       8 allocs/op
BenchmarkIfaceAllocRune-8       	11981491	        99.36 ns/op	      35 B/op	       8 allocs/op
BenchmarkIfaceAllocMyByte-8     	66166123	        17.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceAllocMyStruct-8   	 9552241	       123.8 ns/op	      10 B/op	      10 allocs/op
PASS
ok  	github.com/semistrict/go-experiments/interfacealloc	8.996s
```

Interestingly, all boxing operations I tested did allocate except for `byte` or 
types whose underlying type is `byte`.

I would not have expected this. I would have expected (in the best case) that Go
optimized boxing of data smaller than a pointer by inlining the value instead of
heap-allocating it and then storing a pointer in the interface.

Or I would have expected that no such optimization exists, so that every boxing
operation results in allocation. 

What we actually see is that there is some kind of optimization like this happening
but only for `byte`. Super weird! It can't be that common to box a single byte.

