# Atomic(s) - `sync/atomic`

### Tips

```shell script
# compile the go program and generate an object file
# with pseudo assembly instructions
go tool compile -S main.go
go tool compile -S main.go > main.asm

# check out what gets put in the binary
go build -o exec main.go
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
- [atomic.Value](https://github.com/golang-basics/concurrency/blob/master/atomics/value/main.go)

### Resources

- [Atomicity - Wiki](https://en.wikipedia.org/wiki/Linearizability#Atomic)
- [Processor Register - Wiki](https://en.wikipedia.org/wiki/Processor_register)
- [Assembly Language - Wiki](https://en.wikipedia.org/wiki/Assembly_language)
- [Linker - Wiki](https://en.wikipedia.org/wiki/Linker_(computing))
- [Go Assembler - golang.org](https://golang.org/doc/asm)
- [Atomic Source Code](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L28)
- [Atomic ASM - Loadint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L19)
- [Atomic ASM - Load64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_386.s#L220)
- [Atomic ASM - Storeint64 Source](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L180)
- [386 Ops - Source](https://github.com/golang/go/blob/master/src/cmd/compile/internal/ssa/gen/386Ops.go#L30)
- [arch.Set - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L53)
- [arch.archX86 - Internal Source](https://github.com/golang/go/blob/master/src/cmd/asm/internal/arch/arch.go#L102)

[Home](https://github.com/golang-basics/concurrency)
