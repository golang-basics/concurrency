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

### Presentations

- [Concurrency in Go #3 - Mutexes](https://github.com/golang-basics/concurrency/raw/master/presentations/3_mutexes)

### Examples

- [Atomic WaitGroup](https://github.com/golang-basics/concurrency/blob/master/mutexes/atomic-waitgroup/main.go)
- [Race Condition](https://github.com/golang-basics/concurrency/blob/master/mutexes/race-condition/main.go)
- [RWMutex](https://github.com/golang-basics/concurrency/blob/master/mutexes/read-write/main.go)
- [Mutex vs RWMutex - Basic Comparison](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/basic/main.go)
- [Mutex vs RWMutex - Benchmarks](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/benchmarks/mutex_vs_rwmutex_test.go)
- [sync.Locker](https://github.com/golang-basics/concurrency/blob/master/mutexes/mutex-vs-rwmutex/synclocker/main.go)
- [Lock Contention](https://github.com/golang-basics/concurrency/blob/master/mutexes/lock-contention/main.go)
- [Starvation](https://github.com/golang-basics/concurrency/blob/master/mutexes/starvation/main.go)
- [sync.Once - Simple](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/simple/main.go)
- [sync.Once - Increment/Decrement](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/inc-dec/main.go)
- [sync.Once - Deadlock](https://github.com/golang-basics/concurrency/blob/master/mutexes/once/deadlock/main.go)
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

[Home](https://github.com/golang-basics/concurrency)
