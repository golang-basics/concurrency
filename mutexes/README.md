# Mutexes

### Atomicity

An operation is considered **Atomic**, if within the **context** it is operating it is `Indivisible` or `Uninterruptible`.
The word that's important here is **Context**. Something may be **atomic** in one **context** but not in another.

Operations that are **Atomic** in the **Context** of your **Process**, may not be atomic in the context of the **Operating System**.
Operations that are **Atomic** in the **Context** of your **Operating System**, may not be **Atomic** in the **Context** of your **Machine**,
and **Operations** that are **Atomic** in the **Context** of your **Machine**, may not be **Atomic** in the **Context** of your **Application**.

**Indivisible** and **Uninterruptible** means, that within the **Context** you've defined something that is **Atomic**
`WILL HAPPEN` in its `ENTIRETY`, without anything else `HAPPENING SIMULTANEOUSLY` in the same context.

Let's take for example the statement `i++`. This may look like one Atomic operation, but in reality this happens:

1. Retrieve the value of `i`
2. Increment the value of `i`
3. Store the value of `i`

While **each** of the operations above are **atomic**, **the combination** of these in a certain **Context may not be**,
which also means, **combining several atomic operations** does not necessarily produce a **bigger atomic operation**.

### Memory Access Synchronization

**Memory Access Synchronization** comes hand in hand with **Atomicity**. In order to make a specific operation Atomic,
we MUST allow some kind of Memory Access Synchronization. If **multiple go routines** try to **read/write** from/to the **same
memory space**, they need a way to **communicate**, that **only 1 go routine at a time** can **read or write** from that **memory space**.

This kind of memory access synchronization is done through a process called **Mutual Exclusion**, which provides a **Locking
Mechanism**, so that when 1 concurrent process (go routine) tries to access some kind of memory space, that memory space
can be **guarded by a Mutex** which **holds a Lock** on that **memory space**.

The **Lock** can only be **acquired** by **only 1 go routine** at a time, thus making the rest of concurrent processes (go routines)
**wait** their turn to acquire the Lock, in order to read from or write to that memory space. This way a certain **memory
space** in the **context** of an operation is considered to be **Atomic**, resulting in **deterministic** and **correct** results when
**multiple concurrent operations** are involved in the game.

#### Benefits of Serializability

Let’s first understand the difference between a serial and non-serial schedule for a better understanding of the benefits
that serializability provides. In the case of a serial schedule, the multiple transactions involved are executed one
after the other sequentially with no overlap. This helps maintain the consistency in the database but limits the scope
of concurrency and often a smaller transaction might end up waiting for a long time due to an execution of a previous
longer transaction. Serial schedule also consumes a lot of CPU resources which gets wasted due to the serial execution.

In the case with a non-serial schedule, the multiple transactions executed are interleaved leading to inconsistency in
the database but at the same time helps overcome the disadvantages of a serial schedule such as concurrent execution
and wastage of CPU resources.

It’s established that the execution of multiple transactions in a non-serial schedule takes place concurrently.
And because of the multiple combinations involved, the output obtained may be incorrect at times which cannot be afforded.
This is where serializability comes into the picture and help us determine if the output obtained from a parallelly
executed schedule is correct or not.

In other words, Serializability serves as a measure of correctness for the transactions executed concurrently.
It serves a major role in concurrency control that is crucial for the database and is considered to provide
maximum isolation between the multiple transactions involved. The process of Serializability can also help in achieving
the database consistency which otherwise is not possible for a non-serial schedule.

### Presentations

