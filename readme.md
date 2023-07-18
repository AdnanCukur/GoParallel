# GoParallel
## Introduction
GoParallel is a Go package that provides a convenient way to process a slice in parallel using specified maximum parallelism. It allows you to define the processing logic for each item in the slice and control the number of goroutines running concurrently.

Can be useful if you need to post a large number of items to a web service one by one, or in batches, or process a large number of files in a directory and you want to control the level of parallelism to not overload the webservice.

Inspired by the native Parallel libray in .Net
```csharp
Parallel.ForEach(sourceCollection, item => Process(item));
```
## Installation
```bash 
go get -u github.com/adnancukur/GoParallel
```

## Usage

Example of regular sequential processing without GoParallel
```go
for _, item := range sourceSlice {
    processFunction(item)
}
```

Processing in parallel using GoParallel, where maxParallelism is the maximum number of goroutines to run concurrently
```go
GoParallel.ProcessSlice(sourceSlice, processFunction, maxParallelism)
```
More examples in the example folder