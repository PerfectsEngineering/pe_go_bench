# Benchmark Tests in Golang

This is a benchmark codebase accompanying a blog post about benchmarking in Golang. 

The codebase contains a formatter package that formats a string by replacing all occurrences of a given substring with another string. 
To see the Benchmark code, check out the `formatter/formatter_test.go` file. However, you should read the accompany blog post to have 
better understanding of the benefits of benchmarking your code.

## My Benchmark Results

Results of running the benchmarks on my computer (M2 Pro, 16GB):

```sh
BenchmarkNormalise/WithReplacer-10              50672794               228.3 ns/op            32 B/op          2 allocs/op
BenchmarkNormalise/WithReplacer-10              52982422               229.2 ns/op            32 B/op          2 allocs/op
BenchmarkNormalise/WithReplacer-10              52576513               228.2 ns/op            32 B/op          2 allocs/op
BenchmarkNormalise/WithReplacer-10              52496136               228.9 ns/op            32 B/op          2 allocs/op
BenchmarkNormalise/WithReplacer-10              52274142               229.7 ns/op            32 B/op          2 allocs/op
BenchmarkNormalise/WithRegexp-10                26651758               451.6 ns/op            56 B/op          4 allocs/op
BenchmarkNormalise/WithRegexp-10                26442408               452.4 ns/op            56 B/op          4 allocs/op
BenchmarkNormalise/WithRegexp-10                26480550               452.6 ns/op            56 B/op          4 allocs/op
BenchmarkNormalise/WithRegexp-10                26595258               453.2 ns/op            56 B/op          4 allocs/op
BenchmarkNormalise/WithRegexp-10                26478879               452.6 ns/op            56 B/op          4 allocs/op
```

Clearly the `WithReplacer` version is faster by about 50% (even though the WithReplacer still does a matcher check). 
Another reason why you should be careful when using regex. 

## Running the benchmarks

To run the benchmarks, use the following command:

```sh
go test -bench . -count 5 -benchmem  -benchtime 10s ./...
```

This will run all benchmarks in the codebase and print out the results, including memory allocations.

Feel free to use this codebase as a reference for your own benchmarking needs. If you have any questions or suggestions, please feel free to open an issue or a pull request. 

