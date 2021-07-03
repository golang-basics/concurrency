# WaitGroup(s) - `sync.WaitGroup`

### The Problem

In the real world of applications, not every operation inside a program executes concurrently,
however many operations will most probably do. A lot of the times though, synchronous code
needs a certain **condition to be true** in order to proceed, thus **waiting** for a certain amount
of **concurrent actions** to complete first.

The **order** in which concurrent code executes is always **non-deterministic**, thus the way in which
the `go` keyword **appears**, does not mean that's the order in which way the **Go Scheduler** will
**schedule** them or that, that's the order in which they will **actually** get to **run**.
That being said, sometimes there is a need to **preserve order**, but still **execute** our
operations **concurrently**.

Also in a real world, a service does not always run standalone, meaning sometimes our applications
rely on other **downstream applications** in order to produce a result. In many cases we do not control
those downstream applications, and many of those applications may have **API Rate Limiting**, which means
our upstream application has to be careful on **how many requests** it makes, and at which **rate** they happen.

So we need some kind of **waiting mechanism** for **concurrent code**. We need to **not wait too long**,
so that our system is **efficient** and makes the best use of our resources, but also **wait just enough**,
so that the **order is preserved**, certain **conditions are met**, and we have control over
the amount of concurrent operations happening in our system.

Luckily there's no need to reinvent the wheel, Go introduces the `sync.WaitGroup` type,
which does exactly all the above mentioned.

### Go Concurrency Philosophy

Go is a powerful language, and it gained its popularity because people hear
its **Concurrency Model** is one of the best. While Go has quite an elegant way
of doing Concurrency, Go did not give up the Classic way of doing Concurrency.

In Go there are **2 ways** of making sure **concurrent code** executes **correctly** and **safely**:

1. `Concurrency Primitives`
2. `Channels`

Very important to stress out, it's not one versus the other, and there's no better
way of executing concurrent code using one or the other. It all depends on the
type of scenario and use case. Generally **everything** that can be achieved with **channels**
can be achieved with the `sync` **Concurrency Primitives** as well. It's a matter of choice depending on
**Complexity**, **Composition**, **Performance** or in general on what's being done with the data.

### Concurrency Primitives

All of Go's Concurrency Primitives are stored inside the `sync` package, which stands for
**Synchronization**, because most of the times that's what we as developers do with the
**executing concurrent code**.

These are all the available types under the `sync` package:

- `WaitGroup`
- `Mutex`
- `RWMutex`
- `Locker`
- `Cond`
- `Map`
- `Pool`

### WaitGroup Overview

In order to achieve tasks like:

- **Waiting** on a **Condition** (concurrent operations) to finish
- **Rate Limiting** the amount of max Concurrent Operations
- **Preserve** Concurrent Operations **order**

and a bunch of other things, we can simply make use os the `sync.WaitGroup` type.

Using a `sync.WaitGroup` is fairly simple, all we need to do is make sure to:

1. Create a `sync.WaitGroup`
2. Call the `Add()`
3. Call the `Done()` method inside each concurrent operation, once it's done
4. Call the `Wait()` method at the waiting point

Using a `sync.WaitGroup` in your code may cause a wide variety of issues,
this is why here are some golden rules I recommend everyone who works with WaitGroup(s)

- `Done()` MUST be called as many times as `Add()`
- If calls to `Done()` are less than calls to `Add()` => deadlock
- If calls to `Done()` are more than calls to `Add()` => panic
- Calling `Wait()` without calling `Add()` will return immediately
- `sync.WaitGroup` MUST always be passed by **reference** (as pointer) => (possible panic)
- Calling another `Wait()` before the previous one returns => panic
- Call `Add(n)` when you can, as opposed to `Add(1)` multiple times => (a little faster)

### WaitGroup Implementation

It's fairly easy to implement a WaitGroup. We only need couple of functionalities:

1. We need an `Add` method to **increment** the internal state **counter** (Atomically)
2. We need a `Done` method to **decrement** the internal state **counter** (Atomically)
3. We need a `Wait` method to **infinitely wait** till the internal state **counter** reaches **0** (Atomically)

### Tips

To test programs for race conditions, just use the `-race` flag
before running:

