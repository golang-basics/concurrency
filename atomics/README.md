# Atomic(s) - `sync/atomic`

### atomic

`sync.Value` => `sync.Map` uses `sync.Value`

### Tips

```shell script
# generate object file from Go's pseudo assembly instructions
go tool asm arithmetics.s

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

- [Concurrency in Go #5 - Atomics](https://github.com/golang-basics/concurrency/raw/master/presentations/5_atomics)

### Examples

- [Atomicity](https://github.com/golang-basics/concurrency/blob/master/atomics/atomicity/main.go)
- [Basic Counter](https://github.com/golang-basics/concurrency/blob/master/atomics/basic/main.go)
- [atomic.Value - panic](https://github.com/golang-basics/concurrency/blob/master/atomics/value/panic/main.go)
- [atomic.Value - reader/writer](https://github.com/golang-basics/concurrency/blob/master/atomics/value/reader-writer/main.go)
- [atomic.Value - not atomic](https://github.com/golang-basics/concurrency/blob/master/atomics/value/not-atomic/main.go)
- [Calculator Example](https://github.com/golang-basics/concurrency/blob/master/atomics/calculator/main.go)
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
- [Atomic Source Code](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L28)
- [Atomic ASM - Loadint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L19)
- [`sync/runtime/internal/atomic/atomic_amd64` - Store64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L171)
- [`sync/atomic/asm.s` - LoadInt64 Source](https://github.com/golang/go/blob/master/src/sync/atomic/asm.s#L61)
- [Atomic ASM - Load64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_386.s#L220)
- [Atomic ASM - Storeint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L180)
- [386 Ops - Source](https://github.com/golang/go/blob/master/src/cmd/compile/internal/ssa/gen/386Ops.go#L30)
- [386 Instructions - Source](https://github.com/golang/go/blob/release-branch.go1.5/src/cmd/internal/obj/x86/anames.go#L8)
- [arch.Set - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L53)
- [arch.archX86 - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L102)
- [go build - Internal Source](https://github.com/golang/go/blob/master/src/cmd/go/internal/work/build.go#L366)
- [Function declarations - Spec](https://golang.org/ref/spec#Function_declarations)

[Home](https://github.com/golang-basics/concurrency)
