# Concurrency in Go
### Concepts from Communicating Sequential Processes (CSP), channels and select statements, add value as well
- M:N scheduler, which means it maps M green threads to N OS threads.
> When we have more goroutines than green threads available, Go's runtime distributes the goroutines to new OS thread

- Goroutine is like Thread, but more effiecient thanks to Go's Runtime
> - Many goroutine can be multiplexed on fewer OS's threads m:n where 
m = Goroutines ; n = threads

- TLDR: 1 Goroutine combines concurrency and parallel by Go's runtime, so 1 Goroutine may NOT = 1 Thread, but behave like it. 
>  In Go,we would create a goroutine for each incoming connection, and then return from the goroutine’s function

> A goroutine is a function that is running concurrently



- Fork-join model 
   - Fork: split child branch from parent 
   - Join: Merge child branch with parent

Example
```
sayHello := func() {
   fmt.Println("hello")
}

go sayHello()
```
- Problem:
   - ```main()``` will be executed finished before ```sayHello()``` can run
- Solution:
   - Method 1: Add a time.sleep() after ```go sayHello()```, but it doesn't create join point, only race condition?
   - Method 2: Create a join point by synchronize the main goroutine and ```sayHello()``` goroutine using ```sync.WaitGroup```
   