```shell script
cd into/the/example/dir

go run -race main.go
```

To run the benchmarks:

```shell script
cd benchmarks

# run all the benchmarks in the current directory
go test -bench=.

# run the benchmarks for 3s, by default it runs them for 1s
go test -bench=. -benchtime=3s
```

To run your programs by tracing your go routines, and the way they're scheduled:

```shell script
# run the program by tracing the Go Scheduler
GOMAXPROCS=1 GODEBUG=schedtrace=5000,scheddetail=1 go run main.go

# run the program by tracing the Go Scheduler on already built binary
go build -o exec
GOMAXPROCS=1 GODEBUG=schedtrace=5000,scheddetail=1 ./exec
```

### Zip Archives

- [Concurrency in Go #2 - WaitGroups (Part 1)](https://youtu.be/srb6fbioEY4) - [[Download Zip]](https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-2.tar.gz)
- [Concurrency in Go #3 - WaitGroups (Part 2)](https://youtu.be/zAMUKb6fCO0) - [[Download Zip]](https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-3.tar.gz)
- [Concurrency in Go #4 - WaitGroups (Part 3)](https://youtu.be/_QNcn7LAANY) - [[Download Zip]](https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-4.tar.gz)

### Presentations

- [Concurrency in Go #2, #3, #4 - WaitGroups](https://github.com/golang-basics/concurrency/raw/master/presentations/2_3_4_waitgroups)

### Examples

- [Without WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/without-waitgroup/main.go)
- [Basic Example](https://github.com/golang-basics/concurrency/blob/master/waitgroups/basic/main.go)
- [With WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/with-waitgroup/main.go)
- [Deadlock](https://github.com/golang-basics/concurrency/blob/master/waitgroups/deadlock/main.go)
- [Passed by Value](https://github.com/golang-basics/concurrency/blob/master/waitgroups/passed-by-value/main.go)
- [WaitGroup reuse before Wait() return - simple](https://github.com/golang-basics/concurrency/blob/master/waitgroups/wg-reuse/simple/main.go)
- [WaitGroup reuse before Wait() return - loop](https://github.com/golang-basics/concurrency/blob/master/waitgroups/wg-reuse/loop/main.go)
- [Too many calls to Done()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/done-too-many-times/main.go)
- [No calls to Add()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/no-add/main.go)
- [Limit Go Routines](https://github.com/golang-basics/concurrency/blob/master/waitgroups/limit-goroutines/main.go)
- [Rate Liming Example](https://github.com/golang-basics/concurrency/blob/master/waitgroups/rate-limiting/main.go)
- [Go Routines Order - Simple](https://github.com/golang-basics/concurrency/blob/master/waitgroups/goroutines-order/simple/main.go)
- [Go Routines Order - Preserve Order](https://github.com/golang-basics/concurrency/blob/master/waitgroups/goroutines-order/preserve-order/main.go)
- [Go Routines Order - Different Workloads](https://github.com/golang-basics/concurrency/blob/master/waitgroups/goroutines-order/different-workloads/main.go)
- [WaitGroup Implementation](https://github.com/golang-basics/concurrency/blob/master/waitgroups/waitgroup-implementation/main.go)
- [Benchmark - WaitGroup Add-One vs Add-Many](https://github.com/golang-basics/concurrency/blob/master/waitgroups/benchmarks/add1_vs_addmany_test.go)

### Resources

- [WaitGroup type declaration](https://github.com/golang/go/blob/master/src/sync/waitgroup.go#L20)
- [WaitGroup `Add()` method](https://github.com/golang/go/blob/master/src/sync/waitgroup.go#L53)
- [WaitGroup `Done()` method](https://github.com/golang/go/blob/master/src/sync/waitgroup.go#L98)
- [WaitGroup `Wait()` method](https://github.com/golang/go/blob/master/src/sync/waitgroup.go#L103)
- [`Wait()` synchronized with `Add`](https://github.com/golang/go/blob/master/src/sync/waitgroup.go#L124)
- [WaitGroups - GoByExample](https://gobyexample.com/waitgroups)
- [WaitGroups - CalHoun](https://www.calhoun.io/concurrency-patterns-in-go-sync-waitgroup/)

[Home](https://github.com/golang-basics/concurrency)
