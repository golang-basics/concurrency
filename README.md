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

### Go Scheduler

Primarily the Go scheduler has the opportunity to get triggered on these events:

- The use of the keyword go
- Garbage collection
- System calls (i.e. open file, read file, e.t.c.)
- Synchronization and Orchestration (channel read/write)

#### P, M, G

Once the syscall exists Go tries to apply one of the rules:

- try to acquire the exact same P, and resume the execution
- try to acquire a P in the idle list and resume the execution
- put the goroutine in the global queue and put the associated M back to the idle list

Goroutines do not go in the global queue only when the local queue is full;
it is also pushed in it when Go inject a list of goroutines to the scheduler,
e.g. from the network poller or goroutines asleep during the garbage collection

### Spinning Threads

### Net Poller

### SysMon

`sysmon` is smart enough to not consume resources when there is nothing to do.
Its cycle time is dynamic and depends on the current activity of the running program.
The initial pace is set at 20 nanoseconds, meaning the thread is constantly looking to help.
Then, after some cycles, if the thread is not doing anything, the sleep between two cycles
will double until it reaches 10ms.
If your application does not have many system calls or long-running goroutines,
the thread should back off to a 10ms delay most of its time, giving
a very light overhead to your application.

The thread is also able to detect when it should not run. Here are two cases:

- The garbage collector is going to run. sysmon will resume when the garbage collector ends.
- All the threads are idle, nothing is running.

### Work Stealing

Here's how Go makes sure to equally distribute & balance work
and make use of computer resources as efficient as possible:

- pull work from the local queue
- pull work from the global queue
- pull work from network poller
- steal work from the other P’s local queues

### Tracing

```bash
GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
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
- [GRQ Check - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L3343)
- [Force Preempt Duration for G - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5435)
- [Go Scheduler Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit)
- [Scheduling in Go (Part 2) - Ardan Labs](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
- [Scheduling in Go (Part 3 - Concurrency) - Ardan Labs](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)
- [Scheduler Tracing in Go - Ardan Labs](https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html)
- [The Scheduler Saga](https://about.sourcegraph.com/go/gophercon-2018-the-scheduler-saga/#:~:text=The%20scheduler%20needs%20to%20exist,it%20multiplexes%20goroutines%20onto%20threads.)
- [Stack size](https://github.com/golang/go/blob/master/src/runtime/stack.go#L73)
- [Golang Net Poller Source Code](https://github.com/golang/go/blob/master/src/runtime/netpoll.go)
- [Golang Net Poller](https://morsmachine.dk/netpoller)
- [Preemptive vs Non-Preemptive Scheduling](https://www.guru99.com/preemptive-vs-non-preemptive-scheduling.html#:~:text=In%20Preemptive%20Scheduling%2C%20the%20CPU,Schedulign%20no%20switching%20takes%20place.)
- [Preemptive vs Cooperative](https://www.geeksforgeeks.org/difference-between-preemptive-and-cooperative-multitasking/#:~:text=Preemptive%20multitasking%20is%20a%20task,running%20process%20to%20another%20process.)
- [Go 1.14 Release Notes - Runtime](https://golang.org/doc/go1.14#runtime)
- [Go Asynchronous Preemption](https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c)
- [Go Routine & Preemption](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)
- [Go Routine, OS Thread and CPU Management](https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a)
- [Go Work Stealing - Go Scheduler](https://medium.com/a-journey-with-go/go-work-stealing-in-go-scheduler-d439231be64d)
- [Guarded Command Language](https://en.wikipedia.org/wiki/Guarded_Command_Language)
- [System Monitor - sysmon](https://en.wikipedia.org/wiki/System_monitor)
- [Go SysMon Runtime Monitoring](https://medium.com/@blanchon.vincent/go-sysmon-runtime-monitoring-cff9395060b5)
- [SysMon - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5273)
- [Garbage Collector Period - Source Code](https://github.com/golang/go/blob/master/src/runtime/proc.go#L5268)
