# Atomic(s) - `sync/atomic`

High Level languages are cool, but nothing beats code that executes directly on the CPU Level.
In Go those are atomic(s). Running concurrent code can both be: Slow and Cumbersome
(due to the synchronisation involved). However, the `atomic.Value` type is the best
of both worlds: blazing fast and less cumbersome compared to other sync types.

When you think of Atomic(s), it's nothing fancy or anything Go related,
it's really something which is as close as possible to the CPU Level,
which works directly with CPU registers, hence the blazing fast speed.

When you think of Atoms, these are the smallest Indivisible units. The same approach
pretty much is valid when it comes to Concurrency.

These are some examples of the simplest Atomic Operations: `+`, `-`, `*`, `/`, `=`
that respect all principles of Atomicity. 

So why is it called Atomic, and what exactly does the term ATOMIC mean?
Something is Atomic, if it happens from point A to point Z in its ENTIRETY
without being interrupted by other process in the same Context.

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

### `sync/atomic`

Most functionalities provided by the `sync/atomic` package are there to work with numbers.
Types like:

- `int32`
- `int64`
- `uint32`
- `uint64`
- `uintptr`
- `Pointer`

Respectively available functions for each of those types:

- `Load__`
- `Add__`
- `Store__`
- `Swap__`
- `CompareAndSwap__`


Besides, working mostly with numbers (except `Pointer`) the `sync/atomic` package
also provides the `atomic.Value` type which has the handy methods:

- `Store`
- `Load`

which facilitates Concurrency for a more complex/hybrid type.

`sync.Value` => `sync.Map` uses `sync.Value`

### Behind the Scenes

#### 32bit vs 64bit

#### CPU Registers

#### Cross Compilation Workflow

A 64-bit register can theoretically reference 18,446,744,073,709,551,616 bytes,
or 17,179,869,184 gigabytes (16 exabytes) of memory.
This is several million times more than an average workstation would need to access.

### Tips

```shell script
# cd into asm directory
cd asm

# generate object file from Go's pseudo assembly instructions
go tool asm sqrt_amd64.s

# compiles the go program and generates an object file
# with pseudo assembly instructions
go tool compile -S main.go

# compiles the go program and generates an object file
# with pseudo assembly instructions
# and saves the instructions in the specified file
go tool compile -S main.go > main.s

# generate a binary from the specified object file
go tool link main.o

# compiles and generates a binary from main.go
go build -o exec main.go
# dumps the assembly instructions from the generated binary
go tool objdump -s main.main exec
```

Here are some symbols and abbreviations Go uses internally

- FP: Frame pointer: arguments and locals.
- PC: Program counter: jumps and branches.
- SB: Static base pointer: global symbols.
- SP: Stack pointer: top of stack.

For more checkout out

- `$GOROOT/src/cmd/asm/internal/arch`
- `$GOROOT/src/cmd/internal/obj`

### Zip Archives

- [Concurrency in Go #5 - Atomics](https://youtu.be/srb6fbioEY4) - [[Download Zip]](https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-5.tar.gz)

### Presentations

- [Concurrency in Go #5 - Atomics](https://github.com/golang-basics/concurrency/raw/master/presentations/5_6_atomics)

### Examples

- [Race Example](https://github.com/golang-basics/concurrency/blob/master/atomics/race/main.go)
- [Basic Counter](https://github.com/golang-basics/concurrency/blob/master/atomics/basic-counter/main.go)
- [atomic.Value - Reader/Writer Example](https://github.com/golang-basics/concurrency/blob/master/atomics/value/reader-writer/main.go)
- [atomic.Value - Calculator Example](https://github.com/golang-basics/concurrency/blob/master/atomics/value/calculator/main.go)
- [atomic.Value - panic](https://github.com/golang-basics/concurrency/blob/master/atomics/value/panic/main.go)
- [atomic.Value - not atomic](https://github.com/golang-basics/concurrency/blob/master/atomics/value/not-atomic/main.go)
- [ASM Example](https://github.com/golang-basics/concurrency/blob/master/atomics/asm/main.go)
- [Atomic Implementation](https://github.com/golang-basics/concurrency/blob/master/atomics/atomic-implementation/main.go)
- [Benchmarks](https://github.com/golang-basics/concurrency/blob/master/atomics/benchmarks/number_vs_value_test.go)

### Resources

- [Atomicity - Wiki](https://en.wikipedia.org/wiki/Linearizability#Atomic)
- [Processor Register - Wiki](https://en.wikipedia.org/wiki/Processor_register)
- [Assembly Language - Wiki](https://en.wikipedia.org/wiki/Assembly_language)
- [Linker - Wiki](https://en.wikipedia.org/wiki/Linker_(computing))
- [Go Compiler Intrinsics - Dave Cheney](https://dave.cheney.net/2019/08/20/go-compiler-intrinsics)
- [Go Assembler Docs - golang.org](https://golang.org/doc/asm)
- [Go Tool (asm) - golang.org](https://golang.org/cmd/asm/)
- [Go Tool (compile) - golang.org](https://golang.org/cmd/compile/)
- [Go Tool (link) - golang.org](https://golang.org/cmd/link/)
- [Function declarations - Spec](https://golang.org/ref/spec#Function_declarations)
- [`math.Sqrt` - Source Code Go](https://github.com/golang/go/blob/master/src/math/sqrt.go#L92)
- [`math.Sqrt` - Source Code ASM Go](https://github.com/golang/go/blob/master/src/math/sqrt_asm.go#L12)
- [`math.Sqrt` - Source Code ASM](https://github.com/golang/go/blob/master/src/math/sqrt_amd64.s#L8)
- [`math.Floor` - Source Code Go](https://github.com/golang/go/blob/master/src/math/floor.go#L13)
- [`math.Floor` - Source Code ASM Go](https://github.com/golang/go/blob/master/src/math/floor_asm.go#L12)
- [`math.Floor` - Source Code ASM](https://github.com/golang/go/blob/master/src/math/floor_amd64.s#L10)
- [Atomic Source Code - Doc Go - LoadInt64](https://github.com/golang/go/blob/master/src/sync/atomic/doc.go#L114)
- [Atomic Source Code - ASM - LoadInt64](https://github.com/golang/go/blob/master/src/sync/atomic/asm.s#L61)
- [Atomic Source Code - Internal Go - LoadInt64](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L28)
- [Atomic ASM - Loadint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L19)
- [`sync/runtime/internal/atomic/atomic_amd64` - Store64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L171)
- [Atomic Source Code - Doc Go - StoreInt64](https://github.com/golang/go/blob/master/src/sync/atomic/doc.go#L132)
- [Atomic Source Code - Internal Go - StoreInt64](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L101)
- [Atomic Source Code - ASM - StoreInt64](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L171)
- [Atomic ASM - Storeint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L180)
- [go build - Internal Source](https://github.com/golang/go/blob/master/src/cmd/go/internal/work/build.go#L366)
- [arch.Set - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L53)
- [arch.archX86 - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L102)
- [386 Ops - Source](https://github.com/golang/go/blob/master/src/cmd/compile/internal/ssa/gen/386Ops.go#L30)
- [386 Instructions - Source](https://github.com/golang/go/blob/master/src/cmd/internal/obj/x86/anames.go#L8)
- [`sync.Map` - `atomic.Value` - Source Code](https://github.com/golang/go/blob/master/src/sync/map.go#L39)

[Home](https://github.com/golang-basics/concurrency)