- [Concurrency in Go #7 - Mutexes](https://github.com/golang-basics/concurrency/raw/master/presentations/7_mutexes)

### Examples

- [Atomic WaitGroup](https://github.com/golang-basics/concurrency/blob/master/mutexes/atomic-waitgroup/main.go)
- [Race Condition](https://github.com/golang-basics/concurrency/blob/master/mutexes/race-condition/main.go)
- [RWMutex](https://github.com/golang-basics/concurrency/blob/master/mutexes/read-write/main.go)
- [Mutex vs RWMutex - Basic Comparison](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/basic/main.go)
- [Mutex vs RWMutex - Benchmarks](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/benchmarks/mutex_vs_rwmutex_test.go)
- [sync.Locker](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/synclocker/main.go)
- [Lock Contention](https://github.com/golang-basics/concurrency/blob/master/mutexes/lock-contention/main.go)
- [Starvation](https://github.com/golang-basics/concurrency/blob/master/mutexes/starvation/main.go)
- [Deadlock - Circular Wait](https://github.com/golang-basics/concurrency/blob/master/mutexes/deadlocks/circular-wait/main.go)
- [Deadlock - Mutual Exclusion](https://github.com/golang-basics/concurrency/blob/master/mutexes/deadlocks/mutual-exclusion/main.go)
- [Deadlock - Hold and Wait](https://github.com/golang-basics/concurrency/blob/master/mutexes/deadlocks/hold-and-wait/main.go)
- [Deadlock - No Preemption](https://github.com/golang-basics/concurrency/blob/master/mutexes/deadlocks/no-preemption/main.go)
- [Deadlock - `runtime.Goexit()`](https://github.com/golang-basics/concurrency/blob/master/mutexes/deadlocks/goexit/main.go)
- [`sync.Once` - Simple Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/simple/main.go)
- [`sync.Once` - Increment/Decrement Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/inc-dec/main.go)
- [`sync.Once` - Race](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/race/main.go)
- [`sync.Once` - Deadlock](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/deadlock/main.go)
- [`sync.Once` - Once Implementation](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/once-implementation/main.go)
- [`sync.Once` - Once Implementation with Reset() - resync](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/resync/main.go)
- [`sync.Once` - Caching - Bad Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/bad/main.go)
- [`sync.Once` - Caching - Bad Benchmark](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/bad/bad_test.go)
- [`sync.Once` - Caching - Good Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/good/main.go)
- [`sync.Once` - Caching - Good Benchmark](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/good/good_test.go)
- [`sync.Once` - Caching - Better Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/better/main.go)
- [`sync.Once` - Caching - Better Benchmark](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/caching/better/better_test.go)
- [sync.Map](https://github.com/golang-basics/concurrency/blob/master/mutexes/syncmap/main.go)
- [builtin map vs sync.Map](https://github.com/golang-basics/concurrency/blob/master/mutexes/builtinmap-vs-syncmap/builtinmap_vs_syncmap_test.go)
- [sync.Cond - Too much CPU burst](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/too-much-cpu/main.go)
- [sync.Cond - Too much Wait (sleep)](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/too-much-cpu/main.go)
- [sync.Cond - Simple](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/simple/main.go)
- [sync.Cond - Broadcast](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/broadcast/main.go)
- [sync.Cond - Deadlock](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/deadlock/main.go)
- [sync.Cond - Writer/Validator/Reader Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/writer-validator-reader/main.go)
- [sync.Cond - Shopping Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/shopping/main.go)
- [sync.Cond - Enqueue/Dequeue Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/enqueue-dequeue/main.go)
- [sync.Cond - Button Example](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/enqueue-dequeue/main.go)
- [sync.Cond vs Channel - Signal](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/cond-vs-channel/signal/main.go)
- [sync.Cond vs Channel - Broadcast](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/cond-vs-channel/broadcast/main.go)
- [sync.Cond vs Channel - Benchmarks](https://github.com/golang-basics/concurrency/blob/master/mutexes/cond/cond-vs-channel/benchmarks/cond_vs_channel_test.go)

### Resources

- [Mutex - Wiki](https://en.wikipedia.org/wiki/Mutual_exclusion)
- [Reentrant/Recursive Mutex - Wiki](https://en.wikipedia.org/wiki/Reentrant_mutex)
- [Deadlock - Wiki](https://en.wikipedia.org/wiki/Deadlock)
- [Mutual Exclusion - Wiki](https://en.wikipedia.org/wiki/Mutual_exclusion)
- [Dining Philosophers Problem - Wiki](https://en.wikipedia.org/wiki/Dining_philosophers_problem)
- [Test and Set - Wiki](https://en.wikipedia.org/wiki/Test-and-set)
- [Tuple Space - Wiki](https://en.wikipedia.org/wiki/Tuple_space)
- [Message Passing - Wiki](https://en.wikipedia.org/wiki/Message_passing)
- [Semaphore - Wiki](https://en.wikipedia.org/wiki/Semaphore_(programming))
- [Monitor - Wiki](https://en.wikipedia.org/wiki/Monitor_(synchronization))
- [Spinlock - Wiki](https://en.wikipedia.org/wiki/Spinlock)
- [Concurrency Control - Wiki](https://en.wikipedia.org/wiki/Concurrency_control)
- [Serializability - Wiki](https://en.wikipedia.org/wiki/Serializability)
- [Serializability in DBMS - Educba](https://www.educba.com/serializability-in-dbms/)
- [Result Serializability - Geeks for Geeks](https://www.geeksforgeeks.org/result-serializability-in-dbms/)
- [Schedule - Wiki](https://en.wikipedia.org/wiki/Schedule_(computer_science))
- [Recoverability - Wiki](https://en.wikipedia.org/wiki/Schedule_(computer_science)#Recoverable)
- [2PL - Wiki](https://en.wikipedia.org/wiki/Two-phase_locking)
- [2PC - Wiki](https://en.wikipedia.org/wiki/Two-phase_commit_protocol)
- [Gossip Protocol - Wiki](https://en.wikipedia.org/wiki/Gossip_protocol)
- [Transaction Processing - Wiki](https://en.wikipedia.org/wiki/Transaction_processing)
- [Pessimistic vs Optimistic Locking - StackOverflow](https://stackoverflow.com/questions/129329/optimistic-vs-pessimistic-locking)
- [Pessimistic vs Optimistic Locking - StackOverflow Explanation](https://stackoverflow.com/a/58952004)
- [Check Deadlock - Go Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5122-L5221)
- [Mutex.Lock() - Go Source Code](https://github.com/golang/go/blob/master/src/sync/mutex.go#L76)
- [Mutex.lockSlow() - Go Source Code](https://github.com/golang/go/blob/master/src/sync/mutex.go#L108:17)
- [RaceAcquire - Go Source Code](https://github.com/golang/go/blob/master/src/runtime/race.go#L37)
- [raceacquire - Go Source Code](https://github.com/golang/go/blob/master/src/runtime/race.go#L515)
- [racecall - Go Source Code](https://github.com/golang/go/blob/master/src/runtime/race.go#L348)
- [racecall - GOASM Source Code](https://github.com/golang/go/blob/master/src/runtime/race_amd64.s#L384)
- [RWMutex - Go Source Code](https://github.com/golang/go/blob/master/src/sync/rwmutex.go#L28)

[Home](https://github.com/golang-basics/concurrency)
