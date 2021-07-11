# Atomic(s) - `sync/atomic`

**High Level languages** are cool, but nothing beats code that executes directly on the **CPU Level**.
In Go those are ***atomic(s)***. Running concurrent code can both be: **Slow** and **Cumbersome**
(due to the **synchronisation** involved). However, the `atomic.Value` type is the best
of both worlds: **blazing fast** and **less cumbersome** compared to other **sync types**.

When you think of **Atomic(s)**, it's nothing fancy or anything Go related,
it's really something which is as **close** as possible to the **CPU Level**,
which works directly with **CPU registers**, hence the blazing fast speed.

When you think of **Atoms**, these are the **smallest** **Indivisible** units. The same approach
pretty much is valid when it comes to Concurrency.

These are some examples of the simplest **Atomic Operations**: `+`, `-`, `*`, `/`, `=`
that respect all principles of **Atomicity**. 

So why is it called **Atomic**, and what exactly does the term **ATOMIC** mean?
Something is Atomic, if it **HAPPENS** from point A to point Z in its **ENTIRETY**
**WITHOUT BEING INTERRUPTED** by other processes in the same **Context**.

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

Most functionalities provided by the `sync/atomic` package are there to work with **numbers**.
These types are:

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

which facilitate Concurrency for a more complex/hybrid type.

**Note:** A more complex type `sync.Map` provided by the sync package
also makes use of `atomic.Value`

### Atomic(s) Downsides

- Mostly works with **Numbers**
- `atomic.Value` is generic (**type assertion** required)
- `atomic.Value` has **Limited Context** (cumbersome/impossible to use on sequential data)
- `atomic.Value` is **not always Atomic** (if not careful, we may run into race conditions)

### Behind the Scenes

Whether an operation executes concurrently or not, it has to either to with
some kind of **computation/calculus** or **memory access**. In the computers' era,
things are controlled and coordinated by the CPU, whether it's memory or computation.

#### CPU Registers

Alongside many things a CPU has, one of the most important things
are its **registers**.

What is a Register?
CPU registers are high-speed storage locations inside the microprocessor used
for **Memory Addressing**, **Data Operation**, and **Processing**.

A register is a **small** **high-speed memory** inside the **CPU**. It is used to **store** **temporary results**.
Registers are designed to be assessed at a much **higher speed** than conventional memory.
Some registers are **general purpose**, while others are **classified** according to the function they perform.

Some of these registers are:

- `Memory Address Register` (MAR)
- `Stack Control Register` (ACR)
- `Memory Data Register` (MDR)
- `Flag Register` (FR)
- `Accumulator Register` (AC)
- `Instruction Pointer Register` (IPR)
- `Data Register`
- `Address or Segment Register` (AR or SR)
- `Memory Buffer Register` (MBR)
- `Index Register`
- `Program Counter` (PC)

Shortly every **memory/computation** operation has to deal with **registers**.
A register can either be **used** or **not used**. What that means for concurrent code,
is that every process will get its turn to work with the register and there's
no need for any kind of **memory access synchronization** or other types of
measures we usually take in a high level language which executes concurrent code.

If we go back to **Atomic(s)**, these are **Go ASM (Assembly)** routines that execute directly on CPU
thus, **not needing** any kind of **memory access synchronisation**, and they are **blazing fast**
if benchmarked against other types in the sync package.

For more on what kind of registers are available check out the [CPU Registers](./registers.md)

#### Cross Compilation Workflow

Along with many other reasons why people choose Go, it also provides a beautiful and super
elegant code **cross compilation model**, which works beautifully pretty much on **any
operating system** and **architecture** available nowadays.

When you think of architectures, there's plenty of them. These are the ones Go supports:

- `386`
- `amd64`
- `arm`
- `arm64`
- `mips`
- `mipsle`
- `mips64`
- `mipsle64`
- `ppc64`
- `ppc64le`
- `riscv64`
- `s390x`
- `wasm`

Here's a small overview of how things really happen under the hood, from **compiling** your
Go source code, till you have a **generated binary**, that actually gets to be executed on
your **specific OS** and **Architecture**:

