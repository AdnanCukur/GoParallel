# GoParallel
## Introduction
GoParallel is a Go package that provides a convenient way to process a slice in parallel using specified maximum parallelism in only one line of code. It allows you to define the processing logic for each item in the slice and control the number of goroutines running concurrently.

Can be useful if you need to post a large number of items to a web service one by one, or in batches, or process a large number of files in a directory and you want to control the level of parallelism to not overload the webservice.

Currently used weekly in production to post hundreds of thousands of items to a third party web service that only allows one item to be posted at a time.


Inspired by the simplicity of the native Parallel libray in .Net
```csharp
Parallel.ForEach(sourceCollection, item => Process(item));
```

This package makes use of generics to allow you to process any type of slice.
## Installation
```bash 
go get -u github.com/adnancukur/goparallel
```

## Usage


Processing in parallel using GoParallel, where maxParallelism is the maximum number of goroutines to run concurrently

Example 1
```go
goparallel.Process(sourceSlice, processFunction, maxParallelism)
```
Example 2
```go

sourceSlice := []*TestStruct{
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: true},
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: true},
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: false},
    {IHaveBeenProcessed: false},
}

maxParallelism := 5

goparallel.Process(
    sourceSlice, 
    func(testStruct *TestStruct) 
    {
        testStruct.IHaveBeenProcessed = true
    }, 
    maxParallelism
    )
```
More examples in the example folder