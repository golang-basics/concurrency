# Concurrency in Go

### Go routines

A go routines can block for one of these reasons:

- Sending/Receiving on channel
- Network or I/O
- Blocking System Call
- Timers
- Mutexes

#### Fairness

- Infinite loop — preemption (~10ms time slice)
- Local Run queue — preemption (~10ms time slice)
- Global run queue starvation is avoided by checking the global run queue for every 61 scheduler tick
- Network Poller Starvation Background thread poll network occasionally if not polled by the main worker thread

### Channels

Here are couple of simple rules to make sure channels are used correctly

- Before writing to a channel, make sure someone else is reading from it (deadlock)
- Before reading from a channel, make sure someone else is writing to it (deadlock)
- When ranging over a channel, ALWAYS make sure the producer closes the channel eventually (deadlock)
- Writing to a closed channel will result in a runtime panic
- Reading from a closed channel won't have any effects
- A channel close, is considered a write operation

### Mutexes

### Wait Groups

### Atomics

#### Poll Order

- Local Run queue
- Global Run queue
- Network Poller
- Work Stealing

### OSX `sysctl`

```bash
# get the number of logical CPU cores
sysctl hw.logicalcpu

# get the number of physical CPU cores
sysctl hw.physicalcpu

# get the number of logical cores
sysctl hw.ncpu

# get the number of physical/logical cores
# also thread count meaning the total count of running threads in parallel
sysctl -a | grep machdep.cpu | grep count
```

### Resources

- [OSX - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_darwin.go#L151)
- [Windows - number of CPUs](https://github.com/golang/go/blob/master/src/runtime/os_windows.go#L356)
- [Go Scheduler by rakyll](https://rakyll.org/scheduler/)
- [Go Scheduler by morsmachine](https://morsmachine.dk/go-scheduler)
- [Illustrated Tales of Go Runtime Scheduler](https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go)
- [Main Goroutine](https://github.com/golang/go/blob/master/src/runtime/proc.go#L144)
- [Go Scheduler Implementation](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3470)
- [G - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L403)
- [M - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L503)
- [P - Source Code](https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L596)
- [Go Scheduler Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit)
- [Scheduling in Go (Part 2) - Ardan Labs](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
- [Scheduling in Go (Part 3 - Concurrency) - Ardan Labs](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)
- [Stack size](https://github.com/golang/go/blob/master/src/runtime/stack.go#L73)
- [Golang Net Poller Source Code](https://github.com/golang/go/blob/master/src/runtime/netpoll.go)
- [Golang Net Poller](https://morsmachine.dk/netpoller)
- [Preemptive vs Non-Preemptive Scheduling](https://www.guru99.com/preemptive-vs-non-preemptive-scheduling.html#:~:text=In%20Preemptive%20Scheduling%2C%20the%20CPU,Schedulign%20no%20switching%20takes%20place.)
- [Preemptive vs Cooperative](https://www.geeksforgeeks.org/difference-between-preemptive-and-cooperative-multitasking/#:~:text=Preemptive%20multitasking%20is%20a%20task,running%20process%20to%20another%20process.)
- [Go 1.14 Release Notes - Runtime](https://golang.org/doc/go1.14#runtime)
- [Go Asynchronous Preemption](https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c)
- [Go Routine & Preemption](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)
- [Guarded Command Language](https://en.wikipedia.org/wiki/Guarded_Command_Language)
