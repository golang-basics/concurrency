# Wait Groups

### The Problem

In the real world of applications, not every operation inside a program executes concurrently,
however many operations will most probably do. A lot of the times though synchronous code
needs a certain condition to be true in order to proceed, thus waiting for a certain amount
of concurrent actions to complete first.

The order in which concurrent code executes is always non-deterministic, thus the way in which
the `go` keyword appears, does not mean that's the order in which way the Go Scheduler will
schedule them or that, that's the order in which they will actually get to run.
That being said, sometimes there is a need to preserve order, but still execute our
operations concurrently.

Also in a real world, a service does not always run standalone, meaning sometimes our applications
rely on other downstream applications in order to produce a result. In many cases we do not control
those downstream applications, and many of those applications may have API Rate Limiting, which means
our upstream application has to be careful on how many requests it makes, and at which rate they happen.

So we need some kind of waiting mechanism for concurrent code. We need to not wait too long,
so that our system is efficient and makes the best use of our resources, but also wait just enough,
so that the order is preserved, certain conditions are met, and we have control over
the amount of concurrent operations happening in our system.

Luckily there's no need to reinvent the wheel, Go introduces the `sync.WaitGroup` type,
which does exactly all the above mentioned.

### Go Concurrency Philosophy

Go is a powerful language, and it gained its popularity because people hear
its Concurrency model is one of the best. While Go has quite an elegant way
of doing Concurrency, Go did not give up the Classic way of doing Concurrency.

In Go there are 2 ways of making sure concurrent code executes correctly and safely:

1. Concurrency Primitives
2. Channels

Very important to stress out, it's not one versus the other, and there's no better
way of executing concurrent code using one or the other. It all depends on the
type of scenario and use case. Generally everything that can be achieved with channels
can be achieved with `sync` Concurrency Primitives as well. It's a matter of choice depending on
Complexity,Composition,Performance or in general on what's being done with the data.

### Concurrency Primitives

All of Go's Concurrency Primitives are stored inside the `sync` package, which stands for
Synchronization, because most of the times that's what we as developer do with the
executing concurrent code.

These are all the available types under the `sync` package:

- `WaitGroup`
- `Mutex`
- `RWMutex`
- `Locker`
- `Cond`
- `Map`
- `Pool`

### Atomicity

An operation is considered Atomic, if within the context it is operating it is Indivisible or Uninterruptible.
The word that's important here is Context. Something may be atomic in one context but not in another.

Operations that are Atomic in the Context of your Process, may not be atomic in the context of the Operating System.
Operations that are atomic in the Context of your Operating System, may not be Atomic in the Context of your Machine,
and Operations that are Atomic in the Context of your Machine, may not be Atomic in the Context of your Application.

Indivisible and Uninterruptible means, that within the Context you've defined something that is Atomic
WILL HAPPEN in its ENTIRETY, without anything else HAPPENING SIMULTANEOUSLY in the same context.

Let's take for example the statement `i++`. This may look like one Atomic operation, but in reality this happens:

1. Retrieve the value of `i`
2. Increment the value of `i`
3. Store the value of `i`

While each of the operations above are atomic, the combination of these in a certain Context may not be,
which also means, combining several atomic operations does not necessarily produce a bigger atomic operation.

### WaitGroup Overview

In order to achieve tasks like:

- Waiting on Concurrent Operations to finish
- Rate Limiting the amount of max Concurrent Operations
- Preserve Concurrent Operations order

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

### Examples

- [No WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/no-waitgroup/main.go)
- [Basic Example](https://github.com/golang-basics/concurrency/blob/master/waitgroups/basic/main.go)
- [WaitGroup Parallel](https://github.com/golang-basics/concurrency/blob/master/waitgroups/waitgroup-parallel/main.go)
- [Deadlock](https://github.com/golang-basics/concurrency/blob/master/waitgroups/deadlock/main.go)
- [Passed by Value](https://github.com/golang-basics/concurrency/blob/master/waitgroups/passed-by-value/main.go)
- [WaitGroup reuse before Wait() return](https://github.com/golang-basics/concurrency/blob/master/waitgroups/wg-reuse/main.go)
- [Too many calls to Done()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/done-too-many-times/main.go)
- [No calls to Add()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/no-add/main.go)
- [Limit Go Routines](https://github.com/golang-basics/concurrency/blob/master/waitgroups/limit-goroutines/main.go)
- [Go Routines Order](https://github.com/golang-basics/concurrency/blob/master/waitgroups/goroutines-order/main.go)
- [Atomic WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/atomic-waitgroup/main.go)
- [WaitGroup Implementation](https://github.com/golang-basics/concurrency/blob/master/waitgroups/waitgroup-implementation/main.go)
- [Benchmark - WaitGroup Add-One vs Add-Many](https://github.com/golang-basics/concurrency/blob/master/waitgroups/benchmarks/add1_vs_addmany_test.go)

[Home](https://github.com/golang-basics/concurrency)