1. Go source Code (obviously :D)
2. The `$GOOS` and `$GOARCH` variables, accepting a bunch of OSes and Architectures
3. `Go Compiler` which compiles your code given the Go Code and `$GOOS` & `$GOARCH`
4. Once the compiler is done, it'll generate a `main.s`, which are `Go` pseudo `ASM` (Assembly) instructions
5. That is then picked up by the Go ASM (Assembler) tool and goes through the `obj` library
6. The `obj` library maps the **pseudo instructions** to **real architecture instructions**
generating an object file i.e `main.o` (this process is called assembling)
7. Then the `Go Linker` tool will pick the object file and **generate a binary** out of it (this process is called linking)
8. The generated binary is nothing but 0s and 1s, which is the only thing the CPU gets to eat

#### `go build` Workflow

When running the `go build` command, pretty much the exact same process described before happens.

1. Apply the **Cross Compilation** phase for the `main` package usually the `main.go` containing the `main()` func
2. Look for any **Go ASM files**, usually the ones that end with `.s` for the provided specific architecture.
3. **Create object files** for those ASM files
4. **Link** all **objects** and **generate** the final **binary**.

If your regular good old Go code, contains `functions` with `missing body`, the code
will expect some **external ASM declarations**, pretty much like is done in most C/C++ code.
This is crucial to making the **link between Go code and ASM code**.

Just like regular go files (`.go`) the ASM ones (`.s`) the same
`name_$GOOS_GOARCH.go` and respectively `name_$GOOS_GOARCH.s` rules apply.
Naming your files using the `$GOOS_$GOARCH` combination, will make the
build tool pick only the ones for the provided **OS** and respective **Architecture**.

For a more detailed example, have a look [here](https://github.com/golang-basics/concurrency/blob/master/atomics/asm/main.go)

#### Go ASM

Most statically typed languages only have 3 phases: `Compilation`->`Assembling`->`Linking`

Go has obviously more phases, to easier **accommodate multiple architectures** and make **cross
compilation** easy & make **builds faster** by eliminating some heavy work on
the assembly process.

**Go** has its own **Assembler**, which as I mentioned before generates **ASM like instructions**
which are pretty **architecture agnostic** and are pretty much pseudo instructions
meant to be used internally by the `obj` library, which then **maps** against **real
architecture instructions** and generates real **ASM instructions** for your
specific **OS** and **Architecture**.

Here are some ASM symbols and abbreviations Go uses internally

- `FP` Frame pointer: arguments and locals.
- `PC` Program counter: jumps and branches.
- `SB` Static base pointer: global symbols.
- `SP` Stack pointer: top of stack.

For more checkout out

- `$GOROOT/src/cmd/asm/internal/arch`
- `$GOROOT/src/cmd/internal/obj`

#### 32bit vs 64bit

Most Operating Systems out there run on either **32bit** or **64bit** architecture.
The question is, which one is better?

The answer is pretty simple, obviously 64bit.

There are 2 significant differences.

1. **Maximum number of ALU operations** performed: `32` and `64` respectively.

Which means, the CPU will be able to process that many **arithmetic** and **logical**
operations in a 32bit and respectively 64bit system.

2. **Max RAM memory available**: `3.25 GB` and respectively `16 EB` (16B GB)

Which of course means **larger CPU registers** 32bit vs 64bit values.
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

### Zip Archives

- [Concurrency in Go #5 - Atomics](https://youtu.be/srb6fbioEY4) - [[Download Zip]](https://github.com/golang-basics/concurrency/raw/master/archives/concurrency-5.tar.gz)

### Presentations

- [Concurrency in Go #5, #6 - Atomics](https://github.com/golang-basics/concurrency/raw/master/presentations/5_6_atomics)

### Examples

- [Race Example](https://github.com/golang-basics/concurrency/blob/master/atomics/race/main.go)
- [Basic Counter](https://github.com/golang-basics/concurrency/blob/master/atomics/basic-counter/main.go)
- [Race Fixed Example](https://github.com/golang-basics/concurrency/blob/master/atomics/race-fixed/main.go)
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
- [CPU Registers](https://sciencerack.com/types-of-cpu-registers/)
- [CPU Registers & Memory](https://www.doc.ic.ac.uk/~eedwards/compsys/memory/index.html#:~:text=Registers%20are%20memories%20located%20within,than%2064%20bits%20in%20size.)
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
